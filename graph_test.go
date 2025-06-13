package edsger

import (
	"iter"
	"testing"
)

func FullyConnectedGraph(n int) *Graph[int, int] {
	g := NewUndirectedGraph[int, int]()
	for i := range n {
		g.AddNode(i)
	}
	for i := range n {
		for j := range n {
			if i != j {
				if !g.HasEdge(i, j) {
					g.AddEdge(i, j, 1)
				}
			}
		}
	}
	return g
}

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

func WikipediaDirectedAcyclicGraph() *Graph[int, int] {
	// Source: https://en.wikipedia.org/wiki/File:Directed_acyclic_graph_2.svg
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
	return g
}

func WikipediaGraphs() iter.Seq[*Graph[int, int]] {
	return func(yield func(*Graph[int, int]) bool) {
		if !yield(WikipediaGraph()) {
			return
		}
		if !yield(WikipediaDirectedAcyclicGraph()) {
			return
		}
	}
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
		t.Fatalf("Invalid number of nodes: %d", g.NumberOfNodes())
	}
	if g.NumberOfEdges() != 1 {
		t.Fatalf("Invalid number of edges: %v", g.NumberOfEdges())
	}

	if !g.HasEdge(nodes[0], nodes[1]) {
		t.Fatal("Missing edge")
	}
	if !g.HasEdge(nodes[1], nodes[0]) {
		t.Fatal("Missing edge")
	}
	if g.HasEdge(nodes[1], nodes[2]) {
		t.Fatal("Unexpected edge")
	}

	w, ok := g.GetEdge(nodes[0], nodes[1])
	if !ok {
		t.Fatal("Missing edge")
	}
	if w != 1 {
		t.Fatal("Unexpected weight")
	}

	i := 0
	for n, d := range g.Degree() {
		if n != nodes[0] && n != nodes[1] {
			t.Fatal("Invalid node")
		}
		if d != 1 {
			t.Fatal("Invalid degree")
		}
		i++
	}
	if i != 2 {
		t.Fatal("Invalid number of nodes")
	}
}

func TestRemoveNode(t *testing.T) {
	g := WikipediaGraph()
	g.RemoveNode(1)
	if g.NumberOfNodes() != 5 {
		t.Fatalf("Invalid number of nodes: %d", g.NumberOfNodes())
	}
	if g.NumberOfEdges() != 6 {
		t.Fatalf("Invalid number of edges: %v", g.NumberOfEdges())
	}
}

func TestRemoveEdge(t *testing.T) {
	g := WikipediaGraph()
	n := g.NumberOfEdges()

	g.RemoveEdge(1, 6)
	for edge := range g.Edges() {
		t.Log(edge)
		n--
	}

	if n != 1 {
		t.Fatal("Invalid number of edges")
	}
}

func TestEdgeIterator(t *testing.T) {
	for g := range WikipediaGraphs() {
		n := 0
		for edge := range g.Edges() {
			t.Log(edge)
			n++
		}

		if n != g.NumberOfEdges() {
			t.Fatal("Invalid number of edges")
		}
	}
}

func TestNodeIterator(t *testing.T) {
	for g := range WikipediaGraphs() {
		n := 0
		for node := range g.Nodes() {
			t.Log(node)
			n++
		}

		if n != g.NumberOfNodes() {
			t.Fatal("Invalid number of edges")
		}
	}
}

func validateCount[T any](t *testing.T, node T, seq iter.Seq[T], expected int) {
	n := 0
	for range seq {
		n++
	}

	if n != expected {
		t.Fatalf("Invalid count for node %v. Expected %d, received %d", node, expected, n)
	}
}

func TestSuccessors(t *testing.T) {
	for g := range WikipediaGraphs() {
		successors := make(map[int]int, g.NumberOfNodes())
		for e := range g.Edges() {
			successors[e.From]++
			if !g.IsDirected() {
				successors[e.To]++
			}
		}

		for node, c := range successors {
			validateCount(t, node, g.Successors(node), c)
		}
	}
}

func TestPredecessors(t *testing.T) {
	for g := range WikipediaGraphs() {
		predecessors := make(map[int]int, g.NumberOfNodes())
		for e := range g.Edges() {
			predecessors[e.To]++
			if !g.IsDirected() {
				predecessors[e.From]++
			}
		}

		for node, c := range predecessors {
			validateCount(t, node, g.Predecessors(node), c)
		}
	}
}
