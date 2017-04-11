package maze

import (
	"math/rand"
	"time"
)

// init ...
func init() {
	rand.Seed(time.Now().UnixNano())
}

// BinaryTree implementation ...
func BinaryTree(g Grid) *Grid {
	cs := g.EachCell()

	for _, c := range cs {
		neighbors := make(CellSlice, 0)
		if c.North != nil {
			neighbors = append(neighbors, c.North)
		}
		if c.East != nil {
			neighbors = append(neighbors, c.East)
		}

		if len(neighbors) > 0 {
			index := rand.Intn(len(neighbors))
			neighbor := neighbors[index]
			c.Link(neighbor)
		}
	}

	return &g
}
