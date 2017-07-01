package week4

import (
	"math/rand"
	"time"
)

// graph is a simple graph struct to implement Krager's Algorithm
type graph struct {
	VertexNumber int
	Edges []edge
}

type edge struct {
	start int
	end int
}

// NewGraph create a graph and returns reference to the created graph
func NewGraph() *graph {
	VertexNumber := 0
	Edges := make([]edge, 0)
	return &graph{VertexNumber, Edges}
}

// AddEdge will add an edge starting with start and ending with end
func (g *graph) AddEdge (start, end int) {
	newEdge := edge{start, end}
	g.Edges = append(g.Edges, newEdge)
}

// AddVertex simply bump the vertexNumber of a graph
func (g *graph) AddVertex() {
	g.VertexNumber++
}

// CloneGraph does a deep copy of the graph
func (g *graph) CloneGraph() *graph {
	VertexNumber := g.VertexNumber
	Edges := make([]edge, len(g.Edges))
	copy(Edges, g.Edges)
	return &graph{
		VertexNumber,
		Edges,
	}
}

// ContractEdge implements Karger's Contraction. It will contract the edge with index i, remove the endpoint of edge(i)
// and rewrite all the edges related to the endpoint of edge(i)
func (g *graph) ContractEdge(i int) {
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
func KargerMinCut(g *graph, repeatTime int) int {
	vertexNumber := g.VertexNumber;
	if vertexNumber <= 2 {
		return len(g.Edges)
	}

	minTestResult := len(g.Edges)

	ch := make(chan int)

	for i := 0; i < repeatTime; i++ {
		go func(g *graph) {
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

func kargerMinCutSingleRun(g *graph) int {
	for g.VertexNumber > 2 {
		r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		g.ContractEdge(r.Intn(len(g.Edges)))
		r = nil
	}
	return len(g.Edges);
}