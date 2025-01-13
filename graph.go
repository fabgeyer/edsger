package edsger

import (
	"iter"
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
	if g.HasEdge(source, dest) {
		panic("Edge already defined")
	}

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

func (g *Graph[T, N]) GetEdge(source, dest T) (N, bool) {
	g.validatePathNodes(source, dest)
	for _, edge := range g.edges[source] {
		if edge.Node == dest {
			return edge.Weight, true
		}
	}
	return N(0), false
}

func (g *Graph[T, N]) HasEdge(source, dest T) bool {
	_, ok := g.GetEdge(source, dest)
	return ok
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

func (g *Graph[T, N]) Degree() iter.Seq2[T, int] {
	return func(yield func(n T, d int) bool) {
		for node, edges := range g.edges {
			if !yield(node, len(edges)) {
				return
			}
		}
	}
}

func (g *Graph[T, N]) Nodes() iter.Seq[T] {
	return func(yield func(n T) bool) {
		for n := range g.nodes {
			if !yield(n) {
				return
			}
		}
	}
}

func (g *Graph[T, N]) NodesList() []T {
	res := make([]T, len(g.nodes))
	for n, i := range g.nodes {
		res[i] = n
	}
	return res
}

func (g *Graph[T, N]) Edges() iter.Seq[*WeightedEdge[T, N]] {
	return func(yield func(x *WeightedEdge[T, N]) bool) {
		for src, edges := range g.edges {
			srcid := g.nodes[src]
			for _, edge := range edges {
				if !g.directed && srcid <= g.nodes[edge.Node] {
					continue
				}
				if !yield(&WeightedEdge[T, N]{
					From:   src,
					To:     edge.Node,
					Weight: edge.Weight,
				}) {
					return
				}
			}
		}
	}
}
