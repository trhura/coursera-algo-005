package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

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
	Neighbors []*Edge
}

type Edge struct {
	Node1 *Vertex
	Node2 *Vertex
	Weight int
}

func (v *Vertex) String () string { return fmt.Sprintf("%v", v.Id) }
func (e *Edge) String () string { return fmt.Sprintf("(%d-%d-%d)", e.Node1.Id, e.Node2.Id, e.Weight) }

func NewVertex (id int) *Vertex {
	return &Vertex {
		id,
		make([]*Edge, 0),
	}
}

// func (e *Edge) OtherNode (v *Vertex) *Vertex {
// 	if v == e.Node1 {
// 		return e.Node2
// 	} else {
// 		return e.Node1
// 	}
// }

/***
 * Graph
 */
type Graph struct {
	Vertices map[int]*Vertex
	Edges []*Edge
}

func NewGraph () *Graph {
	return &Graph{
		make(map[int]*Vertex),
		make([]*Edge, 0),
	}
}

func (g *Graph) NumOfVertices () int {
	return len(g.Vertices)

}

func (g *Graph) NumOfEdges () int {
	return len(g.Edges)

}

func (g *Graph) Vertex (id int) *Vertex {
	if _, ok := g.Vertices[id]; !ok {
		g.Vertices[id] = NewVertex(id)
	}
	return g.Vertices[id]
}

func (g *Graph) AddEdge (n1 int, n2 int, weight int) {
	node1 := g.Vertex(n1)
	node2 := g.Vertex(n2)

	edge := Edge {node1, node2, weight}
	node1.Neighbors = append(node1.Neighbors, &edge)
	g.Edges = append(g.Edges, &edge)
}

func (g *Graph) ShortestPathsFromVertex (vid int) []int {
	distances := make ([]int, g.NumOfVertices())
	for i, _ := range distances {
		distances[i] = math.MaxInt64
		if i == vid - 1 {
			distances[i] = 0

		}
	}

	visited := make(map[int]bool)
	is_visited := func (node *Vertex) bool {
		_, ok := visited[node.Id-1]
		return ok
	}

	// unvisited := make([]*Vertex, 0)
	// for _, node := range g.Vertices {
	// 	unvisited = append(unvisited, node)
	// }

	make_visited := func (node *Vertex) {
		visited[node.Id-1] = true
		// unvisited[node.Id-1] = unvisited[len(unvisited)-1]
		// unvisited = unvisited[:len(unvisited)-1]
	}

	cvid := vid
	for {
		current := g.Vertex(cvid)
		make_visited(current)

		smallest := math.MaxInt64
		smallestnode := current

		for _, edge := range g.Edges {
			if is_visited (edge.Node1) && !is_visited (edge.Node2) {
				neighbor := edge.Node2
				if distances[current.Id-1]  == math.MaxInt64 {panic("MAXINT")}
				tdis := distances[current.Id-1] + edge.Weight

				if tdis < distances[neighbor.Id-1] {
					distances[neighbor.Id-1] = tdis
				}

				if distances[neighbor.Id-1] < smallest {
					smallestnode = neighbor
					smallest = distances[neighbor.Id-1]
				}
			}
		}

		//fmt.Println(cvid, distances)
		if smallest == math.MaxInt64 {
			break
		}

		if smallestnode == current { panic("NOCURRENT")}
		cvid = smallestnode.Id
	}

	return distances
}


func loadGraph (file *os.File) *Graph {
	scanner := bufio.NewScanner(file)
	graph := NewGraph()

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, "\t")

		node1, err := strconv.Atoi(splitted[0]); checkError(err)

		for _, x := range splitted[1:] {
		 	edge := strings.Split(x, ",")
			if len(edge) == 2 {
				node2, err := strconv.Atoi(edge[0]); checkError(err)
				weight, err := strconv.Atoi(edge[1]); checkError(err)
				graph.AddEdge (node1, node2, weight)
			}
		}
	}

	return graph
}

/***
 * Main
 */
func main () {
	g := loadGraph(os.Stdin)
	fmt.Println(g.NumOfVertices(), g.NumOfEdges())
	from := 1

	distances := g.ShortestPathsFromVertex(from)
	for i, d := range distances {
		fmt.Printf("%d -> %d = %d\n", from, i+1, d)
	}

	fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n",
		distances[7-1],
		distances[37-1],
		distances[59-1],
		distances[82-1],
		distances[99-1],
		distances[115-1],
		distances[133-1],
		distances[165-1],
		distances[188-1],
		distances[197-1])
}
