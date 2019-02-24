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

func (g *Graph) Dijkstra() ([]*Node, []*Node) {
	q := make([]*Node, len(g.Nodes))
	copy(q, g.Nodes)

	var prev []*Node
	var u *Node
	for len(q) > 0 {
		u, q = getMinDistNode(q)
		for _, v := range u.Cons {
			alt := u.Dist + v.Weight
			if (alt < v.Other.Dist) {
				println("V:", v.Other.ID, "Alt:", alt, "other dist:", v.Other.Dist)
				v.Other.Dist = alt
				prev = append(prev, u)
			}
		}
	}

	return g.Nodes, prev
}

func getMinDistNode(list []*Node) (*Node, []*Node) {
	lowest := new(Node)
	lowest.Dist = math.Inf(1)
	var ind int = -1
	for i := 0; i < len(list); i++ {
		if list[i].Dist < lowest.Dist {
			lowest = list[i]
			ind = i
		}
	}

	if ind > -1 {
		list = append(list[:ind], list[ind+1:]...)
	}
	return lowest, list
}

func nodeIn(list []*Node, node *Node) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == node {
			return true
		}
	}

	return false
}
