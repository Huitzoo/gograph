package graph

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func e(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type node struct {
	name map[string]int
}

type edge struct {
	cost  int
	nodes []string
}

type Graph struct {
	kind    string
	edges   []edge
	nodes   node
	n_edg   int
	n_nodes int
	Walk    string
	check   []bool
	stack   []string
}

/*
func (g *Graph) findIndex(a string) int {
	for i, n := range g.nodes {
		if n.name == a {
			return i
		}
	}
	return -1
}
*/
func (g *Graph) CreateGraph(name string, kind string) {

	data, err := ioutil.ReadFile(name)
	nodes := node{}
	nodes.name = make(map[string]int)
	e(err)
	lines := strings.Split(string(data), "\n")
	g.n_nodes = len(strings.Replace(lines[0], " ", "", -1))
	i := 0
	/*Create nodes*/
	for _, c := range lines[0] {
		if c != 32 {
			nodes.name[string(c)] = i
			//nodes = append(nodes, node{name: string(c)})
			i += 1
		}
	}
	/*Create edge*/
	for _, line := range lines[1 : len(lines)-1] {
		con := strings.Split(line, " ")
		ed := edge{}
		ed.nodes = append(ed.nodes, con[0])
		ed.nodes = append(ed.nodes, con[1])
		c, err := strconv.Atoi(con[2])
		e(err)
		ed.cost = c
		g.edges = append(g.edges, ed)

	}
	g.nodes = nodes
	g.n_edg = len(g.edges)
	g.kind = kind
	g.Walk = ""
	g.check = make([]bool, len(g.nodes.name))
	//len(strings.Replace(lines[0], " ", "", -1))
}

func (g *Graph) DeepGraph(start string, end string) {

	//Check is a list to check if you passed in node
	flag := 0
	/*flag is a flag to use for first iteration*/
	for flag == 0 || len(g.stack) != 0 {
		flag = 1
		g.Walk += start + " "
		g.check[g.nodes.name[start]] = true
		if g.kind != "0" {
			e(errors.New("Kind of Graph doesn't support this type of graph"))
		}

		for _, edge := range g.edges {
			if edge.nodes[0] == start {
				if !g.check[g.nodes.name[edge.nodes[1]]] {
					g.check[g.nodes.name[edge.nodes[1]]] = true
					g.stack = append(g.stack, edge.nodes[1])
				}
			}
		}
		x, stack := g.stack[len(g.stack)-1], g.stack[:len(g.stack)-1]
		g.stack = stack
		start = x
	}
	g.Walk += start
}
