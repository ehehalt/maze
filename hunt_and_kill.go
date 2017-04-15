package maze

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hunt and Kill implementation
func HuntAndKill(g Grid) *Grid {
	current := g.Sample()

	for current != nil {

		// get the unvisited neighbors of the current cell
		unvisited_neighbors := CellSlice{}
		for _, nb := range current.Neighbors() {
			if len(nb.Links()) == 0 {
				unvisited_neighbors = append(unvisited_neighbors, nb)
			}
		}

		if unvisited_neighbors.Any() {
			neighbor := unvisited_neighbors.Sample()
			current.Link(neighbor)
			current = neighbor
		} else {
			current = nil

			for _, cell := range g.EachCell() {

				// get the visited neighbors of the cell
				visited_neighbors := CellSlice{}
				for _, nb := range cell.Neighbors() {
					if len(nb.Links()) != 0 {
						visited_neighbors = append(visited_neighbors, nb)
					}
				}

				if len(cell.Links()) == 0 && visited_neighbors.Any() {
					current = cell

					neighbor := visited_neighbors.Sample()
					current.Link(neighbor)

					break
				}
			}
		}
	}

	return &g
}
