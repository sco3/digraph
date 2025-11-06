package digraph

// DirectedGraph represents a simple directed graph using adjacency lists.
type DirectedGraph struct {
	Adj map[rune][]rune
}

// NewDirectedGraph creates and returns a new DirectedGraph.
func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		Adj: make(map[rune][]rune),
	}
}

// AddEdge adds a directed edge from node u to node v.
func (g *DirectedGraph) AddEdge(u, v rune) {
	g.Adj[u] = append(g.Adj[u], v)
}

// Neighbors returns the neighbors of node u.
func (g *DirectedGraph) Neighbors(u rune) []rune {
	return g.Adj[u]
}
