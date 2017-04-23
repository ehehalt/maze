package maze

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	runeWall  rune = '*'
	runeFree  rune = ' '
	runeStart rune = 'S'
	runeEnd   rune = 'L'
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Grid struct {
	Rows, Columns int
	Cells         [][]*Cell
	Dists         *Dist
	Paths         *Dist
}

type Generator func(Grid) *Grid

// prepareGrid creates the multidimensional Cell slice
func (g *Grid) prepareCells() {
	g.Cells = make([][]*Cell, g.Rows)
	for i := range g.Cells {
		g.Cells[i] = make([]*Cell, g.Columns)
	}
}

// configureCells connect the neighbors
func (g *Grid) configureCells() {
	for row, cells := range g.Cells {
		for col, _ := range cells {
			g.Cells[row][col] = NewCell(row, col)
			if row > 0 {
				g.Cells[row][col].North = g.Cells[row-1][col]
				g.Cells[row-1][col].South = g.Cells[row][col]
			}
			if col > 0 {
				g.Cells[row][col-1].East = g.Cells[row][col]
				g.Cells[row][col].West = g.Cells[row][col-1]
			}
		}
	}
}

// NewGrid creates a new Grid
func NewGrid(rows, columns int) *Grid {
	g := &Grid{Rows: rows, Columns: columns}
	g.prepareCells()
	g.configureCells()
	return g
}

// CellAt returns a cell at a given position or nil if the position is outside the grid
func (g *Grid) CellAt(row, col int) *Cell {
	if row >= 0 && row < g.Rows {
		if col >= 0 && col < g.Columns {
			return g.Cells[row][col]
		}
	}
	return nil
}

// RandomCell returns a random cell from the grid
// Same as method Sample()?!
func (g *Grid) RandomCell() *Cell {
	row := rand.Intn(g.Rows)
	col := rand.Intn(g.Columns)
	return g.CellAt(row, col)
}

// Size returns the amount of cells in a grid
func (g *Grid) Size() int {
	return g.Rows * g.Columns
}

// Converts the multidimensional grid into a one dimensional slice
func (g *Grid) EachCell() CellSlice {
	cells := []*Cell{}
	for _, row := range g.Cells {
		for _, cell := range row {
			cells = append(cells, cell)
		}
	}
	return cells
}

// Sample returns a random Cell from a grid
// Same as method RandomCell()?!
func (g *Grid) Sample() *Cell {
	return g.EachCell().Sample()
}

func (g *Grid) contentsOf(cell *Cell) string {
	return " "
}

func (g *Grid) contentsOfDistance(cell *Cell) string {
	if g.Dists != nil {
		if dist, ok := g.Dists.Cells[cell]; ok {
			return strconv.Itoa(dist)
		}
	}
	return " "
}

// DeadEnds returns a slice with DeadEnd cells.
// DeadEnd cells are cells with only one linked cell.
func (g *Grid) DeadEnds() CellSlice {
	return g.EachCell().DeadEnds()
}

// ToString returns the grid as string.
func (g *Grid) ToString() string {
	output := "+" + strings.Repeat("---+", g.Columns) + "\n"
	for _, row := range g.Cells {
		top := "|"
		bottom := "+"
		for _, cell := range row {
			body := " " + g.contentsOfDistance(cell) + " "
			var eastBoundary string
			if cell.IsLinked(cell.East) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}
			top = top + body + eastBoundary
			var southBoundary string
			if cell.IsLinked(cell.South) {
				southBoundary = "   "
			} else {
				southBoundary = "---"
			}
			corner := "+"
			bottom = bottom + southBoundary + corner
		}
		output = output + top + "\n"
		output = output + bottom + "\n"
	}
	return output
}

// ToRunes converts the grid into a 2-dimensional rune slice.
// This is needed from the maze-box sample.
func (g *Grid) ToRunes() [][]rune {
	rCols := g.Columns*2 + 1
	rRows := g.Rows*2 + 1
	wall := runeWall
	free := runeFree
	strt := runeStart
	ende := runeEnd

	maze := make([][]rune, rRows)
	for row := range maze {
		maze[row] = make([]rune, rCols)
		for i := 0; i < rCols; i++ {
			switch i {
			case 0, rCols - 1:
				maze[row][i] = wall
			default:
				if row == 0 || row == rRows-1 {
					maze[row][i] = wall
				} else {
					maze[row][i] = free
				}
			}
		}
	}

	for _, gRow := range g.Cells {
		for _, cell := range gRow {
			pRow := cell.Row*2 + 1
			pCol := cell.Column*2 + 1

			if cell.IsLinked(cell.East) {
				maze[pRow][pCol+1] = free
			} else {
				maze[pRow][pCol+1] = wall
			}

			if cell.IsLinked(cell.West) {
				maze[pRow][pCol-1] = free
			} else {
				maze[pRow][pCol-1] = wall
			}

			if cell.IsLinked(cell.North) {
				maze[pRow-1][pCol] = free
			} else {
				maze[pRow-1][pCol] = wall
			}

			if cell.IsLinked(cell.South) {
				maze[pRow+1][pCol] = free
			} else {
				maze[pRow+1][pCol] = wall
			}

			maze[pRow-1][pCol-1] = wall
			maze[pRow+1][pCol-1] = wall
			maze[pRow][pCol] = free
		}
	}

	maze[1][1] = strt
	maze[rRows-2][rCols-2] = ende

	/*
		for row := range maze {
			for col := range maze[row] {
				fmt.Print(string(maze[row][col]))
			}
			fmt.Println("")
		}
	*/

	return maze
}
