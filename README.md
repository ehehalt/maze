# maze

Maze algorithms from the book [Mazes for Programmers](https://pragprog.com/book/jbmaze/mazes-for-programmers) translated to Go.

Translated algorithms:

- Aldous Broder
- Binary Tree
- Sidewinder
- Wilson

## Samples

Create a 10x10 algorithm with Aldous Broder as ASCII output for example:

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
