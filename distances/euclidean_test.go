package distances

import (
	"testing"

	"github.com/facebookgo/ensure"
)

func TestEuclidean(t *testing.T) {
	distance, err := EuclideanDistance([]float64{0, 0}, []float64{4, 0})
	ensure.Nil(t, err)
	ensure.True(t, distance == 4.0)

	distance, err = EuclideanDistance([]float64{2, 1}, []float64{2, 3})
	ensure.Nil(t, err)
	ensure.True(t, distance == 2.0)
}
