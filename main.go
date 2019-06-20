package main

import (
	"fmt"

	"./graph"
)

func main() {

	g := graph.Graph{}
	g.CreateGraph("struct.txt", "0")
	g.DepthGraph("H", "H")
	fmt.Println(g.Walk)
}
