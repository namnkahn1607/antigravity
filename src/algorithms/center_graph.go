/* algorithms: Centroid of a Graph */
// a minimum height tree algorithm
package main

import "fmt"

/*
	Analysis on Complexity

- Time: O(V + E) average/worst case.
- Space: O(V + E).
*/
func centroid() {
	centroid := &Centroid{}
	V := 6
	edges := [][]int{
		{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4},
	}

	ans := centroid.centerGraph(V, edges)
	fmt.Println("The center(s) of Graph are:")

	for _, center := range ans {
		fmt.Println(center)
	}
}

func (c *Centroid) centerGraph(V int, edges [][]int) []int {
	adj := make([][]int, V)
	degree := make([]int, V)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		degree[u]++
		adj[v] = append(adj[v], u)
		degree[v]++
	}

	remain := V
	leaves := []int{}

	// Start with 'leaf' vertices: 0 or 1 in degree.
	for i, deg := range degree {
		if deg <= 1 {
			leaves = append(leaves, i)
		}
	}

	// Maximum of 2 possible center(s).
	for remain > 2 {
		newLeaves := []int{}

		for _, leaf := range leaves {
			for _, nei := range adj[leaf] {
				// Shear out the leaves -> Decrease parents' degree.
				degree[nei]--

				// If parents become leaves -> Queue for shearing.
				if degree[nei] == 1 {
					newLeaves = append(newLeaves, nei)
				}
			}
		}

		remain -= len(newLeaves)
		leaves = newLeaves
	}

	return leaves
}

type Centroid struct{}
