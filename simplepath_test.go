package edsger

import "testing"

func TestHasSimplePath(t *testing.T) {
	{
		g := WikipediaGraph()
		if g.HasSimplePath(1, 5) != true {
			t.Fatal("Invalid result")
		}
	}

	{
		g := NewUndirectedGraph[int, int]()
		for i := range 5 {
			g.AddNode(i + 1)
		}
		g.AddEdge(1, 2, 1)
		g.AddEdge(1, 3, 1)

		if g.HasSimplePath(1, 5) != false {
			t.Fatal("Invalid result")
		}
	}
}

func TestSimplePath(t *testing.T) {
	g := WikipediaGraph()
	t.Log(g.SimplePath(1, 5))
}

func validateSimplePathIterator[T comparable, N Number](t *testing.T, it *SimplePathIterator[T, N], shortestPath []T, shortestTotal N) {
	t.Log(shortestPath, shortestTotal, "(shortest)")
	source, dest := shortestPath[0], shortestPath[len(shortestPath)-1]

	for it.Next() {
		path, weight := it.Get()
		t.Log(path, weight)
		if path[0] != source {
			t.Fatal("Invalid path")
		}
		if path[len(path)-1] != dest {
			t.Fatal("Invalid path")
		}
		if weight < shortestTotal {
			t.Fatal("Invalid weight")
		}
		if weight > it.CutoffWeight {
			t.Fatal("Invalid weight")
		}
		if len(path) > it.CutoffHops {
			t.Fatal("Invalid path")
		}
	}
	t.Log()
}

func TestAllSimplePaths(t *testing.T) {
	g := WikipediaGraph()
	source, dest := 1, 5

	shortestPath, shortestTotal := g.DijkstraShortestPath(source, dest)

	it := g.AllSimplePaths(source, dest)
	validateSimplePathIterator(t, it, shortestPath, shortestTotal)

	it = g.AllSimplePaths(source, dest)
	it.CutoffWeight = 35
	validateSimplePathIterator(t, it, shortestPath, shortestTotal)

	it = g.AllSimplePaths(source, dest)
	it.CutoffHops = 4
	validateSimplePathIterator(t, it, shortestPath, shortestTotal)

}

func TestAllSimplePathsWithHeuristic(t *testing.T) {
	g := WikipediaGraph()
	source, dest := 1, 5

	distances := make(map[int]int)
	for n := range g.nodes {
		_, distances[n] = g.DijkstraShortestPath(n, dest)
	}
	t.Log(distances)

	it := g.AllSimplePathsWithHeuristic(source, dest, func(i, j NodeWeight[int, int]) int {
		return distances[i.Node] - distances[j.Node]
	})

	for it.Next() {
		path, weight := it.Get()
		t.Log(path, weight)
	}
}
