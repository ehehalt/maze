package maze

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RecursiveBacktracker algorithm ...
func RecursiveBacktracker(g Grid) *Grid {
	stack := CellSlice{g.Sample()}

	for stack.Any() {
		current := stack.Last()

		neighbors := CellSlice{}
		for _, nb := range current.Neighbors() {
			if nb.Links().Empty() {
				neighbors = append(neighbors, nb)
			}
		}

		if neighbors.Empty() {
			_, stack = stack.Pop()
		} else {
			neighbor := neighbors.Sample()
			current.Link(neighbor)
			stack = append(stack, neighbor)
		}
	}

	return &g
}
