package week4

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func ExampleKargerMinCut() {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(1, 0)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 0)
	g.AddEdge(3, 1)
	g.VertexNumber = 4
	fmt.Println(KargerMinCut(g, 100))
	// Output: 3
}

func importGraph(path string) *graph {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	g := NewGraph()
	for scanner.Scan() {
		vertexes := strings.Split(scanner.Text(), "\t")
		g.AddVertex()
		startVertex, err := strconv.Atoi(vertexes[0])
		if err != nil {
			panic(err)
		}
		for _, endVertexString := range vertexes[1:] {
			endVertex, err := strconv.Atoi(endVertexString)
			if err != nil {
				if (endVertexString == "") {
					break
				} else {
					panic(err)
				}
			}
			g.AddEdge(startVertex, endVertex)
		}
	}
	return g
}
func ExampleKargerMinCut_Large() {
	g := importGraph("kargerMinCut.txt")
	fmt.Println(KargerMinCut(g, 1000))
	// Output: 34
}