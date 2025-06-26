package edsger

import (
	"sync"
	"sync/atomic"
)

// Computes betweenness centrality
// Betweenness centrality of a node $v$ is the sum of all-pairs shortest paths that pass through $v$
func (g *Graph[T, N]) BetweennessCentrality() map[T]int {
	return g.BetweennessCentralitySubset(g.NodesList())
}

// Computes betweenness centrality for a subset of nodes
// Betweenness centrality of a node $v$ is the sum of all-pairs shortest paths that pass through $v$
func (g *Graph[T, N]) BetweennessCentralitySubset(subset []T) map[T]int {
	m := make(map[T]*atomic.Uint64, len(subset))
	// Initialize the map to make sure it can be accessed in parallel
	for _, n := range subset {
		m[n] = new(atomic.Uint64)
	}

	var wg sync.WaitGroup
	wg.Add(len(subset))
	for _, source := range subset {
		go func() {
			defer wg.Done()

			smap := g.sourceShortestPathMap(source, false, nil)
			for _, dest := range subset {
				if source == dest || g.HasEdge(source, dest) {
					continue
				}

				path, _ := pathFromShortestPathMap(dest, smap, 0)
				if len(path) < 3 {
					continue
				}

				for _, node := range path[1 : len(path)-1] {
					m[node].Add(1)
				}
			}
		}()
	}
	wg.Wait()

	return atomicUint64toIntMap(m)
}

// Computes betweenness centrality using all shortest paths
// Betweenness centrality of a node $v$ is the sum of all-pairs all shortest paths that pass through $v$
func (g *Graph[T, N]) AllShortestPathsBetweennessCentrality() map[T]int {
	return g.AllShortestPathsBetweennessCentralitySubset(g.NodesList())
}

// Computes betweenness centrality using all shortest paths
// Betweenness centrality of a node $v$ is the sum of all-pairs all shortest paths that pass through $v$
func (g *Graph[T, N]) AllShortestPathsBetweennessCentralitySubset(subset []T) map[T]int {
	m := make(map[T]*atomic.Uint64, len(subset))
	// Initialize the map to make sure it can be accessed in parallel
	for _, n := range subset {
		m[n] = new(atomic.Uint64)
	}

	var wg sync.WaitGroup
	wg.Add(len(subset))
	for _, source := range subset {
		go func() {
			defer wg.Done()

			for _, dest := range subset {
				if source == dest || g.HasEdge(source, dest) {
					continue
				}

				nodes, _ := g.AllShortestPathsNodes(source, dest)
				for _, node := range nodes {
					if node != source && node != dest {
						m[node].Add(1)
					}
				}
			}
		}()
	}
	wg.Wait()

	return atomicUint64toIntMap(m)
}

func atomicUint64toIntMap[T comparable](m map[T]*atomic.Uint64) map[T]int {
	res := make(map[T]int, len(m))
	for n, v := range m {
		u := int(v.Load())
		if u != 0 {
			res[n] = u
		}
	}
	return res
}
