package main

import "github.com/ehehalt/maze"

func main() {
	g := maze.NewGrid(10, 10)

	// g = maze.BinaryTree(*g)
	// g = maze.Sidewinder(*g)
	g = maze.AldousBroder(*g)
	// g = maze.Wilson(*g)

	maze.GridToPng(g, "maze.png", true, true)
}
