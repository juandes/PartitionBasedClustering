package kmedoids

import (
	"testing"

	"github.com/facebookgo/ensure"
)

func TestKMedoids(t *testing.T) {
	data := [][]float64{[]float64{0, 0},
		[]float64{2, 2}, []float64{3, 4}}
	kmedoids := NewKMedoids(2, 2)
	err := kmedoids.Fit(data)
	ensure.Nil(t, err)
}
