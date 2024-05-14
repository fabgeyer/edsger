package edsger

import (
	"testing"
)

func TestDijkstraWikipediaGraph(t *testing.T) {
	g := WikipediaGraph()
	if _, total := g.DijkstraShortestPath(1, 5); total != 20 {
		t.Fatal("Invalid path")
	}

	t.Log(g.AllDijkstraShortestPathsMap(1, 5))
	t.Log(g.AllShortestPathsNodes(1, 5))
}

func TestDijkstraMultiplePaths(t *testing.T) {
	g := NewUndirectedGraph[int, int]()
	for i := range 5 {
		g.AddNode(i + 1)
	}
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(3, 5, 1)
	g.AddEdge(1, 4, 5)
	g.AddEdge(4, 5, 10)
	g.AddEdge(1, 5, 2)

	if _, total := g.DijkstraShortestPath(1, 5); total != 2 {
		t.Fatal("Invalid path")
	}

	t.Log(g.AllDijkstraShortestPathsMap(1, 5))
	t.Log(g.AllShortestPathsNodes(1, 5))
}

func TestDijkstraHackerrankGraph(t *testing.T) {
	g := HackerrankGraph()

	if _, total := g.DijkstraShortestPath(1, 2); total != 24 {
		t.Fatal("Invalid path")
	}
	if _, total := g.DijkstraShortestPath(1, 3); total != 3 {
		t.Fatal("Invalid path")
	}
	if _, total := g.DijkstraShortestPath(1, 4); total != 15 {
		t.Fatal("Invalid path")
	}
}

func TestDijkstraNoPath(t *testing.T) {
	g := NoPathGraph()
	path, total := g.DijkstraShortestPath(1, 5)
	if len(path) != 0 || total != 0 {
		t.Fatal("Invalid path")
	}
}

func TestDijkstraAllDijkstraDisjointShortestPathsWikipedia(t *testing.T) {
	g := WikipediaGraph()

	it := g.AllDijkstraDisjointShortestPaths(1, 5)
	for it.Next() {
		t.Log(it.Get())
	}
}

func TestDijkstraAllDijkstraDisjointShortestPaths(t *testing.T) {
	g := NewDirectedGraph[int, int]()

	N := 8
	for i := range 2 + N {
		g.AddNode(i)
	}

	for i := range N {
		g.AddEdge(0, i+2, 1)
		g.AddEdge(i+2, 1, 1)
	}

	i := 0
	it := g.AllDijkstraDisjointShortestPaths(0, 1)
	for it.Next() {
		t.Log(it.Get())
		i++
	}

	if i != N {
		t.Fatal("Invalid number of paths")
	}
}
