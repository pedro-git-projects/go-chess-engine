package utils

// Coordinate represents an (y,x) chess board coordinate
type Coordinate Pair[string, int]

// NewCoordinate returns a new instance of the
// Coordinate type
func NewCoordinate(y string, x int) Coordinate {
	return Coordinate{y, x}
}
