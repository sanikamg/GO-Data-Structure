package main

import "fmt"

//structure for vertex
type vertex struct {
	key            int
	adjacentVertex []*vertex
}

// structure for graph
type Graph struct {
	vertices []*vertex
}

func (g *Graph) AddVertex(key int) {
	if g.contains(g.vertices, key) {
		err := fmt.Errorf("vertex %v already exists", key)
		fmt.Println(err.Error())
	} else {
		newNode := new(vertex)
		newNode.key = key
		g.vertices = append(g.vertices, newNode)
	}
}

func (g *Graph) contains(s []*vertex, key int) bool {
	for _, v := range s {
		if v.key == key {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge %v--->%v", from, to)
		fmt.Println(err.Error())
	} else if g.contains(fromVertex.adjacentVertex, to) {
		err := fmt.Errorf("edge %v--->%v already exists", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacentVertex = append(fromVertex.adjacentVertex, toVertex)
	}
}

func (g *Graph) getVertex(key int) *vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nvertex %v : ", v.key)
		for _, v := range v.adjacentVertex {
			fmt.Printf(" %v ", v.key)
		}
	}
}

func (g *Graph) removeEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if g.contains(fromVertex.adjacentVertex, toVertex.key) {
		for i, v := range fromVertex.adjacentVertex {
			if v.key == toVertex.key {
				fromVertex.adjacentVertex = append(fromVertex.adjacentVertex[:i], fromVertex.adjacentVertex[i+1:]...)
				return
			}
		}
	}
	fmt.Printf("\nedge %v-->%v does not exist", from, to)
}

func (g *Graph) removeVertex(key int) {
	for _, v := range g.vertices {
		g.removeEdge(v.key, key)
	}

	for i, v := range g.vertices {
		if v.key == key {
			g.vertices = append(g.vertices[:i], g.vertices[i+1:]...)
			return
		}

	}
	fmt.Printf("\n %v vertex does not exist ", key)

}

func main() {
	myGraph := Graph{}
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

	myGraph.Print()
	fmt.Println(" ")
	myGraph.removeVertex(20)
	myGraph.Print()

}
