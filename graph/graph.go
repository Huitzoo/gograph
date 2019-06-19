package graph

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func errors(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type node struct {
	name string
}

type edge struct {
	cost  int
	nodes []node
}

type Graph struct {
	kind    string
	edg     []edge
	n_edg   int
	n_nodes int
}

func (g *Graph) CreateGraph(name string, kind string) {

	data, err := ioutil.ReadFile(name)
	nodes := []node{}
	edges := []edge{}

	errors(err)
	lines := strings.Split(string(data), "\n")
	g.n_nodes = len(strings.Replace(lines[0], " ", "", -1))

	for _, c := range lines[0] {
		if c != 32 {
			nodes = append(nodes, node{name: string(c)})
		}
	}

	for _, line := range lines[1 : len(lines)-1] {
		con := strings.Split(line, " ")
		e := edge{}
		for n := range nodes {
			if nodes[n].name == con[0] || nodes[n].name == con[1] {
				e.nodes = append(e.nodes, nodes[n])
			}
		}
		c, err := strconv.Atoi(con[2])
		errors(err)
		e.cost = c
		edges = append(edges, e)
	}
	g.edg = edges
	g.n_edg = len(edges)
	g.kind = kind

	//len(strings.Replace(lines[0], " ", "", -1))
}
func (g *Graph) DeepGraph() {

}
