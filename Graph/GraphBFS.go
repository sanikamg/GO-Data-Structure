package main

import (
	"fmt"
)

// Graph struct
type Graph struct {
	adjacencyList map[int][]int
}

// addEdge method adds an edge to the graph
func (g *Graph) addEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
}

// bfs method performs BFS on the graph
func (g *Graph) bfs(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		// Dequeue a vertex from queue
		u := queue[0]
		queue = queue[1:]

		// If the vertex is not visited yet, mark it as visited
		if !visited[u] {
			visited[u] = true

			fmt.Printf("%d ", u)

			// Enqueue all adjacent vertices of the dequeued vertex u
			for _, v := range g.adjacencyList[u] {
				if !visited[v] {
					queue = append(queue, v)
				}
			}
		}
	}
}

func main() {
	g := Graph{adjacencyList: make(map[int][]int)}
 
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	fmt.Println("Following is Breadth First Traversal (starting from vertex 2)")
	g.bfs(2)
}
