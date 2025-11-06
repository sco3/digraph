package digraph

import (
	"fmt"
	"strings"
)

func (g *DirectedGraph) String() string {
	var builder strings.Builder
	builder.WriteRune('[')
	for k, v := range g.Adj {
		for _, c := range v {
			if builder.Len() > 1 {
				builder.WriteString(", ")

			}
			builder.WriteString(
				fmt.Sprintf("%s -> %s", string(k), string(c)),
			)
		}
	}
	builder.WriteRune(']')
	return builder.String()

}
