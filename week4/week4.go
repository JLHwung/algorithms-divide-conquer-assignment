package week4

import (
	"math/rand"
	"time"
)

// Graph is a simple Graph struct to implement Krager's Algorithm
type Graph struct {
	VertexNumber int
	Edges []edge
}

type edge struct {
	start int
	end int
}

// NewGraph create a Graph and returns reference to the created Graph
func NewGraph() *Graph {
	VertexNumber := 0
	Edges := make([]edge, 0)
	return &Graph{VertexNumber, Edges}
}

// AddEdge will add an edge starting with start and ending with end
func (g *Graph) AddEdge (start, end int) {
	newEdge := edge{start, end}
	g.Edges = append(g.Edges, newEdge)
}

// AddVertex simply bump the vertexNumber of a Graph
func (g *Graph) AddVertex() {
	g.VertexNumber++
}

// CloneGraph does a deep copy of the Graph
func (g *Graph) CloneGraph() *Graph {
	VertexNumber := g.VertexNumber
	Edges := make([]edge, len(g.Edges))
	copy(Edges, g.Edges)
	return &Graph{
		VertexNumber,
		Edges,
	}
}

// ContractEdge implements Karger's Contraction. It will contract the edge with index i, remove the endpoint of edge(i)
// and rewrite all the edges related to the endpoint of edge(i)
func (g *Graph) ContractEdge(i int) {
	e := g.Edges[i]

	// update edges
	g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)

	startVertex := e.start
	endVertex := e.end

	// Now we stick endVertex to startVertex

	// update edges where start/end is the endVertex
	for j, edge := range g.Edges {
		if edge.start == endVertex {
			g.Edges[j].start = startVertex
		}
		if edge.end == endVertex {
			g.Edges[j].end = startVertex
		}
	}

	// remove self loops
	edges := []edge{}
	for _, edge := range g.Edges {
		if edge.start != edge.end {
			edges = append(edges, edge)
		}
	}
	g.Edges = edges
	g.VertexNumber--
}

// KargerMinCut implements Krager's Contraction with repeatTime as repeating Time. The repeatTime should follow the O()
// inequality of the vertex number n: repeatTime = O(n^2 log n)
func KargerMinCut(g *Graph, repeatTime int) int {
	vertexNumber := g.VertexNumber;
	if vertexNumber <= 2 {
		return len(g.Edges)
	}

	minTestResult := len(g.Edges)

	ch := make(chan int)

	for i := 0; i < repeatTime; i++ {
		go func(g *Graph) {
			newGraph := g.CloneGraph()
			trialResult := kargerMinCutSingleRun(newGraph)
			newGraph = nil
			ch <- trialResult
		}(g)
	}

	for i := 0; i < repeatTime; i++ {
		trialResult := <-ch
		if trialResult < minTestResult {
			minTestResult = trialResult
		}
	}

	return minTestResult
}

func kargerMinCutSingleRun(g *Graph) int {
	for g.VertexNumber > 2 {
		r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		g.ContractEdge(r.Intn(len(g.Edges)))
		r = nil
	}
	return len(g.Edges);
}