package main

import "fmt"

const vertices = 5

type graph struct {
	adjMatrix [5][5]int
}

// initialize matrix
func (g *graph) init() [5][5]int {
	for i := 0; i < vertices; i++ {
		for j := 0; j < vertices; j++ {
			g.adjMatrix[i][j] = 0
		}
	}

	return g.adjMatrix
}

func (g *graph) addEdge(i int, j int) {
	g.adjMatrix[i][j] = 1
}

func (g *graph) BFS(num int) {
	Queue := []int{num}
	visited := make(map[int]bool)
	for len(Queue) > 0 {
		// dequeue
		u := Queue[0]
		Queue = Queue[1:]

		//mark as visited
		if !visited[u] {
			visited[u] = true

			fmt.Printf("%d ", u)

			//enqueue
			i := u
			for j := 0; j < vertices; j++ {
				if g.adjMatrix[i][j] == 1 {
					if !visited[j] {
						Queue = append(Queue, j)
					}

				}
			}

		}
	}
	fmt.Println(visited)
}

//dfs traversal

func (g *graph) DFS(num int) {
	visited := make(map[int]bool)
	Stack := []int{num}
	for len(Stack) > 0 {
		u := Stack[len(Stack)-1]
		Stack = Stack[:len(Stack)-1]

		if !visited[u] {
			visited[u] = true
			fmt.Printf("%d ", u)

			for j := 0; j < vertices; j++ {
				if g.adjMatrix[u][j] == 1 && !visited[j] {
					Stack = append(Stack, j)

				}

			}

		}

	}

}

//print graph
func (g *graph) printGraph() {
	for i := 0; i < vertices; i++ {
		if i != 0 {
			fmt.Printf("%d ", i)
		}
		for j := 0; j < vertices; j++ {
			if i == 0 && j == 0 {
				fmt.Printf("  ")
				for k := 0; k < vertices; k++ {
					fmt.Printf("%d ", k)
				}
				fmt.Printf("\n")
				fmt.Printf("%d ", i)
			}
			fmt.Printf("%d ", g.adjMatrix[i][j])
		}
		fmt.Println(" ")
	}
}
func main() {
	mygraph := graph{}
	mygraph.init()
	mygraph.addEdge(0, 2)
	mygraph.addEdge(0, 1)
	mygraph.addEdge(2, 0)
	mygraph.addEdge(2, 3)
	mygraph.addEdge(3, 2)
	mygraph.addEdge(3, 4)
	mygraph.addEdge(4, 3)
	mygraph.addEdge(4, 1)
	mygraph.addEdge(1, 0)
	mygraph.addEdge(1, 4)
	mygraph.printGraph()
	mygraph.BFS(2)
	mygraph.DFS(2)
}
