package edsger

import (
	"maps"
	"slices"
)

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
	edges    map[T][]*NodeWeight[T, N]
	directed bool
}

// Returns a new directed graph
func NewDirectedGraph[T comparable, N Number]() *Graph[T, N] {
	return &Graph[T, N]{
		nodes:    make(map[T]int),
		edges:    make(map[T][]*NodeWeight[T, N]),
		directed: true,
	}
}

// Returns a new undirected graph
func NewUndirectedGraph[T comparable, N Number]() *Graph[T, N] {
	return &Graph[T, N]{
		nodes:    make(map[T]int),
		edges:    make(map[T][]*NodeWeight[T, N]),
		directed: false,
	}
}

func (g *Graph[T, N]) Clone() *Graph[T, N] {
	edges := make(map[T][]*NodeWeight[T, N], len(g.edges))
	for k, v := range g.edges {
		edges[k] = slices.Clone(v)
	}
	return &Graph[T, N]{
		nodes:    maps.Clone(g.nodes),
		edges:    edges,
		directed: g.directed,
	}
}

func (g *Graph[T, N]) IsDirected() bool {
	return g.directed
}

func (g *Graph[T, N]) NumberOfNodes() int {
	return len(g.nodes)
}

func (g *Graph[T, N]) NumberOfEdges() int {
	n := 0
	for _, edges := range g.edges {
		n += len(edges)
	}

	if g.directed {
		return n
	} else {
		return n / 2
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
	g.validatePathNodes(source, dest)
	g.edges[source] = append(g.edges[source], &NodeWeight[T, N]{
		Node:   dest,
		Weight: weight,
	})
}

func (g *Graph[T, N]) Neighbors(n T) []*NodeWeight[T, N] {
	return g.edges[n]
}

func (g *Graph[T, N]) AllSuccessors() map[T]map[T]bool {
	res := make(map[T]map[T]bool, len(g.nodes))
	for src, edges := range g.edges {
		if len(edges) == 0 {
			continue
		}
		res[src] = make(map[T]bool, len(edges))
		for _, dst := range edges {
			res[src][dst.Node] = true
		}
	}
	return res
}

func (g *Graph[T, N]) AllPredecessors() map[T]map[T]bool {
	res := make(map[T]map[T]bool, len(g.nodes))
	for src, edges := range g.edges {
		for _, dst := range edges {
			if res[dst.Node] == nil {
				res[dst.Node] = make(map[T]bool)
			}
			res[dst.Node][src] = true
		}
	}
	return res
}

func (g *Graph[T, N]) validatePathNodes(source, dest T) {
	if !g.HasNode(source) {
		panic("Invalid source node")
	}
	if !g.HasNode(dest) {
		panic("Invalid destination node")
	}
}

func (g *Graph[T, N]) RemoveNode(node T) {
	if !g.HasNode(node) {
		panic("Invalid node")
	}
	delete(g.nodes, node)
	delete(g.edges, node)

	for other, edges := range g.edges {
		g.edges[other] = slices.DeleteFunc(edges, func(e *NodeWeight[T, N]) bool {
			if e.Node == node {
				return true
			}
			return false
		})
	}
}

func (g *Graph[T, N]) RemoveEdge(source, dest T) {
	g.removeEdge(source, dest)
	if !g.directed {
		g.removeEdge(dest, source)
	}
}

func (g *Graph[T, N]) removeEdge(source, dest T) {
	g.validatePathNodes(source, dest)

	g.edges[source] = slices.DeleteFunc(g.edges[source], func(e *NodeWeight[T, N]) bool {
		if e.Node == dest {
			return true
		}
		return false
	})
}

func (g *Graph[T, N]) Degree() map[T]int {
	res := make(map[T]int, len(g.edges))
	for node, edges := range g.edges {
		res[node] = len(edges)
	}
	return res
}

func (g *Graph[T, N]) Nodes() []T {
	res := make([]T, len(g.nodes))
	for n, i := range g.nodes {
		res[i] = n
	}
	return res
}

func (g *Graph[T, N]) ApplyNodes(fn func(T) bool) {
	for n := range g.nodes {
		if !fn(n) {
			return
		}
	}
}

func (g *Graph[T, N]) ApplyEdges(fn func(T, T, N) bool) {
	for src, edges := range g.edges {
		for _, edge := range edges {
			if !fn(src, edge.Node, edge.Weight) {
				return
			}
		}
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
	i := 0
	for n := range g.nodes {
		nodes[i] = n
		i++
	}

	return &EdgeIterator[T, N]{
		g:     g,
		nodes: nodes,
	}
}

func (it *EdgeIterator[T, N]) Next() bool {
	for len(it.nodes) > 0 {
		for i := it.i; i < len(it.g.edges[it.nodes[0]]); i++ {
			it.edge = it.g.edges[it.nodes[0]][i]
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

func (it *EdgeIterator[T, N]) Get() *WeightedEdge[T, N] {
	return &WeightedEdge[T, N]{
		From:   it.nodes[0],
		To:     it.edge.Node,
		Weight: it.edge.Weight,
	}
}
