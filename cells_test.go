package maze

import "testing"

// setup ...
func createCells() (CellSlice, *Cell) {
	cs := []*Cell{}
	c1 := NewCell(1, 2)
	c2 := NewCell(2, 3)
	cs = append(cs, c1)
	cs = append(cs, c2)
	return cs, c1
}

func TestCellsContains(t *testing.T) {
	cs, c1 := createCells()
	if cs.Contains(c1) == false {
		t.Errorf("Col c1 should be in the cs slice, get not")
	}
	c3 := NewCell(3, 4)
	if cs.Contains(c3) == true {
		t.Errorf("Col c3 should not be in the cs slice")
	}
}

func TestCellsInclude(t *testing.T) {
	cs, c1 := createCells()
	if cs.Include(c1) == false {
		t.Errorf("Col c1 should be in the cs slice, get not")
	}
	c3 := NewCell(3, 4)
	if cs.Include(c3) == true {
		t.Errorf("Col c3 should not be in the cs slice")
	}
}

func TestCellsDelete(t *testing.T) {
	cs, c1 := createCells()
	cs = cs.Delete(c1)
	for _, n := range cs {
		if n == c1 {
			t.Errorf("Col c1 should not be in the ns slice")
		}
	}
}

func TestCellsSample(t *testing.T) {
	cs, _ := createCells()
	cx := cs.Sample()
	if cx == nil {
		t.Errorf("Col cx should not be nil")
	}
}

func TestCellsAny(t *testing.T) {
	cs, _ := createCells()
	if cs.Any() == false {
		t.Errorf("Any should not be false!")
	}

	cs = CellSlice{}
	if cs.Any() {
		t.Errorf("Any should be true!")
	}
}

func TestIndex(t *testing.T) {
	cs, c1 := createCells()
	if cs.Index(c1) != 0 {
		t.Errorf("Index should be 0")
	}

	c3 := NewCell(0, 0)
	if cs.Index(c3) != -1 {
		t.Error("Index shold be -1")
	}
}

func TestCellSliceDeadEnds(t *testing.T) {
	cs, _ := createCells()
	if len(cs.DeadEnds()) != 0 {
		t.Error("There should be 0 dead ends.")
	}

	c0 := cs[0]
	c1 := cs[1]

	c0.Link(c1)

	if len(cs.DeadEnds()) != 2 {
		t.Error("There should be 2 dead ends.")
	}
}

func TestCellSliceLast(t *testing.T) {
	cs := CellSlice{NewCell(1, 1)}
	cs = append(cs, NewCell(2, 2))
	if cs.Last().Column != 2 {
		t.Error("The column of the last element should be 2")
	}
	if cs.Last().Row != 2 {
		t.Error("The row of the last element should be 2")
	}
}

func TestCellSliceEmpty(t *testing.T) {
	cs := CellSlice{}
	if !cs.Empty() {
		t.Error("The cell slice has to be empty")
	}

	cs = append(cs, NewCell(1, 1))
	if cs.Empty() {
		t.Error("The cell slice should not empty")
	}
}

func TestCellSlicePop(t *testing.T) {
	cs, _ := createCells()
	c := cs.Pop()
	if len(cs) != 1 {
		t.Error("The cell slice should be a size of 1 element")
	}
}
