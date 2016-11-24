package main

import (
	io "partitionclustering/cmd/testcases/io"
	"partitionclustering/kmedoids"

	log "github.com/Sirupsen/logrus"
)

func main() {
	kmedoids := kmedoids.NewKMedoids(3)
	data, err := io.ReadData("../../../static/datasets/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = kmedoids.Fit(data)
	if err != nil {
		log.Fatal(err)
	}
}
