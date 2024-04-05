package main

// Graph represents a set of vertices connected by edges.
type Graph struct {
	Vertices map[int]*Vertex
}

// Vertex is a node in the graph that stores the int value at that node
// along with a map to the vertices it is connected to via edges.
type Vertex struct {
	Val   int
	Edges map[int]*Edge
}

// Edge represents an edge in the graph and the destination vertex.
type Edge struct {
	Weight int
	Vertex *Vertex
}

func main() {
	g := &Graph{Vertices: map[int]*Vertex{}}
	sn := NewSocialNetwork(g)
	sn.AddUser(1)
	sn.AddUser(2)
	sn.AddUser(3)
	sn.AddUser(4)
	sn.MakeFriends(1, 2)
	sn.MakeFriends(1, 3)
	sn.MakeFriends(3, 2)
	sn.MakeFriends(4, 2)
	sn.DisplayNetwork()
}
