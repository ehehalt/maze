package maze

// Dist is used for distances between cells
type Dist struct {
	Root  *Cell
	Cells map[*Cell]int
}

// NewDist creates a distance
func NewDist(rootCell *Cell) *Dist {
	dist := &Dist{Root: rootCell, Cells: make(map[*Cell]int)}
	dist.Cells[rootCell] = 0
	return dist
}

// Max returs the longest path in the distances dictionary
func (d *Dist) Max() int {
	max := 0
	for _, v := range d.Cells {
		if v > max {
			max = v
		}
	}
	return max
}

// PathTo ...
func (d *Dist) PathTo(goal *Cell) *Dist {
	current := goal
	breadcrumbs := NewDist(d.Root)
	breadcrumbs.Cells[current] = d.Cells[current]
	for current != d.Root {
		for _, neighbor := range current.Links() {
			if d.Cells[neighbor] < d.Cells[current] {
				breadcrumbs.Cells[neighbor] = d.Cells[neighbor]
				current = neighbor
				break
			}
		}
	}
	return breadcrumbs
}
