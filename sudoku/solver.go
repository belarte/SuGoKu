package sudoku

type Solver interface {
	Solve(Grid) bool
}

type RecursiveSolver struct {
	grid *Grid
}

func (solver *RecursiveSolver) Solve(grid *Grid) bool {
	solver.grid = grid
	firstCell := solver.grid.GetNextEmptyCell(Coord{1, 1})
	return solver.solve(firstCell)
}

func (solver *RecursiveSolver) solve(c Coord) bool {
	if EqualCoord(c, Coord{0, 0}) {
		return true
	}

	for _, value := range solver.grid.GetPossibleValues(c) {
		solver.grid.SetValue(c, value)
		res := solver.solve(solver.grid.GetNextEmptyCell(c))
		if res == true {
			return true
		}
	}

	solver.grid.SetValue(c, 0)
	return false
}
