package kmeans

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"

	"log"
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
	ClusterDistance  [][]float64
	initRandom       bool
}

// Result is made of the results of the clustering
type Result struct {
	NumClusters      int
	PointsAssignment []int

	// Cluster validation
	// InterclusterDistance is the distance between all the clusters
	InterclusterDistance float64
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

func (km *KMeans) newResult() *Result {
	return &Result{
		NumClusters:          km.NumClusters,
		PointsAssignment:     km.PointsAssignment,
		InterclusterDistance: km.interclusterDistance(),
	}
}

// Fit trains the KMeans model
func (km *KMeans) Fit(dataset [][]float64) (*Result, error) {
	data := make([][]float64, len(dataset))
	// Copy the array
	for i := range dataset {
		for j := range dataset[i] {
			data[i] = append(data[i], dataset[i][j])
		}
	}

	previousAssignments := make([]int, len(data))

	// Use a random point from the original dataset as centroid
	for i := 0; i < km.NumClusters; i++ {
		km.Centroids[i].Position = data[rand.Intn(len(data))]
	}

	if len(data) < 1 {
		return nil, fmt.Errorf("Dataset is empty")
	}

	if len(data[0]) != km.NumFeatures {
		return nil, fmt.Errorf("Number of dimensions is distinct than the number of features")
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

	km.ClusterDistance = km.clustersDistanceMatrix(dataset)

	return km.newResult(), nil
}

// Assigns each datapoint to a cluster
func (km *KMeans) assign(data [][]float64) {
	km.PointsAssignment = make([]int, len(data))

	// Step 1: Assign each point to its nearest centroid
	for i, entry := range data {
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

}

func (km *KMeans) updateCentroids() {
	for _, centroid := range km.Centroids {
		for j := 0; j < km.NumFeatures; j++ {
			// This occurence means that the centroid has never the closest to any of the datapoints
			if centroid.TimesItWon != 0 {
				centroid.Position[j] = centroid.DimensionSums[j] / float64(centroid.TimesItWon)
			}
		}
	}
}

// This function returns a matrix with the distances between clusters.
// It is calculated by getting the distance between all the points that belongs
// to different clusters
func (km *KMeans) clustersDistanceMatrix(data [][]float64) [][]float64 {
	// Create empty matrix
	clusterDistanceMatrix := make([][]float64, km.NumClusters)
	for i := range clusterDistanceMatrix {
		clusterDistanceMatrix[i] = make([]float64, km.NumClusters)
	}

	// Calculate the distance between all the points.
	// Since the distance between A and B is the same as B and A, aka dist(A,B) = dist(B,A)
	// we avoid calculating the second case
	// Once the distance is calculated, its result is added to the distance matrix
	for j := 0; j < len(km.PointsAssignment); j++ {
		for k := j; k < len(km.PointsAssignment); k++ {
			// Avoid calculating the distance between the same points, aka dist(A,A)
			if j == k || km.PointsAssignment[j] == km.PointsAssignment[k] {
				continue
			}

			distance, err := distances.EuclideanDistance(data[j], data[k])
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			// Because dist(A,B) = dist(B,A). The resulting matrix is symmetric with 0
			// along the diagonal
			clusterDistanceMatrix[km.PointsAssignment[k]][km.PointsAssignment[j]] += distance
			clusterDistanceMatrix[km.PointsAssignment[j]][km.PointsAssignment[k]] += distance

		}
	}
	return clusterDistanceMatrix
}

// Calculate the intercluster distance of the model.
// Intercluster distance is the distance between the clusters
func (km *KMeans) interclusterDistance() float64 {
	sum := 0.0
	// Since the matrix and symmetric, and the diagonal is 0,
	// only the lower diagonal part is used for calculating the
	// intercluster distance
	for i := 1; i < km.NumClusters; i++ {
		for j := 0; j < i; j++ {
			sum += km.ClusterDistance[i][j]
		}
	}
	return sum
}
