package maze

import "testing"

func TestNewDistance(t *testing.T) {
	c := NewCell(0, 0)
	d := NewDist(c)

	if d.Root != c {
		t.Errorf("d.Root should c")
	}

	if d.Cells[c] != 0 {
		t.Errorf("c should have a distance 0")
	}
}

func TestPathTo(t *testing.T) {
	g := NewGrid(2, 2)
	if g.Size() != 4 {
		t.Errorf("g should have a size of 16, got %d", g.Size())
	}
	c00 := g.CellAt(0, 0)
	c01 := g.CellAt(0, 1)
	c10 := g.CellAt(1, 0)
	c11 := g.CellAt(1, 1)
	c00.Link(c01)
	c01.Link(c11)
	c11.Link(c10)

	d := c00.Distances()
	p := d.PathTo(c11)

	d00 := d.Cells[c00]
	p00 := p.Cells[c00]

	if d00 != p00 {
		t.Errorf("distance in Distances and PathTo for cell(0,0) should the same")
	}
}
