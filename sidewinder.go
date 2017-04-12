package maze

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Sidewinder implementation ...
func Sidewinder(g Grid) *Grid {

	for _, row := range g.Cells {

		run := CellSlice{}

		for _, cl := range row {
			run = append(run, cl)

			at_eastern_boundary := (cl.East == nil)
			at_northern_boundary := (cl.North == nil)

			should_close_out := at_eastern_boundary || (!at_northern_boundary && rand.Intn(2) == 0)

			if should_close_out {
				member := run[rand.Intn(len(run))]
				if member.North != nil {
					member.Link(member.North)
				}
				run = nil
				run = CellSlice{}
			} else {
				cl.Link(cl.East)
			}
		}
	}
	return &g
}
