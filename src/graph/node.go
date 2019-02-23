package graph

type Node struct {
	ID string
	Cons []Conn
	Dist float64
}

type Conn struct {
	Other *Node
	Weight float64
}
