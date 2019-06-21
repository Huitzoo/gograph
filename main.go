package main

import (
	"./graph"
)

func main() {
	/*
		t := tree.Tree{}

		t.Id = -1
		t.Content = "F"
		t.Father = nil

		for i, c := range "Tree" {
			t.Childs = append(t.Childs, &tree.Tree{i, string(c), nil, &t})
			//v = append(v, string(c))
		}

		child := t.Childs[0]
		for i, c := range "child" {
			child.Childs = append(child.Childs, &tree.Tree{i, string(c), nil, child})
		}
		fmt.Println("Father of all: ", t.Content)
		t.PrintTree()
	*/

	g := graph.Graph{}
	g.CreateGraph("struct1.txt", "no-ham")
	t := g.WideGraphNoHamiltonian("S", -1)
	t.PrintTree()
}
