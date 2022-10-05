package main

import (
	"fmt"
)

//global variable storing the string used to build suffix tree
var x string

//tree structure
type Node struct {
	children   map[rune]*Node
	startIndex int
	endIndex   int
}

//should take root as first parameter, string to search for as second param.
func search(v *Node, y string) (*Node, int) {
	l := 0

	for l < len(y) {

		w := v.children[rune(y[l])]

		sf_len := suffix(w, y[l:])

		if sf_len < edgeLength(w) {
			return w, sf_len
		}
		v = w
		l += sf_len
	}
	return v, edgeLength(v)
}

//return how many chars we match
func suffix(v *Node, y_part string) int {
	suffix := x[v.startIndex:v.endIndex]

	for i, char := range []byte(suffix) {

		if i == len(y_part) {
			return i
		}
		if char != suffix[i] {
			return i
		}

	}
	return len(suffix)
}

func buildSuffixTree(x string) *Node {
	root := insertNode(0, 0, nil)

	for i := range x {
		v, val := search(&root, x[i:])
		if val == edgeLength(v) {
			insertNode(v.endIndex, len(x), v)
		}
	}
	return &root
}

func edgeLength(node *Node) int {
	return node.endIndex - node.startIndex
}

func insertNode(startIndex int, endIndex int, parent *Node) Node {
	node := new(Node)
	node.endIndex = endIndex
	node.startIndex = startIndex
	node.children = make(map[rune]*Node)
	parent.children[rune(x[startIndex])] = node
	return *node
}

func splitEdge(parent *Node, child *Node, splitIndex int, mismatch rune) {

	//clear children (there is only 1)
	for k := range parent.children {
		delete(parent.children, k)
	}

	new_internal := insertNode(parent.endIndex, splitIndex, parent)
	new_leaf := insertNode(splitIndex, len(x), &new_internal)

	parent.children[rune(splitIndex)+rune(parent.startIndex)] = &new_internal

	new_internal.children[rune(parent.endIndex)-rune(splitIndex)] = child
	new_internal.children[mismatch] = &new_leaf

	parent.endIndex = splitIndex
	child.startIndex = splitIndex

}

func main() {
	/*if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}*/
	x = "monke"
	/*
		genome := os.Args[1]
		reads := os.Args[2]

		fmt.Println(shared.Todo(genome, reads))
	*/
	fmt.Println(buildSuffixTree(x))
}
