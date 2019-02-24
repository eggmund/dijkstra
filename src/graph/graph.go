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

func (g *Graph) getOut() []Node {
	var out []Node
	for i := 0; i < len(g.Nodes); i++ {
		out = append(out, *g.Nodes[i])
	}
	return out
}

func (g *Graph) Dijkstra(start *Node, targetID string) []Node {
	start.Dist = 0

	q := make([]*Node, len(g.Nodes))
	copy(q, g.Nodes)

	var u *Node
	for len(q) > 0 {
		u, q = getMinDistNode(q)
		for _, v := range u.Cons {
			alt := u.Dist + v.Weight
			if (alt < v.Other.Dist) {
				v.Other.Dist = alt
			}
		}
	}

	out := g.getOut()
	g.resetDistances()

	return out
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
