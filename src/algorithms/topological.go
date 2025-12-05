/* algorithms: Topological Sort */
// a topological sorting algorithm (Digraph).
package main

import "fmt"

/*
	Analysis on Complexity

- Time: linear O(V + E) average/worst case.
- Space: O(V + E).
*/
func topological() {
	topo := &Topological{}
	V := 7
	edges := [][]int{
		{1, 2}, {1, 4}, {2, 3}, {4, 5},
		{4, 2}, {4, 6}, {5, 6},
	}

	ans := topo.topoOrder(V, edges)
	fmt.Println("Topological order of the Digraph:")

	for i := range V {
		fmt.Println(ans[i])
	}
}

func (t *Topological) topoOrder(V int, edges [][]int) []int {
	adj := make([][]int, V)
	order := []int{}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
	}

	onPath := make([]bool, V)
	marked := make([]bool, V)

	var dfs func(vertex int) bool
	dfs = func(vertex int) bool {
		marked[vertex] = true
		onPath[vertex] = true

		for _, nei := range adj[vertex] {
			if onPath[nei] {
				return false
			}

			if marked[nei] {
				continue
			}

			if !dfs(nei) {
				return false
			}
		}

		order = append(order, vertex)
		onPath[vertex] = false

		return true
	}

	for i := range V {
		if !marked[i] && !dfs(i) {
			return []int{}
		}
	}

	return order
}

type Topological struct{}
