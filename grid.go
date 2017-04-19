package maze

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

// DeadEnds ...
func (g *Grid) DeadEnds() CellSlice {
	return g.EachCell().DeadEnds()
}

// ToString()  ...
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

func (g *Grid) ToRunes() [][]rune {
	cols := g.Columns*2 + 1
	rows := g.Rows*2 + 1
	wall := '*'
	free := ' '

	maze := make([][]rune, rows)
	for row := range maze {
		maze[row] = make([]rune, cols)
		for i := 0; i < cols; i++ {
			switch i {
			case 0, cols - 1:
				maze[row][i] = wall
			default:
				if row == 0 || row == rows-1 {
					maze[row][i] = wall
				} else {
					maze[row][i] = free
				}
			}
		}
	}

	for row := range maze {
		for col := range maze[row] {
			fmt.Print(string(maze[row][col]))
		}
		fmt.Println("")
	}

	/*

		// maze to rune

		maze[last.x][last.y] = 'L'
		bordered := make([][]rune, len(maze)+2)
		for r := range bordered {
			bordered[r] = make([]rune, len(maze[0])+2)
			for c := range bordered[r] {
				if r == 0 || r == len(maze)+1 || c == 0 || c == len(maze[0])+1 {
					bordered[r][c] = '*'
				} else {
					bordered[r][c] = maze[r-1][c-1]
				}
			}
		}
		return bordered
	*/
	return maze
}
