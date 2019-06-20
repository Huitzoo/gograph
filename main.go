package main

import (
	"fmt"

	"./graph"
)

func main() {

	g := graph.Graph{}
	g.CreateGraph("struct.txt", "0")
	g.DeepGraph("A", "H")
	fmt.Println(g.Walk)
}
