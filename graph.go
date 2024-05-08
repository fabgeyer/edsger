package edsger

type NodeWeight[T comparable, N Number] struct {
	Node   T
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
