package maze

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Wilson implementation
func Wilson(g Grid) *Grid {
	unvisited := g.EachCell()
	first := unvisited.Sample()
	unvisited = unvisited.Delete(first)

	for unvisited.Any() {
		cell := unvisited.Sample()
		path := CellSlice{cell}

		for unvisited.Include(cell) {
			cell = cell.Neighbors().Sample()
			position := path.Index(cell)
			if position >= 0 {
				path = path[0 : position+1]
			} else {
				path = append(path, cell)
			}
		}

		for idx := 0; idx < len(path)-1; idx++ {
			path[idx].Link(path[idx+1])
			unvisited = unvisited.Delete(path[idx])
		}
	}

	return &g
}
