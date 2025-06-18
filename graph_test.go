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

// Zachary's Karate Club graph
// https://en.wikipedia.org/wiki/Zachary's_karate_club
// Auto-generated from networkx.karate_club_graph
func KarateClubGraph() *Graph[int, float64] {
	g := NewUndirectedGraph[int, float64]()

	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddNode(6)
	g.AddNode(7)
	g.AddNode(8)
	g.AddNode(9)
	g.AddNode(10)
	g.AddNode(11)
	g.AddNode(12)
	g.AddNode(13)
	g.AddNode(14)
	g.AddNode(15)
	g.AddNode(16)
	g.AddNode(17)
	g.AddNode(18)
	g.AddNode(19)
	g.AddNode(20)
	g.AddNode(21)
	g.AddNode(22)
	g.AddNode(23)
	g.AddNode(24)
	g.AddNode(25)
	g.AddNode(26)
	g.AddNode(27)
	g.AddNode(28)
	g.AddNode(29)
	g.AddNode(30)
	g.AddNode(31)
	g.AddNode(32)
	g.AddNode(33)

	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 5)
	g.AddEdge(0, 3, 3)
	g.AddEdge(0, 4, 3)
	g.AddEdge(0, 5, 3)
	g.AddEdge(0, 6, 3)
	g.AddEdge(0, 7, 2)
	g.AddEdge(0, 8, 2)
	g.AddEdge(0, 10, 2)
	g.AddEdge(0, 11, 3)
	g.AddEdge(0, 12, 1)
	g.AddEdge(0, 13, 3)
	g.AddEdge(0, 17, 2)
	g.AddEdge(0, 19, 2)
	g.AddEdge(0, 21, 2)
	g.AddEdge(0, 31, 2)
	g.AddEdge(1, 2, 6)
	g.AddEdge(1, 3, 3)
	g.AddEdge(1, 7, 4)
	g.AddEdge(1, 13, 5)
	g.AddEdge(1, 17, 1)
	g.AddEdge(1, 19, 2)
	g.AddEdge(1, 21, 2)
	g.AddEdge(1, 30, 2)
	g.AddEdge(2, 3, 3)
	g.AddEdge(2, 7, 4)
	g.AddEdge(2, 8, 5)
	g.AddEdge(2, 9, 1)
	g.AddEdge(2, 13, 3)
	g.AddEdge(2, 27, 2)
	g.AddEdge(2, 28, 2)
	g.AddEdge(2, 32, 2)
	g.AddEdge(3, 7, 3)
	g.AddEdge(3, 12, 3)
	g.AddEdge(3, 13, 3)
	g.AddEdge(4, 6, 2)
	g.AddEdge(4, 10, 3)
	g.AddEdge(5, 6, 5)
	g.AddEdge(5, 10, 3)
	g.AddEdge(5, 16, 3)
	g.AddEdge(6, 16, 3)
	g.AddEdge(8, 30, 3)
	g.AddEdge(8, 32, 3)
	g.AddEdge(8, 33, 4)
	g.AddEdge(9, 33, 2)
	g.AddEdge(13, 33, 3)
	g.AddEdge(14, 32, 3)
	g.AddEdge(14, 33, 2)
	g.AddEdge(15, 32, 3)
	g.AddEdge(15, 33, 4)
	g.AddEdge(18, 32, 1)
	g.AddEdge(18, 33, 2)
	g.AddEdge(19, 33, 1)
	g.AddEdge(20, 32, 3)
	g.AddEdge(20, 33, 1)
	g.AddEdge(22, 32, 2)
	g.AddEdge(22, 33, 3)
	g.AddEdge(23, 25, 5)
	g.AddEdge(23, 27, 4)
	g.AddEdge(23, 29, 3)
	g.AddEdge(23, 32, 5)
	g.AddEdge(23, 33, 4)
	g.AddEdge(24, 25, 2)
	g.AddEdge(24, 27, 3)
	g.AddEdge(24, 31, 2)
	g.AddEdge(25, 31, 7)
	g.AddEdge(26, 29, 4)
	g.AddEdge(26, 33, 2)
	g.AddEdge(27, 33, 4)
	g.AddEdge(28, 31, 2)
	g.AddEdge(28, 33, 2)
	g.AddEdge(29, 32, 4)
	g.AddEdge(29, 33, 2)
	g.AddEdge(30, 32, 3)
	g.AddEdge(30, 33, 3)
	g.AddEdge(31, 32, 4)
	g.AddEdge(31, 33, 4)
	g.AddEdge(32, 33, 5)

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
