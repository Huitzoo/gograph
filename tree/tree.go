package tree

import "fmt"

type Tree struct {
	Id      int
	Content string
	Childs  []*Tree
	Father  *Tree
}

func (t *Tree) DeleteSheet() {

}

func (t *Tree) PrintTree() {
	sheet := ""
	for _, child := range t.Childs {
		if len(child.Childs) != 0 {
			child.PrintTree()
		}
		sheet += child.Content
	}
	fmt.Println("Father: ", t.Content, "childs: ", sheet)
	return
}
