# maze

Maze algorithms from the book [Mazes for Programmers](https://pragprog.com/book/jbmaze/mazes-for-programmers) translated to Go.

Translated algorithms:

- Aldous Broder
- Binary Tree
- Sidewinder
- Wilson

## Samples

The first sample creates a 10x10 maze with the [Aldous Broder Algorithm](http://weblog.jamisbuck.org/2011/1/17/maze-generation-aldous-broder-algorithm) and print out the result:

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
