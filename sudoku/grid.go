package sudoku

import (
	"fmt"
	"strings"
	"unicode"
)

type Grid struct {
	cells [81]int
}

func NewGrid(str string) *Grid {
	s := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, str)

	var grid Grid
	for i, c := range s {
		grid.cells[i] = int(c - '0')
	}

	return &grid
}

func CopyGrid(grid *Grid) *Grid {
	var cells [81]int
	copy(cells[:], grid.cells[:])
	return &Grid{cells: cells}
}

func EqualGrids(left, right *Grid) bool {
	return EqualValues(left.cells[:], right.cells[:])
}

func index(i, j int) int {
	return i + j*9
}

func fmtLine(line []int) string {
	return fmt.Sprintf("|%d%d%d|%d%d%d|%d%d%d|", line[0], line[1], line[2], line[3], line[4], line[5], line[6], line[7], line[8])
}

func (grid *Grid) String() string {
	var output string

	output += "-------------\n"
	output += fmtLine(grid.cells[0:9]) + "\n"
	output += fmtLine(grid.cells[9:18]) + "\n"
	output += fmtLine(grid.cells[18:27]) + "\n"
	output += "-------------\n"
	output += fmtLine(grid.cells[27:36]) + "\n"
	output += fmtLine(grid.cells[36:45]) + "\n"
	output += fmtLine(grid.cells[45:54]) + "\n"
	output += "-------------\n"
	output += fmtLine(grid.cells[54:63]) + "\n"
	output += fmtLine(grid.cells[63:72]) + "\n"
	output += fmtLine(grid.cells[72:81]) + "\n"
	output += "-------------\n"

	return output
}

func (grid *Grid) GetValue(c Coord) int {
	return grid.cells[index(c.X-1, c.Y-1)]
}

func (grid *Grid) SetValue(c Coord, value int) {
	grid.cells[index(c.X-1, c.Y-1)] = value
}

func (grid *Grid) GetRowNeighbours(c Coord) Coords {
	result := make(Coords, 0, 8)

	for i := 1; i < 10; i++ {
		if i != c.X {
			result = append(result, Coord{i, c.Y})
		}
	}

	return result
}

func (grid *Grid) GetColumnNeighbours(c Coord) Coords {
	result := make(Coords, 0, 8)

	for i := 1; i < 10; i++ {
		if i != c.Y {
			result = append(result, Coord{c.X, i})
		}
	}

	return result
}

func (grid *Grid) GetCellNeighbours(c Coord) Coords {
	result := make(Coords, 0, 8)

	i_offset := ((c.X - 1) / 3) * 3
	j_offset := ((c.Y - 1) / 3) * 3
	for j := 1 + j_offset; j < 4+j_offset; j++ {
		for i := 1 + i_offset; i < 4+i_offset; i++ {
			if !(i == c.X && j == c.Y) {
				result = append(result, Coord{i, j})
			}
		}
	}

	return result
}

func (grid *Grid) GetCellExclusiveNeighbours(c Coord) Coords {
	result := make(Coords, 0, 8)

	i_offset := ((c.X - 1) / 3) * 3
	j_offset := ((c.Y - 1) / 3) * 3
	for j := 1 + j_offset; j < 4+j_offset; j++ {
		for i := 1 + i_offset; i < 4+i_offset; i++ {
			if i != c.X && j != c.Y {
				result = append(result, Coord{i, j})
			}
		}
	}

	return result
}

func (grid *Grid) GetNeighbours(c Coord) Coords {
	var result Coords
	result = append(result, grid.GetRowNeighbours(c)...)
	result = append(result, grid.GetColumnNeighbours(c)...)
	result = append(result, grid.GetCellExclusiveNeighbours(c)...)
	return result
}

func (grid *Grid) GetNeighboursValuesAsMap(c Coord) map[int]bool {
	neighbours := grid.GetNeighbours(c)
	values := map[int]bool{}

	for _, coord := range neighbours {
		value := grid.GetValue(coord)
		if value != 0 {
			values[value] = true
		}
	}

	return values
}

func (grid *Grid) GetNeighboursValues(c Coord) Values {
	values := grid.GetNeighboursValuesAsMap(c)

	result := make(Values, 0, 9)
	for i := 1; i < 10; i++ {
		if values[i] {
			result = append(result, i)
		}
	}

	return result
}

func (grid *Grid) GetPossibleValues(c Coord) Values {
	values := grid.GetNeighboursValuesAsMap(c)

	result := make(Values, 0, 9)
	for i := 1; i < 10; i++ {
		if grid.GetValue(c) == 0 && values[i] == false {
			result = append(result, i)
		}
	}

	return result
}

func (grid *Grid) GetNextEmptyCell(c Coord) Coord {
	for i := index(c.X-1, c.Y-1); i < 81; i++ {
		if grid.cells[i] == 0 {
			return Coord{i%9 + 1, i/9 + 1}
		}
	}

	return Coord{0, 0}
}

func (grid *Grid) GetNextCell(c Coord) Coord {
	if c.X < 9 {
		return Coord{c.X + 1, c.Y}
	}
	if c.Y < 9 {
		return Coord{1, c.Y + 1}
	}

	return Coord{0, 0}
}
