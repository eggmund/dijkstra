package main

import (
	"fmt"
	"math"
	"graph"
)

/*
Graph:
A/0 -> (D, 3), (B, 7)
B/1 -> (A, 7), (D, 2), (E, 6), (C, 3)
C/2 -> (B, 3), (E, 1), (D, 4)
D/3 -> (A, 3), (B, 2), (C, 4), (E, 7)
E/4 -> (D, 7), (B, 6), (C, 1)
*/
func loadTestGraph() *graph.Graph {
	g := new(graph.Graph)
	for i := 0; i < 5; i++ {
		g.Nodes = append(g.Nodes, new(graph.Node))
		g.Nodes[i].Dist = math.Inf(1)
	}

	g.Nodes[0].ID = "A"
	g.Nodes[0].Cons = []graph.Conn{graph.Conn{Other: g.Nodes[3], Weight: 3},
																 graph.Conn{Other: g.Nodes[1], Weight: 7}}
	g.Nodes[0].Dist = 0

  g.Nodes[1].ID = "B"
	g.Nodes[1].Cons = []graph.Conn{graph.Conn{Other: g.Nodes[0], Weight: 7},
                                 graph.Conn{Other: g.Nodes[3], Weight: 2},
                                 graph.Conn{Other: g.Nodes[4], Weight: 6},
                                 graph.Conn{Other: g.Nodes[2], Weight: 3}}

  g.Nodes[2].ID = "C"
	g.Nodes[2].Cons = []graph.Conn{graph.Conn{Other: g.Nodes[1], Weight: 3},
                                 graph.Conn{Other: g.Nodes[4], Weight: 1},
                                 graph.Conn{Other: g.Nodes[3], Weight: 4}}

  g.Nodes[3].ID = "D"
	g.Nodes[3].Cons = []graph.Conn{graph.Conn{Other: g.Nodes[0], Weight: 3},
                                 graph.Conn{Other: g.Nodes[1], Weight: 2},
                                 graph.Conn{Other: g.Nodes[2], Weight: 4},
                                 graph.Conn{Other: g.Nodes[4], Weight: 7}}

  g.Nodes[4].ID = "E"
	g.Nodes[4].Cons = []graph.Conn{graph.Conn{Other: g.Nodes[3], Weight: 7},
                                 graph.Conn{Other: g.Nodes[1], Weight: 6},
                                 graph.Conn{Other: g.Nodes[2], Weight: 1}}

  return g
}

func main() {
  g := loadTestGraph()
	results := g.Dijkstra(g.Nodes[0], "C")
	println("Done")
	for i := range results {
		fmt.Println(results[i].ID, results[i].Dist)
	}
}
