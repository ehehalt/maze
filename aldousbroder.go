package maze

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// AldousBroder algorithm ...
func AldousBroder(g Grid) *Grid {
	c := g.RandomCell()
	unvisited := g.Size() - 1

	for unvisited > 0 { // && step < 10 {
		nbs := c.Neighbors()
		rnd := rand.Intn(len(nbs))
		tmp := fmt.Sprintf("rand[0-%d]=%d", len(nbs), rnd)
		fmt.Print(tmp)

		nb := nbs[rnd]

		fmt.Println(":", c.ToString(), " => ", nb.ToString())

		if len(nb.Links()) == 0 {
			c.Link(nb)
			unvisited--
		}

		c = nb
	}

	return &g
}
