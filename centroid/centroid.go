package centroid

// Centroid represents a centroid. The members of the struct
// are the position of said centroid, the sum of the position per dimension
// of all the points belonging to the same centroid, and the
// number of times this centroid was the closest one to a giving
// point during the updating process.
type Centroid struct {
	Position      []float64
	DimensionSums []float64
	TimesItWon    int
}

func NewCentroid(dimension int, initialPosition []float64) *Centroid {
	return &Centroid{
		Position:      initialPosition,
		DimensionSums: make([]float64, dimension),
	}
}
