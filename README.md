# maze

Maze algorithms from the book [Mazes for Programmers](https://pragprog.com/book/jbmaze/mazes-for-programmers) translated to Go.

Translated algorithms:

- Aldous Broder
- Binary Tree
- Sidewinder
- Wilson

## Samples

The first sample creates a 10x10 maze with the *Aldous Broder* algorithm and print the result to the output:

``` go
package main

import (
	"fmt"

	maze "github.com/ehehalt/maze"
)

func main() {
	g := maze.NewGrid(10, 10)
	g = maze.AldousBroder(*g)
	ascii := g.ToString()
	fmt.Println(ascii)
}
```
