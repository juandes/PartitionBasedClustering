package datapoint

type Datapoint struct {
	Data     []float64
	Centroid int
	IsMedoid bool
}

func NewDataPoint(data []float64) *Datapoint {
	return &Datapoint{
		Data:     data,
		Centroid: -1,
	}
}
