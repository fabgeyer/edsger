package edsger

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"slices"
)

// Based on https://pkg.go.dev/container/heap
type priorityQueue[T comparable, N Number] struct {
	items []T
	m     map[T]int // value to index
	pr    map[T]N   // value to priority
}

func newPriorityQueue[T comparable, N Number](n int) *priorityQueue[T, N] {
	return &priorityQueue[T, N]{
		items: make([]T, 0, n),
		m:     make(map[T]int, n),
		pr:    make(map[T]N, n),
	}
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

// Implementation of Dijkstra's shortest path algorithm using a priority queue
func (g *Graph[T, N]) sourceShortestPathMap(source T, withMultiplePaths bool, excludedNodes map[T]bool) map[T][]T {
	maxW := MaxValue[N]()
	prev := make(map[T][]T, g.NumberOfNodes())

	// Manually initialize the priority queue
	q := newPriorityQueue[T, N](g.NumberOfNodes())
	q.items = append(q.items, source)
	q.m[source] = 0
	q.pr[source] = 0
	for n := range g.nodes {
		if n == source {
		} else if _, ok := excludedNodes[n]; !ok {
			q.items = append(q.items, n)
			q.pr[n] = maxW
			q.m[n] = len(q.m)
		}
	}

	for q.Len() > 0 {
		u := heap.Pop(q).(T)
		for _, v := range g.Neighbors(u) {
			if _, ok := excludedNodes[v.Node]; ok {
				continue
			}

			var alt N
			if q.pr[u] == maxW {
				// We prevent here any integer overflow
				alt = maxW
			} else if v.Weight < 0 {
				panic(fmt.Sprintf("Edge (%v, %v) has a negative weight!", u, v.Node))
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

	return prev
}

// Implementation of Dijkstra's shortest path algorithm using a priority queue
func (g *Graph[T, N]) shortestPathMap(source, dest T, withMultiplePaths bool, excludedNodes map[T]bool) (map[T][]T, N) {
	g.validatePathNodes(source, dest)

	maxW := MaxValue[N]()
	prev := make(map[T][]T, g.NumberOfNodes())

	// Manually initialize the priority queue
	q := newPriorityQueue[T, N](g.NumberOfNodes())
	q.items = append(q.items, source)
	q.m[source] = 0
	q.pr[source] = 0
	for n := range g.nodes {
		if n == source {
		} else if _, ok := excludedNodes[n]; !ok {
			q.items = append(q.items, n)
			q.pr[n] = maxW
			q.m[n] = len(q.m)
		}
	}

	for q.Len() > 0 {
		u := heap.Pop(q).(T)
		if u == dest {
			break
		}

		for _, v := range g.Neighbors(u) {
			if _, ok := excludedNodes[v.Node]; ok {
				continue
			}

			var alt N
			if q.pr[u] == maxW {
				// We prevent here any integer overflow
				alt = maxW
			} else if v.Weight < 0 {
				panic(fmt.Sprintf("Edge (%v, %v) has a negative weight!", u, v.Node))
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

func pathFromShortestPathMap[T comparable, N Number](dest T, prev map[T][]T, dist N) ([]T, N) {
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

func (g *Graph[T, N]) DijkstraShortestPath(source, dest T) ([]T, N) {
	prev, dist := g.shortestPathMap(source, dest, false, nil)
	return pathFromShortestPathMap(dest, prev, dist)
}

func (g *Graph[T, N]) DijkstraShortestPathWithExclusionMap(source, dest T, excludedNodes map[T]bool) ([]T, N) {
	prev, dist := g.shortestPathMap(source, dest, false, excludedNodes)
	return pathFromShortestPathMap(dest, prev, dist)
}

func (g *Graph[T, N]) AllDijkstraShortestPathsMap(source, dest T) (map[T][]T, N) {
	return g.shortestPathMap(source, dest, true, nil)
}

func (g *Graph[T, N]) AllShortestPathsNodes(source, dest T) ([]T, N) {
	// Returns all nodes which are part of the shortest path

	prev, dist := g.shortestPathMap(source, dest, true, nil)
	if prev == nil {
		// No path was found
		return nil, 0
	}

	visited := make(map[T]bool, g.NumberOfNodes())
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

func (g *Graph[T, N]) DijkstraShortestPathWithoutNodes(source, dest T) ([]T, N) {
	prev, dist := g.shortestPathMap(source, dest, false, nil)
	return pathFromShortestPathMap(dest, prev, dist)
}

type DijkstraDisjointShortestPathIterator[T comparable, N Number] struct {
	dest T
	prev map[T][]T
	dist N
	path []T
}

func (g *Graph[T, N]) AllDijkstraDisjointShortestPaths(source, dest T) *DijkstraDisjointShortestPathIterator[T, N] {
	prev, dist := g.shortestPathMap(source, dest, true, nil)
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

type item[T comparable, N Number] struct {
	node T
	path []T
	cost N
}

// A basicPriorityQueue implements heap.Interface and holds items.
type basicPriorityQueue[T comparable, N Number] []*item[T, N]

func (pq basicPriorityQueue[T, N]) Len() int { return len(pq) }

func (pq basicPriorityQueue[T, N]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, cost so we use greater than here.
	return pq[i].cost < pq[j].cost
}

func (pq basicPriorityQueue[T, N]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *basicPriorityQueue[T, N]) Push(x any) {
	item := x.(*item[T, N])
	*pq = append(*pq, item)
}

func (pq *basicPriorityQueue[T, N]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func (g *Graph[T, N]) ShortestPathWithMinCost(source, dest T, minCost N) ([]T, N) {
	q := basicPriorityQueue[T, N]{&item[T, N]{
		node: source,
		cost: 0,
		path: []T{source},
	}}

	visited := make(map[T]map[N]bool, g.NumberOfNodes())
	for node := range g.Nodes() {
		visited[node] = make(map[N]bool, g.NumberOfNodes())
	}

	maxW := MaxValue[N]()
	for q.Len() > 0 {
		u := heap.Pop(&q).(*item[T, N])
		if u.node == dest && u.cost >= minCost {
			return u.path, u.cost
		}

		for _, v := range g.Neighbors(u.node) {
			var alt N
			if u.cost == maxW {
				// We prevent here any integer overflow
				alt = maxW
			} else {
				alt = u.cost + v.Weight
			}

			if visited[v.Node][alt] {
				continue
			}
			visited[v.Node][alt] = true

			if slices.Contains(u.path, v.Node) { // Avoid loops
				continue
			}

			if v.Node != dest || alt >= minCost {
				path := make([]T, len(u.path)+1)
				copy(path, u.path)
				path[len(path)-1] = v.Node

				heap.Push(&q, &item[T, N]{
					node: v.Node,
					cost: alt,
					path: path,
				})
			}
		}
	}

	return nil, maxW
}

func (g *Graph[T, N]) ShortestPathWithMinNodes(source, dest T, minNodes int) ([]T, int) {
	q := basicPriorityQueue[T, int]{&item[T, int]{
		node: source,
		cost: 1,
		path: []T{source},
	}}

	visited := make(map[T]map[int]bool, g.NumberOfNodes())
	for node := range g.Nodes() {
		visited[node] = make(map[int]bool, g.NumberOfNodes())
	}

	for q.Len() > 0 {
		u := heap.Pop(&q).(*item[T, int])
		if u.node == dest && u.cost >= minNodes {
			return u.path, u.cost
		}

		for _, v := range g.Neighbors(u.node) {
			alt := u.cost + 1
			if visited[v.Node][alt] {
				continue
			}
			visited[v.Node][alt] = true

			if slices.Contains(u.path, v.Node) { // Avoid loops
				continue
			}

			if v.Node != dest || alt >= minNodes {
				path := make([]T, len(u.path)+1)
				copy(path, u.path)
				path[len(path)-1] = v.Node

				heap.Push(&q, &item[T, int]{
					node: v.Node,
					cost: alt,
					path: path,
				})
			}
		}
	}

	return nil, math.MaxInt
}
