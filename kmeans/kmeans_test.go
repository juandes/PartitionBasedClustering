package kmeans

import (
	"fmt"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestKMeans(t *testing.T) {
	kmeans := NewKMeans(3, 3, false)
	data := [][]float64{[]float64{0, 0, 0},
		[]float64{2, 2, 2}}
	err := kmeans.Fit(data)
	ensure.Nil(t, err)
	for i, element := range kmeans.PointsAssignment {
		fmt.Printf("%v %v\n", i, element)
	}
}
