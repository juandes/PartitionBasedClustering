package distances

import (
	"testing"

	"github.com/facebookgo/ensure"
)

func TestMinkowski(t *testing.T) {
	euclideanDistance, err := Distance([]float64{0, 0}, []float64{4, 0}, Euclidean)
	ensure.Nil(t, err)
	ensure.True(t, euclideanDistance == 4.0)

	euclideanDistance, err = Distance([]float64{2, 1}, []float64{2, 3}, Euclidean)
	ensure.Nil(t, err)
	ensure.True(t, euclideanDistance == 2.0)

	euclideanDistance, err = Distance([]float64{3, 0}, []float64{7, 4}, Euclidean)
	ensure.Nil(t, err)
	ensure.True(t, euclideanDistance > 5.65 && euclideanDistance < 5.66)

	manhattanDistance, err := Distance([]float64{0, 0}, []float64{4, 0}, Manhattan)
	ensure.Nil(t, err)
	ensure.True(t, manhattanDistance == 4.0)

	manhattanDistance, err = Distance([]float64{2, 1}, []float64{2, 3}, Manhattan)
	ensure.Nil(t, err)
	ensure.True(t, manhattanDistance == 2.0)

	manhattanDistance, err = Distance([]float64{3, 0}, []float64{7, 4}, Manhattan)
	ensure.Nil(t, err)
	ensure.True(t, manhattanDistance == 8.0)
}
