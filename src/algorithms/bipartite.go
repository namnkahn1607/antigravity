/* algorithms: Bipartite */
// check if Graph is bipartite (Graph).
package main

import "fmt"

/*
	Analysis on Complexity

- Time: O(V + E) average/worst case.
- Space: O(V + E) for adjacency list.
*/
func bipartite() {
	bipartite := &Bipartite{}
	V := 4
	edges := [][2]int{
		{1, 3}, {0, 2}, {1, 3}, {0, 2},
	}

	ans := bipartite.isBipartite(V, edges)
	fmt.Printf("The given Graph is %t\n", ans)
}

func (b *Bipartite) isBipartite(V int, edges [][2]int) bool {
	adj := make([][]int, V)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
	}

	// Coloring: black -> true, white -> false.
	color := make([]bool, V)
	marked := make([]bool, V)

	for i := 0; i <= V; i++ {
		marked[i] = false
		color[i] = false
	}

	var dfs func(curr int) bool
	dfs = func(curr int) bool {
		marked[curr] = true

		for _, nei := range adj[curr] {
			if marked[nei] {
				// Visited & same color -> not bipartite.
				if color[nei] == color[curr] {
					return false
				}

				continue
			}

			// Color the opposite to all neighbors.
			color[nei] = !color[curr]

			// Already found a falsy clue -> return.
			if !dfs(nei) {
				return false
			}
		}

		return true
	}

	for i := 0; i <= V; i++ {
		if !marked[i] && !dfs(i) {
			return false
		}
	}

	return true
}

type Bipartite struct{}
