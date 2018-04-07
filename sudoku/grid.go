package sudoku

type Grid struct {
	cells [81]int
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
