package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

// import "encoding/csv"
// import "math/rand"
//import "reflect"

func checkError (err error) {
	if err != nil {
		panic(err)
	}
}

/***
 * Vertex
 */

type Vertex struct {
	Id int
	Edges []*Edge
	Explored bool
}

type Edge struct {
	From *Vertex
	To *Vertex
}

func (v *Vertex) String () string { return fmt.Sprintf("%v", v.Edges) }
func (e *Edge) String () string { return fmt.Sprintf("(%d-%d)", e.From.Id, e.To.Id) }

func NewVertex (id int) *Vertex {
	return &Vertex {
		id,
		make([]*Edge, 0),
		false,
	}
}

func (v *Vertex) AddEdge (to *Vertex) {
	edge := Edge {v, to}
	v.Edges = append(v.Edges, &edge)
}

/***
 * Graph
 */
type Graph struct {
	Vertices map[int]*Vertex
}

func NewGraph () *Graph {
	return &Graph{
		make(map[int]*Vertex),
	}
}

func (g *Graph) NumOfVertices () int {
	return len(g.Vertices)

}

func (g *Graph) Vertex (id int) *Vertex {
	if _, ok := g.Vertices[id]; !ok {
		g.Vertices[id] = NewVertex(id)
	}
	return g.Vertices[id]
}

func (g *Graph) AddEdge (from int, to int) {
	g.Vertex(from).AddEdge(g.Vertex(to))
}

func (g *Graph) Transpose () *Graph {
	gr := NewGraph ()

	for _, v := range g.Vertices {
		for _, e := range v.Edges {
			gr.AddEdge(e.To.Id, e.From.Id)
		}

	}
	return gr
}

func loadGraph (file *os.File) *Graph {
	scanner := bufio.NewScanner(file)
	graph := NewGraph()

	for scanner.Scan() {
		line := scanner.Text()
		strings := strings.Split(line, " ")
		from, err := strconv.Atoi(strings[0]); checkError(err)
		to, err := strconv.Atoi(strings[1]); checkError(err)
		graph.AddEdge(from, to)
	}

	return graph
}

func (g *Graph) _DFSLoop (vertex *Vertex, stack *[]int) {
	vertex.Explored = true
	for _, edge := range vertex.Edges {
		if edge.To.Explored == false {
			g._DFSLoop (edge.To, stack)
		}
	}

	*stack = append(*stack, vertex.Id)
}

func (g *Graph) DFSLoop () {
	stack := make([]int,0)

	for i := 1; i <= g.NumOfVertices(); i++ {
		v := g.Vertex(i)
		if v.Explored == false {
			g._DFSLoop (v, &stack)
		}
	}

	gt := g.Transpose()
	sum := 0
	for j := len(stack) - 1; j >= 0 ; j-- {
		v := gt.Vertex(stack[j])
		if v.Explored == false {
			visited := make([]int, 0)
			gt._DFSLoop(v, &visited)
			fmt.Println(len(visited))
			sum = sum + len(visited)
		}
	}

	fmt.Printf("%d(sum) == %d(vertices)\n", sum, gt.NumOfVertices())
}

/***
 * Main
 */
func main () {
	g := loadGraph(os.Stdin)
	g.DFSLoop()
}
