package main

import (
	"fmt"
	"sco3/digraph/v2/digraph"
)

func main() {

	dg := digraph.NewDirectedGraph()
	dg.AddEdge('a', 'b')
	dg.AddEdge('a', 'c')

	fmt.Printf("Graph: %s]\n", dg)
	fmt.Printf("Neighbours: %v\n", dg.Neighbors('a'))
	dg.Traverse('a')

}
