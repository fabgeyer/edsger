package edsger

import (
	"math/rand"
	"slices"
)

func (g *Graph[T, N]) HasSimplePath(source, dest T) bool {
	g.validatePathNodes(source, dest)
	visited := make(map[T]bool)
	return g.hasSimplePath(source, dest, visited)
}

func (g *Graph[T, N]) hasSimplePath(source, dest T, visited map[T]bool) bool {
	if source == dest {
		return true
	}

	visited[source] = true
	for _, edge := range g.edges[source] {
		if !visited[edge.Node] {
			if g.hasSimplePath(edge.Node, dest, visited) {
				return true
			}
		}
	}
	return false
}

func (g *Graph[T, N]) SimplePath(source, dest T) ([]T, N) {
	g.validatePathNodes(source, dest)
	visited := make(map[T]bool)
	return g.simplePath(source, dest, visited, []T{source}, 0)
}

func (g *Graph[T, N]) simplePath(source, dest T, visited map[T]bool, currentPath []T, totalWeight N) ([]T, N) {
	if source == dest {
		return currentPath, totalWeight
	}

	visited[source] = true
	for _, edge := range g.edges[source] {
		if !visited[edge.Node] {
			path, total := g.simplePath(edge.Node, dest, visited,
				append(currentPath, edge.Node), totalWeight+edge.Weight)
			if path != nil {
				return path, total
			}
		}
	}

	// No path found
	visited[source] = false
	return nil, 0
}

type SPIStackElem[T comparable, N Number] struct {
	node   T
	weight N
	edges  []*NodeWeight[T, N]
}

type SimplePathIterator[T comparable, N Number] struct {
	CutoffWeight N
	CutoffHops   int

	g         *Graph[T, N]
	visited   map[T]bool
	dest      T
	stack     []*SPIStackElem[T, N]
	heuristic func(i, j *NodeWeight[T, N]) int

	// Returned value
	path        []T
	totalWeight N
}

func (g *Graph[T, N]) AllSimplePaths(source, dest T) *SimplePathIterator[T, N] {
	g.validatePathNodes(source, dest)
	return &SimplePathIterator[T, N]{
		g:            g,
		CutoffWeight: MaxValue[N](),
		CutoffHops:   MaxInt[int](),
		visited: map[T]bool{
			source: true,
		},
		dest: dest,
		stack: []*SPIStackElem[T, N]{
			{
				node:  source,
				edges: g.edges[source],
			},
		},
	}
}

func (g *Graph[T, N]) AllSimplePathsWithHeuristic(source, dest T, heuristic func(i, j *NodeWeight[T, N]) int) *SimplePathIterator[T, N] {
	it := g.AllSimplePaths(source, dest)
	it.heuristic = heuristic
	for _, e := range it.stack {
		e.edges = it.applyHeuristic(e.edges)
	}
	return it
}

func (it *SimplePathIterator[T, N]) applyHeuristic(edges []*NodeWeight[T, N]) []*NodeWeight[T, N] {
	if it.heuristic == nil {
		return edges
	}

	res := make([]*NodeWeight[T, N], len(edges))
	copy(res, edges)
	slices.SortFunc(res, it.heuristic)
	return res
}

func (it *SimplePathIterator[T, N]) Next() bool {
	n := len(it.stack) - 1
	for n >= 0 {
		top := it.stack[n]
		if top.edges == nil {
			it.visited[top.node] = false
			it.stack = it.stack[:n]
			n--
			continue
		}

		var i int
		if it.heuristic == nil {
			i = rand.Intn(len(top.edges))
		} else {
			i = 0
		}
		edge := top.edges[i]
		top.edges = slices.Concat(top.edges[:i], top.edges[i+1:])

		if edge.Node == it.dest {
			it.totalWeight = top.weight + edge.Weight
			if it.totalWeight > it.CutoffWeight {
				continue
			}

			// Build the final path based on the stack
			it.path = make([]T, len(it.stack)+1)
			for i := range it.stack {
				it.path[i] = it.stack[i].node
			}
			it.path[len(it.stack)] = it.dest

			return true

		} else if !it.visited[edge.Node] && len(it.stack) < it.CutoffHops-1 {
			weight := top.weight + edge.Weight
			if weight > it.CutoffWeight {
				continue
			}

			it.visited[edge.Node] = true
			it.stack = append(it.stack, &SPIStackElem[T, N]{
				node:   edge.Node,
				weight: weight,
				edges:  it.applyHeuristic(it.g.edges[edge.Node]),
			})
			n++
		}
	}

	it.path = nil
	it.stack = nil
	it.totalWeight = 0
	return false
}

func (it *SimplePathIterator[T, N]) Get() ([]T, N) {
	return it.path, it.totalWeight
}
