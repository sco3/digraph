package main

import (
	"fmt"
	"sco3/digraph/v2/digraph"
)

func main() {

	dg := digraph.NewDirectedGraph()
	dg.AddEdge('a', 'b')
	dg.AddEdge('a', 'c')

	fmt.Printf("Graph: %s", dg)

}
