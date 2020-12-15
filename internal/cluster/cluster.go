package cluster

import (
	"fmt"
	"log"

	"github.com/dh-manoj/ds-cluster-utility/internal/country"
)

type Cluster struct {
	FilePath                     string
	Country                      *country.Country
	Env                          string
	FileNameCode                 string
	Port                         int
	NumberOfTimesSkipClusterName int
	OverwriteZone                string //overwrite with the correct zone in case of zone incorrect in the file
	//Fetched by parsing
	Code        string
	Region      string
	Zone        string
	ClusterName string
}

func newCluster(filePath string, fileNameCode string, env string) *Cluster {
	cluster := &Cluster{
		Country:                      country.Lookup(fileNameCode),
		FilePath:                     filePath,
		FileNameCode:                 fileNameCode,
		Port:                         54000,
		NumberOfTimesSkipClusterName: 0,
		OverwriteZone:                "",
		Env:                          env,
	}
	return cluster
}

func (c *Cluster) HasOverwriteZone() bool {
	return c.OverwriteZone != ""
}

func (c *Cluster) IsLive() bool {
	return c.Env == "live"
}

func (c *Cluster) IsStaging() bool {
	return c.Env == "stg"
}

// >=2 letter code for country
func (c *Cluster) DisplayCountryCode() string {
	if c.Env == "stg" {
		if c.Code == "kw" {
			return "stg"
		} else if c.Code == "lt" {
			return "test"
		} else {
			log.Fatalf("%s:%s", c.FilePath, c.Code)
		}
	}
	return c.Code
}

// 2 letter code for country (except for stg environment)
func (c *Cluster) DisplayCode() string {
	if c.Env == "stg" {
		if c.Code == "kw" {
			return "stg"
		} else if c.Code == "lt" {
			return "test"
		} else {
			log.Fatalf("%s:%s", c.FilePath, c.Code)
		}
	}
	return c.Code[:2]
}

func (c *Cluster) DisplayGKEClusterName() string {
	return fmt.Sprintf("gke_dh-darkstores-%s_%s_%s", c.Env, c.Zone, c.ClusterName)
}

func (c *Cluster) DisplayDB() string {
	// Example:
	//        "dh-darkstores-stg:europe-west1:cloudsql-stg-europe-v2=tcp:54325" # stg
	//        "dh-darkstores-live:asia-southeast1:cloudsql-live-tw${REPLICA}=tcp:54322" # tw
	//        "dh-darkstores-live:europe-west1:cloudsql-live-kwt${REPLICA}=tcp:54323" # kw
	prefix := fmt.Sprintf("dh-darkstores-%s:%s:cloudsql-%s-", c.Env, c.Region, c.Env)
	postfix := c.Code
	if c.Code == "pk" {
		postfix += "-v1"
	}
	if c.Env == "stg" {
		postfix = "europe-v2"
	}
	return prefix + postfix
}

var Clusters = make(map[string]*Cluster)

func Register(filePath string, fileNameCode string, env string) {
	tempCode := fileNameCode
	if fileNameCode == "arg" {
		tempCode = "ar"
	}
	cl := newCluster(filePath, fileNameCode, env)
	Clusters[tempCode] = cl
}
