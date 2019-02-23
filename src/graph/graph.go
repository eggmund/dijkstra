package graph

import (
	"math"
)

type Graph struct {
	Nodes []*Node
}

func (g *Graph) resetDistances() {
	for i := 0; i < len(g.Nodes); i++ {
		g.Nodes[i].Dist = math.Inf(1)
	}
}

func (g *Graph) Dijkstra() []*Node {
	out := make([]*Node, len(g.Nodes))

	q := make([]*Node, len(g.Nodes))
	copy(q, g.Nodes)

	var prev []*Node

	var u *Node
	for len(q) > 0 {
		u, q = q[0], q[1:]

		for _, v := range u.Cons {
			alt := u.Dist + v.Weight
			if alt < v.Other.Dist {
				v.Other.Dist = alt
				prev = append(prev, u)
			}
		}
	}

	copy(out, g.Nodes)
	g.resetDistances()
	return out
}
