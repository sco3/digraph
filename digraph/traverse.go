package digraph

import "fmt"

func (g *DirectedGraph) Traverse(root rune) {
	fmt.Printf("%s", string(root))
}
