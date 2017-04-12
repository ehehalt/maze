package maze

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// GridToPng creates a png graphic with the given filename.
// If colorize is true, then the graphic is colored.
// If solved is true, the shortest path is marked.
func GridToPng(g *Grid, filename string, colorize bool, solved bool) {

	// size constants
	const (
		sz     = 10
		szhalf = 5
	)

	// Creates the Image as RGBA graphic with the given dynamic size
	img := image.NewRGBA(image.Rect(0, 0, sz*(g.Columns+1), sz*(g.Rows+1)))

	// Creates different colors for the output
	clrRed := color.RGBA{155, 0, 0, 255}
	clrBlue := color.RGBA{0, 0, 255, 255}
	clrGray := color.RGBA{155, 155, 155, 255}

	// Set the default color as clr
	clr := clrRed

	// Set the colors for the solved variant of the picture
	clrSolved := color.RGBA{255, 230, 230, 255}
	clrUnsolved := color.RGBA{255, 255, 255, 255}

	// closure for a horizontal line
	lineHorizontal := func(x, y int) {
		// left: x position
		x1 := ((x + 1) * sz) - szhalf
		// right: x position
		x2 := ((x + 1) * sz) + szhalf
		// y position
		py := ((y + 1) * sz) - szhalf

		// draw the pixels from x1 to x2 @ y
		for px := x1; px <= x2; px++ {
			img.Set(px, py, clr)
		}
	}

	// closure for a vertical line
	lineVertical := func(x, y int) {
		// x position
		px := ((x + 1) * sz) - szhalf
		// up: y1 position
		y1 := ((y + 1) * sz) - szhalf
		// down: y2 position
		y2 := ((y + 1) * sz) + szhalf

		// draw the pixels from y1 to y2 @ x
		for py := y1; py <= y2; py++ {
			img.Set(px, py, clr)
		}
	}

	// closure for a block with size d in both directions
	block := func(x, y, d int, clr color.RGBA) {
		// left: x position
		x = x * sz
		// up: y position
		y = y * sz

		// count from left to right
		for px := (x) - d/2; px <= x+d/2; px++ {
			// count from up to down
			for py := y - d/2; py <= y+d/2; py++ {
				// draw the pixel
				img.Set(px, py, clr)
			}
		}
	}

	// draw the background

	if colorize {
		// the colored version
		startCell := g.CellAt(0, 0)

		// get the distances for the colors
		dists := startCell.Distances()
		if solved {
			// in the solved version get the distances only for the shortest path
			paths := dists.PathTo(g.CellAt(g.Rows-1, g.Columns-1))
			// change the distances to only the shortest path ones
			dists = paths
		}

		// calculate the color delta on the Max distance calculated
		delta := 120.0 / float32(dists.Max()+1)

		// run the rows ...
		for y, row := range g.Cells {
			// ... and the columns
			for x, cell := range row {
				if solved {
					// in the solved variant: only two colors
					clr2 := clrUnsolved
					if _, ok := dists.Cells[cell]; ok {
						clr2 = clrSolved
					}
					// draw the block
					block(x+1, y+1, sz, clr2)
				} else {
					// in the colored variant: calculate the distance
					dist := float32(dists.Cells[cell])
					// the color for the calculated distance
					clr2 := uint8(120.0 + delta*dist)
					// draw the block
					block(x+1, y+1, sz, color.RGBA{255, clr2, clr2, 255})
				}
			}
		}
	} else {
		// the bw version
		if solved {
			// only the solved variant has drawn blocks
			startCell := g.CellAt(0, 0)
			dists := startCell.Distances()
			paths := dists.PathTo(g.CellAt(g.Rows-1, g.Columns-1))
			dists = paths
			for y, row := range g.Cells {
				for x, cell := range row {
					if _, ok := dists.Cells[cell]; ok {
						block(x+1, y+1, 4, clrGray)
					}
				}
			}
		}
	}

	// draw the grid

	for y, row := range g.Cells {
		for x, cell := range row {
			if !cell.IsLinked(cell.East) || cell.East == nil {
				lineVertical(x+1, y)
			}
			if !cell.IsLinked(cell.North) || cell.North == nil {
				lineHorizontal(x, y)
			}
			if !cell.IsLinked(cell.West) || cell.West == nil {
				lineVertical(x, y)
			}
			if !cell.IsLinked(cell.South) || cell.South == nil {
				lineHorizontal(x, y+1)
			}
		}
	}

	// draw the start ...
	block(1, 1, 4, clrRed)

	// ... and end point of the grid
	block(g.Columns, g.Rows, 4, clrBlue)

	// write the image file
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
