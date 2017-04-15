package main

import (
	"fmt"

	"github.com/ehehalt/maze"
)

func main() {
	g := maze.NewGrid(10, 10)

	// g = maze.BinaryTree(*g)
	// g = maze.Sidewinder(*g)
	// g = maze.AldousBroder(*g)
	// g = maze.Wilson(*g)
	g = maze.HuntAndKill(*g)

	ascii := g.ToString()

	fmt.Println(ascii)
}
