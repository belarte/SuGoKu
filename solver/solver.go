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

type node struct {
	coord          sudoku.Coord
	possibleValues sudoku.Values
	currentIndex   int
}

type Moves []node

type IterativeSolver struct {
	grid  *sudoku.Grid
	moves Moves
}

func NewIterativeSolver() *IterativeSolver {
	return &IterativeSolver{}
}

func (solve *IterativeSolver) Solve(grid *sudoku.Grid) bool {
	solve.grid = grid
	solve.moves = Moves{}

	c := sudoku.Coord{0, 1}
	for {
		next := solve.grid.GetNextCell(c)
		c = solve.grid.GetNextEmptyCell(next)
		if sudoku.EqualCoord(next, sudoku.Coord{0, 0}) || sudoku.EqualCoord(c, sudoku.Coord{0, 0}) {
			break
		}

		solve.moves = append(solve.moves, node{c, sudoku.Values{}, 0})
	}

	i := 0
	for i < len(solve.moves) {
		solve.grid.SetValue(solve.moves[i].coord, 0)
		solve.moves[i].possibleValues = solve.grid.GetPossibleValues(solve.moves[i].coord)

		if solve.moves[i].currentIndex < len(solve.moves[i].possibleValues) {
			solve.grid.SetValue(solve.moves[i].coord, solve.moves[i].possibleValues[solve.moves[i].currentIndex])
			solve.moves[i].currentIndex++
			i++
			if i < len(solve.moves) {
				solve.moves[i].currentIndex = 0
			}
		} else {
			solve.grid.SetValue(solve.moves[i].coord, 0)
			i--
		}
	}

	return i == len(solve.moves)
}

type RecursiveParallelSolver struct {
	grid      *sudoku.Grid
	semaphore chan struct{}
}

func NewRecursiveParallelSolver(numChannel int) *RecursiveParallelSolver {
	var sem = make(chan struct{}, numChannel)
	return &RecursiveParallelSolver{grid: &sudoku.Grid{}, semaphore: sem}
}

func (solve *RecursiveParallelSolver) Solve(grid *sudoku.Grid) bool {
	firstCell := grid.GetNextEmptyCell(sudoku.Coord{1, 1})
	res := solve.solve(grid, firstCell)
	*grid = *solve.grid
	return res
}

func (solve *RecursiveParallelSolver) solve(grid *sudoku.Grid, c sudoku.Coord) bool {
	if sudoku.EqualCoord(c, sudoku.Coord{0, 0}) {
		*solve.grid = *grid
		return true
	}

	possibleValues := grid.GetPossibleValues(c)
	outputs := make(chan bool, len(possibleValues))

	for _, value := range possibleValues {
		copy := sudoku.CopyGrid(grid)
		copy.SetValue(c, value)
		select {
		case solve.semaphore <- struct{}{}:
			go func() {
				outputs <- solve.solve(copy, copy.GetNextEmptyCell(c))
				<-solve.semaphore
			}()
		default:
			outputs <- solve.solve(copy, copy.GetNextEmptyCell(c))
		}
	}

	res := false
	for i := 0; i < len(possibleValues); i++ {
		out := <-outputs
		res = res || out
	}

	if res == false {
		grid.SetValue(c, 0)
	}

	return res
}
