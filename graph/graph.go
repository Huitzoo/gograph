package graph

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"../tree"
)

func bug(err error) {
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
}

func (g *Graph) CreateGraph(name string, kind string) {

	data, err := ioutil.ReadFile(name)
	nodes := node{}
	nodes.name = make(map[string]int)
	bug(err)
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
	for _, line := range lines[1:len(lines)] {
		con := strings.Split(line, " ")
		ed := edge{}
		ed.nodes = append(ed.nodes, con[0])
		ed.nodes = append(ed.nodes, con[1])
		c, err := strconv.Atoi(con[2])

		bug(err)
		ed.cost = c
		g.edges = append(g.edges, ed)

	}
	g.nodes = nodes
	g.n_edg = len(g.edges)
	g.kind = kind
	//len(strings.Replace(lines[0], " ", "", -1))
}

func (g *Graph) DepthGraphHamiltonian(start string, end string) string {
	//g.Walk = ""
	walk := ""
	check := make([]bool, len(g.nodes.name))
	var stack []string
	flag := 0
	/*flag is a flag to use for first iteration*/
	for flag == 0 || len(stack) != 0 {
		flag = 1
		walk += start + " "
		check[g.nodes.name[start]] = true
		if g.kind != "ham" {
			bug(errors.New("Kind of Graph doesn't support this type of graph"))
		}
		for _, edge := range g.edges {
			if edge.nodes[0] == start {
				if !check[g.nodes.name[edge.nodes[1]]] {
					check[g.nodes.name[edge.nodes[1]]] = true
					stack = append(stack, edge.nodes[1])
				}
			}
		}
		x := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		start = x
	}
	walk += start
	return walk
}

func (g *Graph) WideGraphHamiltonian(start string, end string) string {

	walk := ""
	check := make([]bool, len(g.nodes.name))
	var stack []string
	flag := 0
	/*flag is a flag to use for first iteration*/
	for flag == 0 || len(stack) != 0 {
		flag = 1
		walk += start + " "
		check[g.nodes.name[start]] = true
		if g.kind != "ham" {
			bug(errors.New("Kind of Graph doesn't support this type of graph"))
		}
		for _, edge := range g.edges {
			if edge.nodes[0] == start {
				if !check[g.nodes.name[edge.nodes[1]]] {
					check[g.nodes.name[edge.nodes[1]]] = true
					stack = append(stack, edge.nodes[1])
				}
			}
		}
		x := stack[0]
		stack = stack[1:len(stack)]
		start = x
	}
	walk += start
	return walk

}
func (g *Graph) checkHamiltonianGraph() {

}

func (g *Graph) WideGraphNoHamiltonian(start string, i int) *tree.Tree {
	t := tree.Tree{i, start, nil, nil}

	//check := make([]bool, len(g.nodes.name))
	//var stack []string
	var flag int = 0

	//flag := 0
	/*flag is a flag to use for first iteration*/
	for {

		//check[g.nodes.name[start]] = true
		if g.kind != "no-ham" {
			bug(errors.New("Kind of Graph doesn't support this type of graph"))
		}
		for _, edge := range g.edges {
			if edge.nodes[0] == start {
				//if !check[g.nodes.name[edge.nodes[1]]] {
				t.Childs = append(t.Childs, &tree.Tree{i, string(edge.nodes[1]), nil, &t})
				//check[g.nodes.name[edge.nodes[1]]] = true
				g.WideGraphNoHamiltonian(edge.nodes[0], i)
				i += 1
				//}
				flag = 1
			}
		}
		if flag != 1 {
			return nil
		}
		/*
			if len(stack) == 0 {
				break
			}
			fmt.Println(stack)
			x := stack[0]
			stack = stack[1:len(stack)]
			start = x
		*/
	}
	return &t
}
