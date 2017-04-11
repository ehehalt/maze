package maze

import "testing"

func TestNewCell(t *testing.T) {
	c := NewCell(1, 2)

	if c.Row != 1 {
		t.Errorf("Row should returns 1, get %d", c.Row)
	}

	if c.Column != 2 {
		t.Errorf("Column should return 2, get %d", c.Column)
	}
}

func TestLink(t *testing.T) {
	a := NewCell(1, 1)
	b := NewCell(2, 2)

	a.Link(b)

	if !a.IsLinked(b) {
		t.Error("a should be linked with b!")
	}

	if !b.IsLinked(a) {
		t.Error("b should be linked with a!")
	}
}

func TestUnlink(t *testing.T) {
	a := NewCell(1, 1)
	b := NewCell(2, 2)

	a.Link(b)
	b.Unlink(a)

	if a.IsLinked(b) {
		t.Error("a should not linked with b!")
	}

	if b.IsLinked(a) {
		t.Error("b should not linked with a!")
	}
}

func TestLinks(t *testing.T) {
	a := NewCell(1, 1)

	if len(a.Links()) != 0 {
		t.Errorf("a should have zero links, get %d!", len(a.Links()))
	}

	b := NewCell(2, 2)
	a.Link(b)

	if len(a.Links()) != 1 {
		t.Errorf("a should have one link, get %d", len(a.Links()))
	}

	c := NewCell(3, 3)
	c.Link(a)

	if len(a.Links()) != 2 {
		t.Errorf("a should have two links, get %d!", len(a.Links()))
	}
}

func TestNeighbors(t *testing.T) {
	x := NewCell(1, 1)
	n := NewCell(2, 2)
	s := NewCell(3, 3)
	e := NewCell(4, 4)
	w := NewCell(5, 5)

	if len(x.Neighbors()) != 0 {
		t.Errorf("x should have zero neighbors, get %d", len(x.Neighbors()))
	}

	x.North = n
	if len(x.Neighbors()) != 1 {
		t.Errorf("x should have one neighbor, get %d", len(x.Neighbors()))
	}

	x.South = s
	if len(x.Neighbors()) != 2 {
		t.Errorf("x should have two neighbors, get %d", len(x.Neighbors()))
	}

	x.East = e
	if len(x.Neighbors()) != 3 {
		t.Errorf("x should have three neighbors, get %d", len(x.Neighbors()))
	}

	x.West = w
	if len(x.Neighbors()) != 4 {
		t.Errorf("x should have four neighbors, get %d", len(x.Neighbors()))
	}
}

func TestDistRoot(t *testing.T) {
	c := NewCell(1, 1)
	d := NewDist(c)

	if d.Root != c {
		t.Errorf("d.Root should be c")
	}

	if d.Cells[c] != 0 {
		t.Errorf("d.Cells[c] should be 0")
	}
}
