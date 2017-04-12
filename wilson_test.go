package maze

import (
	"testing"
)

func TestWilson(t *testing.T) {
	grid1 := NewGrid(2, 2)
	if len(grid1.EachCell()) != 4 {
		t.Errorf("grid1 should has 4 cells, get %d cells", len(grid1.EachCell()))
	}

	grid2 := Wilson(*grid1)
	if len(grid2.EachCell()) != 4 {
		t.Errorf("grid2 should has 4 cells, get %d cells", len(grid2.EachCell()))
	}
}
