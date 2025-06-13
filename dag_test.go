package edsger

import "testing"

func TestTopologicalOrdering(t *testing.T) {
	g := WikipediaDirectedAcyclicGraph()
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
