package edsger

import "testing"

func TestTopologicalOrdering(t *testing.T) {
	g := NewDirectedGraph[int, int]()
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(5)
	g.AddNode(7)
	g.AddNode(8)
	g.AddNode(9)
	g.AddNode(10)
	g.AddNode(11)

	g.AddEdge(5, 11, 0)
	g.AddEdge(11, 2, 0)
	g.AddEdge(11, 9, 0)
	g.AddEdge(11, 10, 0)
	g.AddEdge(7, 11, 0)
	g.AddEdge(7, 8, 0)
	g.AddEdge(8, 9, 0)
	g.AddEdge(3, 8, 0)
	g.AddEdge(3, 10, 0)

	_, err := g.TopologicalOrdering()
	if err != nil {
		t.Fatal("Invalid result")
	}

	g.AddEdge(10, 5, 0)
	_, err = g.TopologicalOrdering()
	if err == nil {
		t.Fatal("Invalid result")
	}
}
