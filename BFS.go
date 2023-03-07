package main

import "fmt"

type graph struct {
	adjMatrix [5][5]int
}

func (g *graph) init() [5][5]int {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			g.adjMatrix[i][j] = 0
		}
	}
	return g.adjMatrix
}

func (g *graph) addEdge(vertex1, vertex2 int) {
	g.adjMatrix[vertex1][vertex2] = 1
}

func (g *graph) BFS(num int) {
	Queue := []int{num}
	visited := make(map[int]bool)

	for len(Queue) > 0 {

		u := Queue[0]
		Queue = Queue[1:]

		if !visited[u] {
			visited[u] = true

			fmt.Printf("%d ", u)

			for j := 0; j < 5; j++ {
				if g.adjMatrix[u][j] == 1 {
					Queue = append(Queue, j)
				}
			}
		}
	}
}

func main() {
	myGraph := graph{}
	myGraph.init()
	myGraph.addEdge(0, 1)
	myGraph.addEdge(0, 2)
	myGraph.addEdge(1, 0)
	myGraph.addEdge(1, 3)
	myGraph.addEdge(2, 0)
	fmt.Println()
	myGraph.BFS(2)
}
