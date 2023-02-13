package utils

import "fmt"

// Coordinate represents an (y,x) chess board coordinate
type Coordinate Pair[rune, int]

// NewCoordinate returns a new instance of the
// Coordinate type
func NewCoordinate(y rune, x int) Coordinate {
	return Coordinate{y, x}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%c, %d)", c.First, c.Second)
}
