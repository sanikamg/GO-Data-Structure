package main

import "fmt"

//graph structure
type graph struct {
	vertices []*vertex
}

//vertex structure
type vertex struct {
	key            int
	adjacentVertex []*vertex
}

func (g *graph) AddVertex(num int) {
	if g.contains(g.vertices, num) {
		return
	} else {
		newNode := new(vertex)
		newNode.key = num
		g.vertices = append(g.vertices, newNode)
	}
}

func (g *graph) contains(s []*vertex, num int) bool {
	for _, v := range s {
		if v.key == num {
			return true
		}

	}
	return false
}

func (g *graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		return
	} else if g.contains(fromVertex.adjacentVertex, to) {
		return
	} else {
		fromVertex.adjacentVertex = append(fromVertex.adjacentVertex, toVertex)
	}

}
func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("\n%d : ", v.key)
		for _, v := range v.adjacentVertex {
			fmt.Printf("%d --> ", v.key)
		}
	}
}

func (g *graph) getVertex(key int) *vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}
	return nil
}

func main() {
	myGraph := graph{}

	myGraph.AddVertex(10)
	myGraph.AddVertex(20)
	myGraph.AddVertex(30)
	myGraph.AddVertex(40)
	myGraph.AddVertex(50)
	myGraph.AddEdge(10, 50)
	myGraph.AddEdge(10, 20)
	myGraph.AddEdge(20, 30)
	myGraph.AddEdge(20, 10)
	myGraph.AddEdge(30, 40)
	myGraph.AddEdge(30, 20)
	myGraph.AddEdge(40, 50)
	myGraph.AddEdge(40, 30)
	myGraph.AddEdge(50, 10)
	myGraph.AddEdge(50, 40)
	myGraph.print()
}
