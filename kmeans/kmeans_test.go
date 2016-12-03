package kmeans

import (
	"testing"

	"github.com/facebookgo/ensure"
)

func TestKMeans(t *testing.T) {
	kmeans := NewKMeans(3, 3, false)
	data := [][]float64{[]float64{0, 0, 0},
		[]float64{2, 2, 2}}
	_, err := kmeans.Fit(data)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, kmeans.PointsAssignment, []int{0, 0})
	ensure.DeepEqual(t, kmeans.Centroids[0].Position, []float64{2.0, 2.0, 2.0})
}

func TestClusterDistanceMatrix(t *testing.T) {
	kmeans := NewKMeans(3, 2, false)
	ab := [][]float64{[]float64{0, 0},
		[]float64{2, 2},
		[]float64{3.0, 3.0},
		[]float64{3.0, 3.0}}
	data := ab
	_, err := kmeans.Fit(data)
	ensure.Nil(t, err)

	cdm := kmeans.ClusterDistance

	ensure.True(t, cdm[0][0] == 0.0)
	ensure.True(t, cdm[1][1] == 0.0)
	ensure.True(t, cdm[2][2] == 0.0)
}

func TestInterclusterDistance(t *testing.T) {
	kmeans := NewKMeans(3, 2, false)
	data := [][]float64{[]float64{0, 0},
		[]float64{2, 2},
		[]float64{3.0, 3.0},
		[]float64{3.0, 3.0}}
	result, err := kmeans.Fit(data)
	ensure.Nil(t, err)

	icd := result.InterclusterDistance
	ensure.True(t, icd > 11.31 && icd < 11.32)
}
