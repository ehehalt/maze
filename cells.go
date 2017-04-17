package maze

import "math/rand"

type CellSlice []*Cell

// Contains checks if a CellSlice contains a Cell
func (cs CellSlice) Contains(cell *Cell) bool {
	for _, current := range cs {
		if current == cell {
			return true
		}
	}
	return false
}

// Include is an alias for Contains
func (cs CellSlice) Include(cell *Cell) bool {
	return cs.Contains(cell)
}

// Delete removes a Cell from a CellSlice
func (cs CellSlice) Delete(cell *Cell) CellSlice {
	ns := []*Cell{}
	for _, current := range cs {
		if current != cell {
			ns = append(ns, current)
		}
	}
	return ns
}

// Sample returns a random Cell from a CellSlice
func (cs CellSlice) Sample() *Cell {
	idx := rand.Intn(len(cs))
	return cs[idx]
}

// Any returns true if the CellSlice is not empty
func (cs CellSlice) Any() bool {
	return len(cs) > 0
}

// Index returns the index of a Cell in a CellSlice or -1 if the CellSlice don't contains the Cell
func (cs CellSlice) Index(cell *Cell) int {
	for idx, current := range cs {
		if current == cell {
			return idx
		}
	}
	return -1
}

func (cs CellSlice) DeadEnds() CellSlice {
	de := CellSlice{}
	for _, current := range cs {
		if len(current.links) == 1 {
			de = append(de, current)
		}
	}
	return de
}
