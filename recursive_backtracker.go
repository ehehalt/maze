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
	return &g
}
