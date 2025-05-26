package edsger

// Computes betweenness centrality
// Betweenness centrality of a node $v$ is the sum of all-pairs shortest paths that pass through $v$
func (g *Graph[T, N]) BetweennessCentrality() map[T]int {
	res := make(map[T]int)
	for source := range g.Nodes() {
		smap := g.sourceShortestPathMap(source, false, nil)
		for dest := range g.Nodes() {
			if source == dest || g.HasEdge(source, dest) {
				continue
			}

			path, _ := pathFromShortestPathMap(dest, smap, 0)
			for _, node := range path[1 : len(path)-1] {
				res[node]++
			}
		}
	}
	return res
}

// Computes betweenness centrality using all shortest paths
// Betweenness centrality of a node $v$ is the sum of all-pairs all shortest paths that pass through $v$
func (g *Graph[T, N]) AllShortestPathsBetweennessCentrality() map[T]int {
	res := make(map[T]int)
	for source := range g.Nodes() {
		for dest := range g.Nodes() {
			if source == dest || g.HasEdge(source, dest) {
				continue
			}

			nodes, _ := g.AllShortestPathsNodes(source, dest)
			for _, node := range nodes {
				if node != source && node != dest {
					res[node]++
				}
			}
		}
	}
	return res
}
