package main

import "fmt"

//structure of a adjecency list graph
type graph struct {
	vertices []*Vertex
}

//vertex structure of a graph
type Vertex struct {
	key      int
	adjacent []*Vertex
}

//AddVertex adds a vertex to the graph
func (g *graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		newNode := new(Vertex)
		newNode.key = k
		g.vertices = append(g.vertices, newNode)
	}
}

// AddEdge adds an edge to the graph
func (g *graph) AddEdge(from int, to int) {
	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing  edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

// getVertex63
func (g *graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// print will print the adjacent list for each vertex of the graph
func (g *graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

func main() {
	myGraph := graph{}
	for i := 1; i <= 5; i++ {
		myGraph.AddVertex(i)
	}
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 4)
	myGraph.AddEdge(2, 1)
	myGraph.AddEdge(2, 3)
	myGraph.AddEdge(2, 4)
	myGraph.AddEdge(3, 2)
	myGraph.AddEdge(3, 4)
	myGraph.AddEdge(3, 5)
	myGraph.AddEdge(4, 1)
	myGraph.AddEdge(4, 2)
	myGraph.AddEdge(4, 3)
	myGraph.AddEdge(4, 5)
	myGraph.AddEdge(5, 3)
	myGraph.AddEdge(5, 4)

	myGraph.Print()
}
