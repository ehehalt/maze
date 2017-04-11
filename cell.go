package maze

import "fmt"

// Cell is part of a Grid
type Cell struct {
	Row, Column              int
	North, South, East, West *Cell
	links                    map[*Cell]bool
}

// NewCell creates a cell
func NewCell(row, column int) *Cell {
	cell := &Cell{Row: row, Column: column, links: make(map[*Cell]bool)}
	return cell
}

// ToString for info
func (c *Cell) ToString() string {
	var nbs = ""
	for _, nb := range c.Neighbors() {
		if len(nbs) > 0 {
			nbs += ","
		}
		nbs += fmt.Sprintf("n[%d,%d]", nb.Row, nb.Column)
	}

	return fmt.Sprintf("c[%d,%d]nb[%s]", c.Row, c.Column, nbs)
}

// Link linkes a cell to another in both directions
func (c *Cell) Link(o *Cell) {
	c.links[o] = true
	if !o.IsLinked(c) {
		o.Link(c)
	}
}

// IsLinked checks if one cell is linked to another
func (c *Cell) IsLinked(o *Cell) bool {
	_, ok := c.links[o]
	return ok
}

// Unlink o from c and c from o
func (c *Cell) Unlink(o *Cell) {
	if c.IsLinked(o) {
		delete(c.links, o)
		if o.IsLinked(c) {
			o.Unlink(c)
		}
	}
}

// Links are the links of the cell c
func (c *Cell) Links() CellSlice {
	keys := make([]*Cell, 0)
	for key, _ := range c.links {
		keys = append(keys, key)
	}
	return keys
}

// Neighbors ...
func (c *Cell) Neighbors() CellSlice {
	neighbors := make([]*Cell, 0)
	if c.North != nil {
		neighbors = append(neighbors, c.North)
	}
	if c.South != nil {
		neighbors = append(neighbors, c.South)
	}
	if c.East != nil {
		neighbors = append(neighbors, c.East)
	}
	if c.West != nil {
		neighbors = append(neighbors, c.West)
	}
	return neighbors
}

// Distances calculates the distance between two cells
func (c *Cell) Distances() *Dist {
	distances := NewDist(c)
	frontier := []*Cell{c}

	for len(frontier) > 0 {
		newFrontier := make([]*Cell, 0)

		for _, cell := range frontier {
			for linked, _ := range cell.links {
				_, ok := distances.Cells[linked]
				if ok {
					continue
				}
				distances.Cells[linked] = distances.Cells[cell] + 1
				newFrontier = append(newFrontier, linked)
			}
		}

		frontier = newFrontier
	}

	return distances
}
