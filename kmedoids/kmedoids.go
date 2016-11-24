package kmedoids

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"

	dp "partitionclustering/datapoint"
	"partitionclustering/distances"
	"partitionclustering/errors"
)

// KMedoids represents a KMedoids
type KMedoids struct {
	NumClusters int
	// Index of the datapoints used as medoids
	Medoids           []int
	Datapoints        []dp.Datapoint
	PointsAssignments []int
	DistanceMetric    float64
}

// NewKMedoids creates a KMedoids object.
// This implementation of KMedoids uses the PAM (partitioning around medoids) method
// http://www.math.le.ac.uk/people/ag153/homepage/KmeansKmedoids/Kmeans_Kmedoids.html
func NewKMedoids(numClusters int) *KMedoids {
	return &KMedoids{
		NumClusters:    numClusters,
		Medoids:        make([]int, numClusters),
		DistanceMetric: distances.Manhattan, // for the moment use this one

	}
}

// Fit trains the KMedoids model
func (km *KMedoids) Fit(data [][]float64) error {
	if km.NumClusters > len(data) {
		return errors.NewNumClustersGreaterNumDataError(km.NumClusters, len(data))
	}

	km.PointsAssignments = make([]int, len(data))

	for _, datapoint := range data {
		km.Datapoints = append(km.Datapoints, *dp.NewDataPoint(datapoint))
	}

	// Create the initial medoids
	initialMedoidsCreated := 0
	for {
		possibleMedoid := rand.Intn(len(data))
		if km.Datapoints[possibleMedoid].IsMedoid == true {
			continue
		}

		km.Datapoints[possibleMedoid].IsMedoid = true
		km.Medoids[initialMedoidsCreated] = possibleMedoid
		initialMedoidsCreated++
		if initialMedoidsCreated >= km.NumClusters {
			break
		}

	}

	fmt.Printf("Initial medoids: %v\n", km.Medoids)

	var currentCost float64
	var previousCost float64
	var previousPointAssignments []int

	for {

		// Assign each remaining object to the cluster with the nearest representative object
		previousCost = km.assign()

		// Check for convergence
		if reflect.DeepEqual(previousPointAssignments, km.PointsAssignments) {
			break
		}

		// For each medoid object O
		for _, medoid := range km.Medoids {
			fmt.Printf("Evaluating medoid: %v\n", medoid)
			// Randomly select a non medoid object O_random
			var swapMedoid int
			for {
				possibleMedoid := rand.Intn(len(data))
				if km.Datapoints[possibleMedoid].IsMedoid == true {
					continue
				}

				swapMedoid = possibleMedoid
				break
			}
			fmt.Printf("Current assignments %v\n", km.PointsAssignments)

			//  Compute the total cost S of swapping medoid 0 with O_random
			previousPointAssignments = km.PointsAssignments
			currentCost = km.calculateCost(medoid, swapMedoid)

			// if the new cost is greater than the previous one, revert the swapping
			if currentCost > previousCost {
				//fmt.Printf("Swapping reverted\n")
				km.PointsAssignments = previousPointAssignments
			}

		}

	}

	fmt.Printf("Final assignments: %v\n", km.PointsAssignments)
	return nil
}

// assign each data point to its nearest medoid
func (km *KMedoids) assign() float64 {
	var currentTotalCost float64
	for i, datapoint := range km.Datapoints {
		currentMedoid := -1
		currentDistance := math.MaxFloat64

		for i, medoid := range km.Datapoints {
			if medoid.IsMedoid == false {
				continue
			}

			distance, err := distances.Distance(datapoint.Data, medoid.Data, km.DistanceMetric)
			if err != nil {
				continue
			}

			if distance < currentDistance {
				currentDistance = distance
				currentMedoid = i
			}
		}

		// At this point we know to which medoid a datapoint i is associated
		km.PointsAssignments[i] = currentMedoid
		currentTotalCost += currentDistance
		fmt.Printf("Data point %v assign to medoid %v distance %v total current cost: %v\n", i, currentMedoid, currentDistance, currentTotalCost)
	}

	return currentTotalCost
}

func (km *KMedoids) calculateCost(originalMedoid, swapMedoid int) float64 {
	fmt.Printf("Original medoid: %v Swapping medoid: %v\n", originalMedoid, swapMedoid)
	var totalCost float64
	for i, datapoint := range km.Datapoints {
		var currentCost float64
		// If the current point is assigned to the medoid we are examining
		// we calculate its distance to the new swap alternative
		if km.PointsAssignments[i] == originalMedoid {
			// TODO: Catch this error don't be lazy :)
			currentCost, _ = distances.Distance(datapoint.Data, km.Datapoints[swapMedoid].Data, km.DistanceMetric)
			km.PointsAssignments[i] = swapMedoid
		} else {
			// Otherwise calculate the distance of the datapoint to its previous medoid
			currentCost, _ = distances.Distance(datapoint.Data, km.Datapoints[km.PointsAssignments[i]].Data, km.DistanceMetric)
		}
		totalCost += currentCost
	}

	return totalCost
}
