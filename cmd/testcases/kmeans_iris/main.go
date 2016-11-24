package main

import (
	"partitionclustering/cmd/testcases/io"
	"partitionclustering/kmeans"

	log "github.com/Sirupsen/logrus"
)

func main() {
	// Create a kmeans object of 3 cluster, 4 features, and random
	// initialization set to false
	kmeans := kmeans.NewKMeans(3, 4, false)
	data, err := io.ReadData("../../../static/datasets/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = kmeans.Fit(data)
	if err != nil {
		log.Fatal(err)
	}
}
