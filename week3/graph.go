package main

import "fmt"
import "io"
import "os"
import "encoding/csv"
import "strconv"
import "math/rand"
//import "reflect"

/***
 * Vertex
 */

type Vertex struct {
	id int
	Neighbors []int
}

func NewVertex (id int) *Vertex {
	return &Vertex {
		id,
		make([]int, 0),
	}
}

func (v Vertex) Copy () *Vertex {
	neighs := make([]int, len(v.Neighbors))
	copy(neighs, v.Neighbors)
	return &Vertex {
		v.id,
		neighs,
	}
}

/***
 * Graph
 */

type Graph struct {
	vertices map[int]*Vertex
}

func checkError (err error) {
	if err != nil {
		panic(err)
	}
}

func NewGraph () *Graph {
	return &Graph{
		make(map[int]*Vertex),
	}
}

func (g Graph) Copy() *Graph {
	gs := NewGraph()
	for k, v := range g.vertices {
		gs.vertices[k] = v.Copy()
	}
	return gs
}

func (g Graph) AddVertex (v *Vertex) {
	g.vertices[v.id] = v

}

func (g Graph) GetVertices () ([]int, []*Vertex) {
	ids := make([]int, 0)
	vertices := make([]*Vertex, 0)
	for k, v := range g.vertices {
		//fmt.Println(k, v)
		ids = append(ids, k)
		vertices = append(vertices, v)
	}
	return ids, vertices
}

func (g Graph) GetRandomVertex () *Vertex {
	ids, _ := g.GetVertices()
	rid :=  ids[rand.Intn(len(ids))]
	return g.vertices[rid]
}

func (g Graph) GetRandomNeigh (v *Vertex) *Vertex {
	rid := v.Neighbors[rand.Intn(len(v.Neighbors))]
	return g.vertices[rid]
}

func (g Graph) RemoveRandomEdge () {
	rv := g.GetRandomVertex() // first random vertex
	rnv := g.GetRandomNeigh(rv) // second random vertex

	// append second vertex neigh to first vertex
	rv.Neighbors = append(rv.Neighbors, rnv.Neighbors...)


	// replace refs to second vertext to first one
	for _, v := range g.vertices {
		for i, n := range v.Neighbors {
			if n == rnv.id {
				v.Neighbors[i] = rv.id
			}
		}
	}

	// remove self-loops
	newneigh := make([]int, 0)
	for _, n:= range rv.Neighbors {
		if n != rv.id {
			newneigh = append(newneigh, n)
		}
	}
	rv.Neighbors = newneigh

	// remove second vertex
	delete(g.vertices, rnv.id)
}

func (g Graph) Contract () int {
	//g.PrintGraph()
	for len(g.vertices) > 2 {
		g.RemoveRandomEdge()
		//g.PrintGraph()
	}

	ids, _ := g.GetVertices()
	if (len(ids) != 2 ) {
		panic("Should left with exactly 2 vertices.")
	}

	if len(g.vertices[ids[0]].Neighbors) != len(g.vertices[ids[1]].Neighbors) {
		fmt.Println(g.vertices[ids[0]], g.vertices[ids[1]])
		panic("Should have same edges.")
	}

	return len(g.vertices[ids[0]].Neighbors)
}

func LoadGraph (r io.Reader) *Graph {
	reader := csv.NewReader(r)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	checkError(err)

	g := NewGraph()
	for _, row := range records {
		vertexid, err := strconv.Atoi(row[0]); checkError(err)
		vertex := NewVertex(vertexid)

		for _, col := range row[1:] {
			if col == "" { continue }
			nid, err := strconv.Atoi(col); checkError(err)
			vertex.Neighbors = append(vertex.Neighbors, nid)
		}
		g.AddVertex(vertex)
	}

	return g;
}

func (g Graph) PrintGraph () {
	for k, v := range g.vertices {
		fmt.Println(k, v)
	}
	fmt.Println()
}

/***
 * Main
 */
func main () {
	g := LoadGraph(os.Stdin)

	min := 1000000000000
	for i := 0; i < 10000 ; i++ {
		gs := g.Copy()
		cut := gs.Contract()
		if cut < min  {
			min = cut
		}
		//fmt.Println("------------------------------")
	}

	fmt.Println("\nMin cut = ", min, "\n")
}
