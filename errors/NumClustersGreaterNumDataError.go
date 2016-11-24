package errors

import "fmt"

// NumClustersGreaterNumDataError is an error triggered when the number
// of clusters assigned is greater than the size of the dataset
type NumClustersGreaterNumDataError struct {
	s string
}

// NewNumClustersGreaterNumDataError creates a new NumClustersGreaterNumDataError
// error.
func NewNumClustersGreaterNumDataError(NumClusters, NumFeatures int) error {
	return &NumClustersGreaterNumDataError{
		fmt.Sprintf(`Number of clusters is greater than the size of the dataset.
			Number of clusters: %v, size of dataset: %v`, NumClusters, NumFeatures)}
}

func (e *NumClustersGreaterNumDataError) Error() string {
	return e.s
}
