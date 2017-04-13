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
	return &g
}
