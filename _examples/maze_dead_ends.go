package main

import (
	"fmt"

	"github.com/ehehalt/maze"
)

var (
	names = []string{
		"BinaryTree",
		"Sidewinder",
		"AldousBroder",
		"Wilson",
		"HuntAndKill",
	}

	algorithms = []maze.Generator{
		maze.BinaryTree,
		maze.Sidewinder,
		maze.AldousBroder,
		maze.Wilson,
		maze.HuntAndKill,
	}
)

func main() {

	for idx, _ := range names {

		name := names[idx]
		algorithm := algorithms[idx]

		source := maze.NewGrid(10, 10)
		destination := algorithm(*source)

		fmt.Printf("%20s: DeadEnds = %d\n", name, len(destination.DeadEnds()))
	}

}
