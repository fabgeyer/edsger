package edsger

import (
	"container/heap"
	"fmt"
	"math/rand"
	"slices"
)

// Based on https://pkg.go.dev/container/heap
type priorityQueue[T comparable, N Number] struct {
	items []T
	m     map[T]int // value to index
	pr    map[T]N   // value to priority
}

func (pq *priorityQueue[T, N]) Len() int { return len(pq.items) }

func (pq *priorityQueue[T, N]) Less(i, j int) bool { return pq.pr[pq.items[i]] < pq.pr[pq.items[j]] }

func (pq *priorityQueue[T, N]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.m[pq.items[i]] = i
	pq.m[pq.items[j]] = j
}

func (pq *priorityQueue[T, N]) Push(x interface{}) {
	n := len(pq.items)
	item := x.(T)
	pq.m[item] = n
	pq.items = append(pq.items, item)
}

func (pq *priorityQueue[T, N]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.m[item] = -1
	pq.items = old[0 : n-1]
	return item
}

// update modifies the priority of an item in the queue.
func (pq *priorityQueue[T, N]) update(item T, priority N) {
	pq.pr[item] = priority
	heap.Fix(pq, pq.m[item])
}

func (pq *priorityQueue[T, N]) addWithPriority(item T, priority N) {
	heap.Push(pq, item)
	pq.update(item, priority)
}

func (g *Graph[T, N]) validateAllWeightsArePositive() {
	if !Signed[N]() {
		return
	}

	for src, dests := range g.edges {
		for _, n := range dests {
			if n.Weight < 0 {
				msg := fmt.Sprintf("Edge (%v, %v) has a negative weight!", src, n.Node)
				panic(msg)
			}
		}
	}
}

func (g *Graph[T, N]) shortestPathMap(source, dest T, withMultiplePaths bool) (map[T][]T, N) {
	// Implementation of Dijkstra's shortest path algorithm using a priority queue

	g.validateAllWeightsArePositive()
	g.validatePathNodes(source, dest)

	prev := make(map[T][]T, len(g.nodes))
	q := &priorityQueue[T, N]{
		items: make([]T, 0, len(g.nodes)),
		m:     make(map[T]int, len(g.nodes)),
		pr:    make(map[T]N, len(g.nodes)),
	}

	maxW := MaxValue[N]()
	for n := range g.nodes {
		if n == source {
			q.addWithPriority(n, 0)
		} else {
			q.addWithPriority(n, maxW)
		}
	}

	for q.Len() > 0 {
		u := heap.Pop(q).(T)
		for _, v := range g.Neighbors(u) {
			var alt N
			if q.pr[u] == maxW {
				// We prevent here any integer overflow
				alt = maxW
			} else {
				alt = q.pr[u] + v.Weight
			}

			if alt < q.pr[v.Node] {
				prev[v.Node] = []T{u}
				q.update(v.Node, alt)
			} else if withMultiplePaths && alt == q.pr[v.Node] {
				prev[v.Node] = append(prev[v.Node], u)
				q.update(v.Node, alt)
			}
		}
	}

	if q.pr[dest] == maxW {
		// No path was found
		return nil, 0
	}
	return prev, q.pr[dest]
}

func (g *Graph[T, N]) DijkstraShortestPath(source, dest T) ([]T, N) {
	prev, dist := g.shortestPathMap(source, dest, false)
	if prev == nil {
		// No path was found
		return nil, 0
	}

	v := dest
	path := []T{v}
	for {
		vs, ok := prev[v]
		if !ok {
			break
		}
		v = vs[0]
		path = append(path, v)
	}
	slices.Reverse(path)
	return path, dist
}

func (g *Graph[T, N]) AllDijkstraShortestPathsMap(source, dest T) (map[T][]T, N) {
	return g.shortestPathMap(source, dest, true)
}

func (g *Graph[T, N]) AllShortestPathsNodes(source, dest T) ([]T, N) {
	// Returns all nodes which are part of the shortest path

	prev, dist := g.shortestPathMap(source, dest, true)
	if prev == nil {
		// No path was found
		return nil, 0
	}

	visited := make(map[T]bool)
	q := []T{dest}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if visited[v] {
			continue
		}
		visited[v] = true

		vs, ok := prev[v]
		if ok {
			q = slices.Concat(q, vs)
		}
	}

	res := make([]T, len(visited))
	i := 0
	for n := range visited {
		res[i] = n
		i++
	}
	return res, dist
}

type DijkstraDisjointShortestPathIterator[T comparable, N Number] struct {
	dest T
	prev map[T][]T
	dist N
	path []T
}

func (g *Graph[T, N]) AllDijkstraDisjointShortestPaths(source, dest T) *DijkstraDisjointShortestPathIterator[T, N] {
	prev, dist := g.shortestPathMap(source, dest, true)
	if prev == nil {
		return &DijkstraDisjointShortestPathIterator[T, N]{}
	}

	return &DijkstraDisjointShortestPathIterator[T, N]{
		dest: dest,
		prev: prev,
		dist: dist,
	}
}

func (it *DijkstraDisjointShortestPathIterator[T, N]) Shuffle() *DijkstraDisjointShortestPathIterator[T, N] {
	for _, v := range it.prev {
		rand.Shuffle(len(v), func(i, j int) {
			v[i], v[j] = v[j], v[i]
		})
	}
	return it
}

func (it *DijkstraDisjointShortestPathIterator[T, N]) Next() bool {
	if it.dist == 0 {
		return false
	}

	v := it.dest
	it.path = []T{v}
	for {
		vs, ok := it.prev[v]
		if !ok {
			break
		}
		if len(vs) == 0 {
			clear(it.prev)
			it.path = nil
			it.dist = 0
			return false
		}
		v, it.prev[v] = vs[0], vs[1:]
		it.path = append(it.path, v)
	}
	slices.Reverse(it.path)
	return true
}

func (it *DijkstraDisjointShortestPathIterator[T, N]) Get() ([]T, N) {
	return it.path, it.dist
}
