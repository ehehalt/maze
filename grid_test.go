package maze

import (
	"fmt"
	"testing"
)

// Check the Rows and Columns after creating a new grid
func TestNewGrid(t *testing.T) {
	g := NewGrid(2, 4)

	if g.Rows != 2 {
		t.Errorf("g.Rows should return 2, get %d", g.Rows)
	}

	if g.Columns != 4 {
		t.Errorf("g.Columns should return 4, get %d", g.Columns)
	}
}

// Check the dimension of the inner grid after create one
func TestPrepareGrid(t *testing.T) {
	g := NewGrid(2, 4)

	if len(g.Cells) != 2 {
		t.Errorf("g.Cells should have 2 rows, get %d", len(g.Cells))
	}

	for idx, cells := range g.Cells {
		if len(cells) != 4 {
			t.Errorf("g.Cells[%d] should have 4 columns, get %d", idx, len(cells))
		}
	}
}

// Check the cells in the north/south/east/west after creating
func TestConfigureCells(t *testing.T) {
	g := NewGrid(2, 4)

	// Cell[0][0] ...

	if g.Cells[0][0].North != nil {
		t.Error("g.Cells[0][0].North should be nil")
	}

	if g.Cells[0][0].East != g.Cells[0][1] {
		t.Error("g.Cells[0][0].West should be g.Cells[0][1]")
	}

	if g.Cells[0][0].South != g.Cells[1][0] {
		t.Error("g.Cells[0][0].South should be g.Cells[1][0]")
	}

	if g.Cells[0][0].West != nil {
		t.Error("g.Cells[0][0].North should be nil")
	}

	// Cell [1][1]

	if g.Cells[1][1].North != g.Cells[0][1] {
		t.Error("g.Cells[1][1].North should be g.Cells[0][1]")
	}

	if g.Cells[1][1].East != g.Cells[1][2] {
		t.Error("g.Cells[1][1].East should be g.Cells[1][2]")
	}

	if g.Cells[1][1].South != nil {
		t.Error("g.Cells[1][1].South should be nil")
	}

	if g.Cells[1][1].West != g.Cells[1][0] {
		t.Error("g.Cells[1][1].West should be g.Cells[1][0]")
	}
}

func TestCellAt(t *testing.T) {
	g := NewGrid(2, 4)

	if g.CellAt(2, 4) != nil {
		t.Error("g.CellAt(2, 4) should be nil")
	}

	a := g.CellAt(0, 0)
	b := g.Cells[0][0]

	if a != b {
		fmt.Printf("%v", g)
		t.Errorf("g.CellAt(0, 0) %v should be g.Cells[0][0] %v", a, b)
	}
}

func TestRandomCell(t *testing.T) {
	g := NewGrid(2, 4)
	c := g.RandomCell()

	if c == nil {
		t.Error("g.RandomCell() should return a cell and not nil")
	}
}

func TestSize(t *testing.T) {
	g := NewGrid(2, 4)

	if (2 * 4) != g.Size() {
		t.Errorf("g.Size() should return 8, get %d", g.Size())
	}
}

func TestEachCell(t *testing.T) {
	g := NewGrid(2, 4)

	if (2 * 4) != len(g.EachCell()) {
		t.Errorf("g.EachCell() should return 8 cells, get %d", len(g.EachCell()))
	}
}

func TestToString(t *testing.T) {
	g := NewGrid(2, 2)
	s := g.ToString()
	if s != "+---+---+\n|   |   |\n+---+---+\n|   |   |\n+---+---+\n" {
		t.Errorf("g.ToString() result ist wrong!")
	}
}

func TestToRunes(t *testing.T) {
	grid := NewGrid(2, 2)
	runes := grid.ToRunes()
	for lineIdx, line := range runes {
		for runeIdx, rune := range line {

			if lineIdx == 0 || lineIdx == 2 || lineIdx == 4 {
				if rune != runeWall {
					t.Errorf("rune should be '%s', found '%s' at (%d,%d)", string(runeWall), string(rune), lineIdx, runeIdx)
				}
			} else {
				if runeIdx == 0 || runeIdx == 2 || runeIdx == 4 {
					if rune != runeWall {
						t.Errorf("rune should be '%s', found '%s' at (%d,%d)", string(runeWall), string(rune), lineIdx, runeIdx)
					}
				}
			}

			if lineIdx == 1 && runeIdx == 1 {
				if rune != runeStart {
					t.Errorf("rune should be '%s', found '%s' at (%d,%d)", string(runeStart), string(rune), lineIdx, runeIdx)
				}
			}

			if lineIdx == 3 && runeIdx == 3 {
				if rune != runeEnd {
					t.Errorf("rune should be '%s', found '%s' at (%d,%d)", string(runeEnd), string(rune), lineIdx, runeIdx)
				}
			}
		}
	}
}

func TestSample(t *testing.T) {
	g := NewGrid(2, 4)
	c := g.Sample()
	if c == nil {
		t.Errorf("Col c should not be nil")
	}
}

func TestGridDeadEnds(t *testing.T) {
	g := NewGrid(2, 4)
	if len(g.DeadEnds()) != 0 {
		t.Errorf("There should be 0 dead ends!")
	}

	c1 := g.CellAt(0, 0)
	c2 := g.CellAt(0, 1)

	c1.Link(c2)

	if len(g.DeadEnds()) != 2 {
		t.Errorf("There should be 2 dead ends!")
	}
}
