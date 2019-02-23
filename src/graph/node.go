package graph

type Node struct {
	ID string
	Cons []Conn
}

type Conn struct {
	other *Node
	weight float64
}
