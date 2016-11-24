package distances

import (
	"fmt"
	"math"
)

const (
	// Manhattan represents the Manhattan distance.
	Manhattan = 1.0
	// Euclidean represents the Euclidean distance.
	Euclidean = 2.0
)

// Distance function uses the Minkowski distance approach,
// a generalization of both Manhattan and Euclidean distance.
func Distance(x []float64, y []float64, metric float64) (float64, error) {
	var sum float64

	if len(x) != len(y) {
		return sum, fmt.Errorf("Vectors dimension different: %v, %v", len(x), len(y))
	}

	for index, element := range x {
		sum += math.Pow(math.Abs(element-y[index]), metric)
	}

	return math.Pow(sum, 1/metric), nil
}
