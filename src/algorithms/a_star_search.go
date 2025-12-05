/* algorithms: A* Search */
// a shortest path algorithm (Graph).
package main

import (
	"container/heap"
	"fmt"
)

/*
	Analysis on Complexity:

- Time: O(N) average/worst case.
- Space: O(N^2) as the PQ get 'bloated' due to laziness.
*/
func main() {
	astar := &Astar{}
	mat := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}

	ans := astar.ShortestPath(mat)
	fmt.Printf("Shortest path length: %d\n", ans)
}

func (a *Astar) ShortestPath(mat [][]int) int {
	dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	N := len(mat)

	if mat[0][0] != 0 || mat[N-1][N-1] != 0 {
		return -1
	}

	manhattan := func(r, c int) int {
		return 2*N - 2 - r - c
	}

	pq := &PriorityQueue{
		heap:      make([]Node, 0),
		heuristic: manhattan,
	}
	heap.Init(pq)
	heap.Push(pq, Node{r: 0, c: 0, g: 0})

	gCost := make(map[Position]int)
	gCost[Position{r: 0, c: 0}] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Node)
		currPos := Position{curr.r, curr.c}

		// Already found a better path to pos.
		if currG, exists := gCost[currPos]; exists && curr.g > currG {
			continue
		}

		// Destination reached -> return.
		if curr.r == N-1 && curr.c == N-1 {
			return curr.g
		}

		for _, dir := range dirs {
			newR, newC := curr.r+dir[0], curr.c+dir[1]
			neiPos := Position{newR, newC}

			if min(newR, newC) < 0 || max(newR, newC) >= N {
				continue
			}

			if mat[newR][newC] != 0 { // blocked
				continue
			}

			newG := curr.g + 1

			// Haven't computed OR found a better path -> relax.
			if existG, exists := gCost[neiPos]; !exists || newG < existG {
				gCost[neiPos] = newG
				heap.Push(pq, Node{newR, newC, newG})
			}
		}
	}

	return -1
}

type Astar struct{}
type Position struct{ r, c int }
type Node struct {
	r, c int
	g    int
}

type PriorityQueue struct {
	heap      []Node
	heuristic func(int, int) int
}

func (pq PriorityQueue) Len() int { return len(pq.heap) }

func (pq *PriorityQueue) Push(x interface{}) {
	pq.heap = append(pq.heap, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.heap)
	last := pq.heap[n-1]
	pq.heap = pq.heap[:n-1]

	return last
}

func (pq PriorityQueue) Less(a, b int) bool {
	nodeA, nodeB := pq.heap[a], pq.heap[b]

	// f(x) = g(x) + h(x).
	fa := nodeA.g + pq.heuristic(nodeA.r, nodeA.c)
	fb := nodeB.g + pq.heuristic(nodeB.r, nodeB.c)

	return fa < fb
}

func (pq PriorityQueue) Swap(a, b int) {
	pq.heap[a], pq.heap[b] = pq.heap[b], pq.heap[a]
}
