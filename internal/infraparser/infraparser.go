package infraparser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/dh-manoj/ds-cluster-utility/internal/cluster"
)

type InfraParser struct {
	CodePath string
	Env      string
}

func NewInfraParser(codePath string, env string) *InfraParser {
	return &InfraParser{
		CodePath: codePath,
		Env:      env,
	}
}

func (ip *InfraParser) IsStaging() bool {
	return ip.Env == "stg"
}

func (ip *InfraParser) IsLive() bool {
	return ip.Env == "live"
}

type FileInfo struct {
	FilePath     string
	FileNameCode string
}

func split(r rune) bool {
	return r == '-' || r == '.'
}

func (ip *InfraParser) ParseCreateCluster() {
	clusterDirPath := fmt.Sprintf("%s/terraform/%s/infra/", ip.CodePath, ip.Env)

	items, _ := ioutil.ReadDir(clusterDirPath)
	for _, item := range items {
		if !item.IsDir() && strings.Contains(item.Name(), "cluster") {
			names := strings.FieldsFunc(item.Name(), split)
			cluster.Register(clusterDirPath+item.Name(), names[1], ip.Env)
		}
	}
}

var regExValueWithinDoubleQuotes = regexp.MustCompile("\"(.*?)\"")

func valueWithinDoubleQuotes(value string) string {
	match := regExValueWithinDoubleQuotes.FindStringSubmatch(value)
	return match[1]
}

func formClusterName(fileNameCode, code, region, clusterName string) string {
	regionVar := fmt.Sprintf("${local.%s_region}", fileNameCode)
	countryCodeVar := fmt.Sprintf("${local.%s_country_code}", fileNameCode)

	if strings.Contains(clusterName, regionVar) {
		clusterName = strings.Replace(clusterName, regionVar, region, -1)
	}
	if strings.Contains(clusterName, countryCodeVar) {
		clusterName = strings.Replace(clusterName, countryCodeVar, code, -1)
	}
	return clusterName
}

func (ip *InfraParser) ParseClusterFile(cl *cluster.Cluster) {
	file, err := os.Open(cl.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rxCode := regexp.MustCompile("(country_code|country)[ =]*")
	rxRegion := regexp.MustCompile(fmt.Sprintf("^ *[a-z_]*region[ =]*"))
	rxZone := regexp.MustCompile(fmt.Sprintf("^ *[a-z_]*zone[ =]*"))
	rxClusterName := regexp.MustCompile("cluster_name[ =]*")
	rxClusterZone := regexp.MustCompile(fmt.Sprintf("^ *zone[ =]*"))

	isCodeFound := false
	isRegionFound := false
	isZoneFound := false
	isClusterNameFound := false
	isClusterZoneFound := false

	code := ""
	region := ""
	zone := ""
	clusterName := ""
	// clusterZone := ""

	skipClusterNameCount := 0
	for scanner.Scan() {
		if isCodeFound && isRegionFound && isZoneFound && isClusterNameFound && isClusterZoneFound {
			break
		} else if !isCodeFound && rxCode.MatchString(scanner.Text()) {
			code = valueWithinDoubleQuotes(scanner.Text())
			isCodeFound = true
		} else if !isRegionFound && rxRegion.MatchString(scanner.Text()) {
			region = valueWithinDoubleQuotes(scanner.Text())
			isRegionFound = true
		} else if !isZoneFound && rxZone.MatchString(scanner.Text()) {
			zone = valueWithinDoubleQuotes(scanner.Text())

			//hacky code - because of different zone name confusion in cluster-xx.tf file
			if cl.HasOverwriteZone() {
				zone = cl.OverwriteZone
			}
			isZoneFound = true
		} else if !isClusterNameFound && rxClusterName.MatchString(scanner.Text()) {
			if skipClusterNameCount < cl.NumberOfTimesSkipClusterName {
				//skip cluster name
				skipClusterNameCount += 1
			} else {
				clusterName = valueWithinDoubleQuotes(scanner.Text())
				isClusterNameFound = true
			}
		} else if isClusterNameFound && !isClusterZoneFound && rxClusterZone.MatchString(scanner.Text()) {
			isClusterZoneFound = true
			if strings.Contains(scanner.Text(), "\"") {
				//overwrite zone in case a diffrent zone is assigned for a cluster
				zone = valueWithinDoubleQuotes(scanner.Text())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cl.Code = code
	cl.Region = region
	cl.Zone = zone
	cl.ClusterName = formClusterName(cl.FileNameCode, cl.Code, cl.Region, clusterName)
}
