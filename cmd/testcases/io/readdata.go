package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadData loads the data file
func ReadData(path string) ([][]float64, error) {
	var data [][]float64

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var features []float64
		splitFeatures := strings.Split(scanner.Text(), ",")
		splitFeatures = splitFeatures[:len(splitFeatures)-1]
		for _, val := range splitFeatures {
			feature, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, fmt.Errorf("Error encountered while parsing data file: %v", err)
			}

			features = append(features, feature)
		}
		data = append(data, features)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data, nil
}
