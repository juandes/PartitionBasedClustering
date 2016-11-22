package kmeans

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"

	"partitionclustering/centroid"
	"partitionclustering/datapoint"
	"partitionclustering/distances"
)

// KMeans represent a KMeans object.
// It is made of the number of clusters, the number of features,
// the position of the centroids, the data used for performing the clustering
// the assignment of each datapoint to their clusters, and a variable that
// states if we wish to initialize the clusters using a random positions
// or by using the position from some of the rows of the dataset as centroid
type KMeans struct {
	NumClusters      int
	NumFeatures      int
	Centroids        []centroid.Centroid
	Datapoints       []datapoint.Datapoint
	PointsAssignment []int
	initRandom       bool
}

// NewKMeans creates a KMeans object
func NewKMeans(numClusters int, numFeatures int, initRandom bool) *KMeans {
	centroids := make([]centroid.Centroid, numClusters)
	for i := 0; i < numClusters; i++ {
		centroids[i] = *centroid.NewCentroid(numFeatures, []float64{})
	}

	return &KMeans{
		NumClusters: numClusters,
		NumFeatures: numFeatures,
		Centroids:   centroids,
		initRandom:  initRandom,
	}
}

// Fit trains the KMeans model
func (km *KMeans) Fit(data [][]float64) error {
	previousAssignments := make([]int, len(data))

	// Use a random point from the original dataset as centroid
	for i := 0; i < km.NumClusters; i++ {
		km.Centroids[i].Position = data[rand.Intn(len(data))]
	}

	if len(data) < 1 {
		return fmt.Errorf("Dataset is empty")
	}

	if len(data[0]) != km.NumFeatures {
		return fmt.Errorf("Number of dimensions is distinct than the number of features")
	}

	iterations := 0
	for {
		iterations++
		km.assign(data)
		if reflect.DeepEqual(previousAssignments, km.PointsAssignment) {
			break
		}
		previousAssignments = km.PointsAssignment
		km.updateCentroids()
	}

	fmt.Printf("Number of iterations: %v\n", iterations)
	fmt.Printf("Clustering vector: %v\n", km.PointsAssignment)

	for _, val := range km.PointsAssignment {
		fmt.Printf("%v,", val)
	}
	return nil
}

// Assigns each datapoint to a cluster
func (km *KMeans) assign(data [][]float64) {
	km.PointsAssignment = make([]int, len(data))
	//datapoints := make([]datapoint.Datapoint, len(data))

	// Step 1: Assign each point to its nearest centroid
	for i, entry := range data {
		//fmt.Printf("Features: %v\n", features)
		currentCentroid := -1
		currentDistance := math.MaxFloat64

		for j, centroid := range km.Centroids {
			distance, err := distances.EuclideanDistance(entry, centroid.Position)
			if err != nil {
				continue
			}

			if distance < currentDistance {
				currentDistance = distance
				currentCentroid = j
			}
		}

		// Up to this point we know which centroid won, aka. the closest centroid
		km.Centroids[currentCentroid].TimesItWon++

		// Keep the sum of the positions per dimension. This will be
		// divided at a later stage by the number of times the centroid won
		for featureIndex, feature := range entry {
			km.Centroids[currentCentroid].DimensionSums[featureIndex] += feature
		}

		// Up to this point we have calculated the sum of the new features per dimension
		km.PointsAssignment[i] = currentCentroid
	}

	//fmt.Printf("CC: %v\n", km.PointsAssignment)
	//km.Datapoints = datapoints

}

func (km *KMeans) updateCentroids() {
	for _, centroid := range km.Centroids {
		for j := 0; j < km.NumFeatures; j++ {
			centroid.Position[j] = centroid.DimensionSums[j] / float64(centroid.TimesItWon)
		}
	}
}
