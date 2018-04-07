package sudoku

import (
	"bytes"
	"fmt"
)

type Coord struct {
	X, Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("{%d, %d}", c.X, c.Y)
}

type Coords []Coord

func (c Coords) String() string {
	if c == nil {
		return "nil"
	}

	var buffer bytes.Buffer
	buffer.WriteString("{")
	for _, coord := range c {
		buffer.WriteString(coord.String() + " ")
	}
	buffer.WriteString("}")

	return buffer.String()
}

func EqualCoords(left, right Coords) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

type Values []int

func EqualValues(left, right Values) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}
