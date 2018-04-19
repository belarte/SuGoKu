package solver

import (
	"github.com/belarte/SuGoKu/sudoku"
)

type Solver interface {
	Solve(sudoku.Grid) bool
}

type RecursiveSolver struct {
	grid *sudoku.Grid
}

func NewRecursiveSolver() *RecursiveSolver {
	return &RecursiveSolver{}
}

func (solve *RecursiveSolver) Solve(grid *sudoku.Grid) bool {
	solve.grid = grid
	firstCell := solve.grid.GetNextEmptyCell(sudoku.Coord{1, 1})
	return solve.solve(firstCell)
}

func (solve *RecursiveSolver) solve(c sudoku.Coord) bool {
	if sudoku.EqualCoord(c, sudoku.Coord{0, 0}) {
		return true
	}

	for _, value := range solve.grid.GetPossibleValues(c) {
		solve.grid.SetValue(c, value)
		res := solve.solve(solve.grid.GetNextEmptyCell(c))
		if res == true {
			return true
		}
	}

	solve.grid.SetValue(c, 0)
	return false
}
