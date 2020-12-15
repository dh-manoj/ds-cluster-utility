package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dh-manoj/ds-cluster-utility/internal/cluster"
	"github.com/dh-manoj/ds-cluster-utility/internal/country"
	"github.com/dh-manoj/ds-cluster-utility/internal/infraparser"
)

func infraPath() string {
	path := os.Getenv("DH_DARKSTORE_INFRA_PATH")

	if path == "" {
		log.Fatalf("env DH_DARKSTORE_INFRA_PATH not set. Please set the enviroment to infra path")
	}
	return path
}

func parseClusterInfra(env string) {
	infraParser := infraparser.NewInfraParser(infraPath(), env)
	infraParser.ParseCreateCluster()

	//Overwrite skip
	cluster.Clusters["qa"].NumberOfTimesSkipClusterName = 1

	for _, cl := range cluster.Clusters {
		infraParser.ParseClusterFile(cl)
	}
}

func printInfo() {
	fmt.Printf("1-code|2-country-code|3-country|4-region|5-zone|6-cluster-name|7-gke-cluster-name|8-port|9-db\n")
	for i, _ := range country.Countries {
		if cl, ok := cluster.Clusters[strings.ToLower(country.Countries[i].Code)]; ok {
			fmt.Printf("%s|%s|%s|%s|%s|%s|%s|%d|%s\n",
				cl.DisplayCountryCode(),
				cl.DisplayCode(),
				cl.Country.Name,
				cl.Region,
				cl.Zone,
				cl.ClusterName,
				cl.DisplayGKEClusterName(),
				cl.Country.Port,
				cl.DisplayDB(),
			)
		}
	}
}

func run() {
	parseClusterInfra("live")
	parseClusterInfra("stg")
	printInfo()
}

func main() {
	run()
}
