package edsger

import "testing"

func WikipediaGraph() *Graph[int, int] {
	// Source: https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
	g := NewUndirectedGraph[int, int]()
	for i := range 6 {
		g.AddNode(i + 1)
	}
	g.AddEdge(1, 2, 7)
	g.AddEdge(1, 6, 14)
	g.AddEdge(1, 3, 9)
	g.AddEdge(2, 4, 15)
	g.AddEdge(2, 3, 10)
	g.AddEdge(3, 6, 2)
	g.AddEdge(3, 4, 11)
	g.AddEdge(4, 5, 6)
	g.AddEdge(6, 5, 9)
	return g
}

func NoPathGraph() *Graph[int, int] {
	g := NewUndirectedGraph[int, int]()
	for i := range 5 {
		g.AddNode(i + 1)
	}
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	return g
}

func HackerrankGraph() *Graph[int, int] {
	// Source: https://www.hackerrank.com/challenges/dijkstrashortreach/problem
	g := NewUndirectedGraph[int, int]()
	for i := range 4 {
		g.AddNode(i + 1)
	}
	g.AddEdge(1, 2, 24)
	g.AddEdge(1, 4, 20)
	g.AddEdge(3, 1, 3)
	g.AddEdge(4, 3, 12)
	return g
}

func TestGraphStruct(t *testing.T) {
	type Node struct {
		Id int
	}
	g := NewUndirectedGraph[*Node, float64]()

	nodes := make([]*Node, 5)
	for i := range nodes {
		nodes[i] = &Node{i}
		g.AddNode(nodes[i])
	}
	g.AddEdge(nodes[0], nodes[1], 1)

	if g.NumberOfNodes() != 5 {
		t.Fatal("Invalid number of nodes")
	}
	if g.NumberOfEdges() != 1 {
		t.Fatal("Invalid number of edges")
	}
}
