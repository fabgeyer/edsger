package edsger

type NodeWeight[T comparable, N Number] struct {
	Node   T
	Weight N
}

type WeightedEdge[T comparable, N Number] struct {
	From   T
	To     T
	Weight N
}

type Graph[T comparable, N Number] struct {
	nodes    map[T]int
	edges    map[T][]NodeWeight[T, N]
	directed bool
}

// Returns a new directed graph
func NewDirectedGraph[T comparable, N Number]() *Graph[T, N] {
	return &Graph[T, N]{
		nodes:    make(map[T]int),
		edges:    make(map[T][]NodeWeight[T, N]),
		directed: true,
	}
}

// Returns a new undirected graph
func NewUndirectedGraph[T comparable, N Number]() *Graph[T, N] {
	return &Graph[T, N]{
		nodes:    make(map[T]int),
		edges:    make(map[T][]NodeWeight[T, N]),
		directed: false,
	}
}

func (g *Graph[T, N]) IsDirected() bool {
	return g.directed
}

func (g *Graph[T, N]) NumberOfNodes() int {
	return len(g.nodes)
}

func (g *Graph[T, N]) NumberOfEdges() int {
	if g.directed {
		return len(g.edges)
	} else {
		return len(g.edges) / 2
	}
}

func (g *Graph[T, N]) AddNode(n T) {
	if g.HasNode(n) {
		panic("Node already in graph!")
	}
	g.nodes[n] = len(g.nodes)
}

func (g *Graph[T, N]) HasNode(n T) bool {
	_, ok := g.nodes[n]
	return ok
}

func (g *Graph[T, N]) AddEdge(source, dest T, weight N) {
	g.addEdge(source, dest, weight)
	if !g.directed {
		g.addEdge(dest, source, weight)
	}
}

func (g *Graph[T, N]) addEdge(source, dest T, weight N) {
	if !g.HasNode(source) {
		panic("Invalid 'source' node")
	}
	if !g.HasNode(dest) {
		panic("Invalid 'dest' node")
	}
	g.edges[source] = append(g.edges[source], NodeWeight[T, N]{
		Node:   dest,
		Weight: weight,
	})
}

func (g *Graph[T, N]) Neighbors(n T) []NodeWeight[T, N] {
	return g.edges[n]
}

func (g *Graph[T, N]) validatePathNodes(source, dest T) {
	if !g.HasNode(source) {
		panic("Invalid source node")
	}
	if !g.HasNode(dest) {
		panic("Invalid destination node")
	}
}

type EdgeIterator[T comparable, N Number] struct {
	g     *Graph[T, N]
	nodes []T
	i     int
	edge  *NodeWeight[T, N]
}

func (g *Graph[T, N]) Edges() *EdgeIterator[T, N] {
	nodes := make([]T, len(g.nodes))
	for n, i := range g.nodes {
		nodes[i] = n
	}

	return &EdgeIterator[T, N]{
		g:     g,
		nodes: nodes,
	}
}

func (it *EdgeIterator[T, N]) Next() bool {
	for len(it.nodes) > 0 {
		for i := it.i; i < len(it.g.edges[it.nodes[0]]); i++ {
			it.edge = &it.g.edges[it.nodes[0]][i]
			if it.g.directed || it.g.nodes[it.nodes[0]] <= it.g.nodes[it.edge.Node] {
				it.i = i + 1
				return true
			}
		}

		it.nodes = it.nodes[1:]
		it.i = 0
	}
	return false
}

func (it *EdgeIterator[T, N]) Get() WeightedEdge[T, N] {
	return WeightedEdge[T, N]{
		From:   it.nodes[0],
		To:     it.edge.Node,
		Weight: it.edge.Weight,
	}
}
