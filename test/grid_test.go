package test

import (
	"testing"

	"github.com/belarte/SuGoKu/sudoku"
)

func TestGridGetRowNeighbours(t *testing.T) {
	var grid sudoku.Grid

	var entries = []struct {
		in       sudoku.Coord
		expected sudoku.Coords
	}{
		{sudoku.Coord{1, 1}, sudoku.Coords{
			sudoku.Coord{2, 1}, sudoku.Coord{3, 1},
			sudoku.Coord{4, 1}, sudoku.Coord{5, 1},
			sudoku.Coord{6, 1}, sudoku.Coord{7, 1},
			sudoku.Coord{8, 1}, sudoku.Coord{9, 1},
		}},
		{sudoku.Coord{4, 7}, sudoku.Coords{
			sudoku.Coord{1, 7}, sudoku.Coord{2, 7},
			sudoku.Coord{3, 7}, sudoku.Coord{5, 7},
			sudoku.Coord{6, 7}, sudoku.Coord{7, 7},
			sudoku.Coord{8, 7}, sudoku.Coord{9, 7},
		}},
		{sudoku.Coord{9, 9}, sudoku.Coords{
			sudoku.Coord{1, 9}, sudoku.Coord{2, 9},
			sudoku.Coord{3, 9}, sudoku.Coord{4, 9},
			sudoku.Coord{5, 9}, sudoku.Coord{6, 9},
			sudoku.Coord{7, 9}, sudoku.Coord{8, 9},
		}},
	}

	for _, entry := range entries {
		val := grid.GetRowNeighbours(entry.in)
		if !sudoku.EqualCoords(val, entry.expected) {
			t.Errorf("\nexpected %s\nbut got  %s", entry.expected, val)
		}
	}
}

func TestGridGetColumnNeighbours(t *testing.T) {
	var grid sudoku.Grid

	var entries = []struct {
		in       sudoku.Coord
		expected sudoku.Coords
	}{
		{sudoku.Coord{1, 1}, sudoku.Coords{
			sudoku.Coord{1, 2}, sudoku.Coord{1, 3},
			sudoku.Coord{1, 4}, sudoku.Coord{1, 5},
			sudoku.Coord{1, 6}, sudoku.Coord{1, 7},
			sudoku.Coord{1, 8}, sudoku.Coord{1, 9},
		}},
		{sudoku.Coord{4, 7}, sudoku.Coords{
			sudoku.Coord{4, 1}, sudoku.Coord{4, 2},
			sudoku.Coord{4, 3}, sudoku.Coord{4, 4},
			sudoku.Coord{4, 5}, sudoku.Coord{4, 6},
			sudoku.Coord{4, 8}, sudoku.Coord{4, 9},
		}},
		{sudoku.Coord{9, 9}, sudoku.Coords{
			sudoku.Coord{9, 1}, sudoku.Coord{9, 2},
			sudoku.Coord{9, 3}, sudoku.Coord{9, 4},
			sudoku.Coord{9, 5}, sudoku.Coord{9, 6},
			sudoku.Coord{9, 7}, sudoku.Coord{9, 8},
		}},
	}

	for _, entry := range entries {
		val := grid.GetColumnNeighbours(entry.in)
		if !sudoku.EqualCoords(val, entry.expected) {
			t.Errorf("\nexpected %s\nbut got  %s", entry.expected, val)
		}
	}
}

func TestGridGetCellNeighbours(t *testing.T) {
	var grid sudoku.Grid

	var entries = []struct {
		in       sudoku.Coord
		expected sudoku.Coords
	}{
		{sudoku.Coord{1, 1}, sudoku.Coords{
			sudoku.Coord{2, 1}, sudoku.Coord{3, 1},
			sudoku.Coord{1, 2}, sudoku.Coord{2, 2},
			sudoku.Coord{3, 2}, sudoku.Coord{1, 3},
			sudoku.Coord{2, 3}, sudoku.Coord{3, 3},
		}},
		{sudoku.Coord{5, 6}, sudoku.Coords{
			sudoku.Coord{4, 4}, sudoku.Coord{5, 4},
			sudoku.Coord{6, 4}, sudoku.Coord{4, 5},
			sudoku.Coord{5, 5}, sudoku.Coord{6, 5},
			sudoku.Coord{4, 6}, sudoku.Coord{6, 6},
		}},
		{sudoku.Coord{9, 9}, sudoku.Coords{
			sudoku.Coord{7, 7}, sudoku.Coord{8, 7},
			sudoku.Coord{9, 7}, sudoku.Coord{7, 8},
			sudoku.Coord{8, 8}, sudoku.Coord{9, 8},
			sudoku.Coord{7, 9}, sudoku.Coord{8, 9},
		}},
	}

	for _, entry := range entries {
		val := grid.GetCellNeighbours(entry.in)
		if !sudoku.EqualCoords(val, entry.expected) {
			t.Errorf("\nexpected %s\nbut got  %s", entry.expected, val)
		}
	}
}

func TestGridGetCellExclusiveNeighbours(t *testing.T) {
	var grid sudoku.Grid

	var entries = []struct {
		in       sudoku.Coord
		expected sudoku.Coords
	}{
		{sudoku.Coord{1, 1}, sudoku.Coords{
			sudoku.Coord{2, 2}, sudoku.Coord{3, 2},
			sudoku.Coord{2, 3}, sudoku.Coord{3, 3},
		}},
		{sudoku.Coord{5, 6}, sudoku.Coords{
			sudoku.Coord{4, 4}, sudoku.Coord{6, 4},
			sudoku.Coord{4, 5}, sudoku.Coord{6, 5},
		}},
		{sudoku.Coord{8, 5}, sudoku.Coords{
			sudoku.Coord{7, 4}, sudoku.Coord{9, 4},
			sudoku.Coord{7, 6}, sudoku.Coord{9, 6},
		}},
		{sudoku.Coord{9, 9}, sudoku.Coords{
			sudoku.Coord{7, 7}, sudoku.Coord{8, 7},
			sudoku.Coord{7, 8}, sudoku.Coord{8, 8},
		}},
	}

	for _, entry := range entries {
		val := grid.GetCellExclusiveNeighbours(entry.in)
		if !sudoku.EqualCoords(val, entry.expected) {
			t.Errorf("\nexpected %s\nbut got  %s", entry.expected, val)
		}
	}
}

func TestGridGetNeighbours(t *testing.T) {
	var grid sudoku.Grid

	var entries = []struct {
		in       sudoku.Coord
		expected sudoku.Coords
	}{
		{sudoku.Coord{4, 7}, sudoku.Coords{
			sudoku.Coord{1, 7}, sudoku.Coord{2, 7},
			sudoku.Coord{3, 7}, sudoku.Coord{5, 7},
			sudoku.Coord{6, 7}, sudoku.Coord{7, 7},
			sudoku.Coord{8, 7}, sudoku.Coord{9, 7},
			sudoku.Coord{4, 1}, sudoku.Coord{4, 2},
			sudoku.Coord{4, 3}, sudoku.Coord{4, 4},
			sudoku.Coord{4, 5}, sudoku.Coord{4, 6},
			sudoku.Coord{4, 8}, sudoku.Coord{4, 9},
			sudoku.Coord{5, 8}, sudoku.Coord{6, 8},
			sudoku.Coord{5, 9}, sudoku.Coord{6, 9},
		}},
	}

	for _, entry := range entries {
		val := grid.GetNeighbours(entry.in)
		if !sudoku.EqualCoords(val, entry.expected) {
			t.Errorf("\nexpected %s\nbut got  %s", entry.expected, val)
		}
	}
}
