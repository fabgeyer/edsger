package edsger

import "errors"

// Implementation using Kahn's algorithm
func (g *Graph[T, N]) TopologicalOrdering() ([]T, error) {
	if !g.directed {
		return nil, errors.New("Graph is not directed")
	}

	allPredecessors := g.AllPredecessors()
	allSuccessors := g.AllSuccessors()
	removeEdge := func(src, dst T) {
		delete(allPredecessors[dst], src)
		if len(allPredecessors[dst]) == 0 {
			delete(allPredecessors, dst)
		}

		delete(allSuccessors[src], dst)
		if len(allSuccessors[src]) == 0 {
			delete(allSuccessors, src)
		}
	}

	// Set of all nodes with no incoming edge
	nodesWithoutPredecessors := make([]T, 0, len(g.nodes))
	g.ApplyNodes(func(n T) bool {
		if len(allPredecessors[n]) == 0 {
			nodesWithoutPredecessors = append(nodesWithoutPredecessors, n)
		}
		return true
	})

	var n T
	res := make([]T, 0, len(g.nodes))
	for len(nodesWithoutPredecessors) > 0 {
		n, nodesWithoutPredecessors = nodesWithoutPredecessors[0], nodesWithoutPredecessors[1:]
		res = append(res, n)

		for m := range allSuccessors[n] {
			removeEdge(n, m)
			if len(allPredecessors[m]) == 0 {
				// m has no other incoming edges
				nodesWithoutPredecessors = append(nodesWithoutPredecessors, m)
			}
		}
	}

	if len(allSuccessors) > 0 {
		return nil, errors.New("Graph contains at least one cycle")
	}
	return res, nil
}

func (g *Graph[T, N]) IsDAG() bool {
	_, err := g.TopologicalOrdering()
	return err == nil
}
