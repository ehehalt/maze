package main

import (
	"fmt"
	"sync"

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

	tries = 100
	size  = 20
)

func main() {

	fmt.Println("Comparing DeadEnds in different algorithms.")
	fmt.Println("Tries per algorithm:", tries)
	fmt.Println("Size of the grids:", size, "x", size)
	fmt.Println()

	var wg sync.WaitGroup
	wg.Add(len(names))

	for idx, _ := range names {
		go func(name string, algorithm maze.Generator) {
			defer wg.Done()
			total_deadends := 0

			for try := 0; try < tries; try++ {
				source := maze.NewGrid(size, size)
				destination := algorithm(*source)

				total_deadends += len(destination.DeadEnds())
			}

			fmt.Printf("%20s = %d\n", name, total_deadends/tries)
		}(names[idx], algorithms[idx])
	}

	wg.Wait()
}
