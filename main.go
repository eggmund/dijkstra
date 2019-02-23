package main

import (
	"graph"
)

/*
Graph:
A -> (D, 3), (B, 7)
B -> (A, 7), (D, 2), (E, 6), (C, 3)
C -> (B, 3), (E, 1), (D, 4)
D -> (A, 3), (B, 2), (C, 4), (E, 7)
E -> (D, 7), (B, 6), (C, 1)
*/

func main() {
	g := new(graph.Graph)
	for i := 0; i < 5; d3di++ {
		g.Nodes = append(g.Nodes, new(graph.Node))
	}

	g.Nodes[0].ID = "A"
	g.Nodes[0].Cons = []*graph.Conn{graph.Conn{g.Nodes[3], 3},
																	graph.Conn{g.Nodes[1], 7}}
}
