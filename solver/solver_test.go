package solver_test

import (
	"testing"

	"github.com/belarte/SuGoKu/solver"
	"github.com/belarte/SuGoKu/sudoku"
)

var grid1 = sudoku.NewGrid("530070000" +
	"600195000" +
	"098000060" +
	"800060003" +
	"400803001" +
	"700020006" +
	"060000280" +
	"000419005" +
	"000080079")

var grid1Solution = sudoku.NewGrid("534678912" +
	"672195348" +
	"198342567" +
	"859761423" +
	"426853791" +
	"713924856" +
	"961537284" +
	"287419635" +
	"345286179")

var grid2 = sudoku.NewGrid("000000000" +
	"000003085" +
	"001020000" +
	"000507000" +
	"004000100" +
	"090000000" +
	"500000073" +
	"002010000" +
	"000040009")

func TestRecursiveSolver(t *testing.T) {
	var entries = []struct {
		in       *sudoku.Grid
		expected bool
		out      *sudoku.Grid
	}{
		{sudoku.CopyGrid(grid1), true, grid1Solution},
	}

	solver := solver.NewRecursiveSolver()
	for _, entry := range entries {
		if entry.expected != solver.Solve(entry.in) {
			t.Errorf("\n%s\n\n%s", entry.in, entry.out)
		}
	}
}

func BenchmarkRecursiveGrid1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solver := solver.NewRecursiveSolver()
		solver.Solve(sudoku.CopyGrid(grid1))
	}
}

func BenchmarkRecursiveGrid2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solver := solver.NewRecursiveSolver()
		solver.Solve(sudoku.CopyGrid(grid2))
	}
}
