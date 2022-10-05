package main

import (
	"fmt"
)

// global variable storing the string used to build suffix tree
var x string

// tree structure
type Node struct {
	Children   map[rune]*Node
	startIndex int
	endIndex   int
}

// should take root as first parameter, string to search for as second param.
func search(v *Node, y string) (*Node, *Node, int) {
	l := 0
	parent := v

	fmt.Println()
	for l < len(y) {
		fmt.Println(parent)
		fmt.Println(l)

		child := parent.Children[rune(y[l])]

		//case where we ended in a node and now have a mismatch
		if child == nil {
			return parent, child, edgeLength(parent)
		}

		sf_len := suffix(child, y[l:])

		//case where we have a mismatch on some edge, we need both parent and child
		if sf_len < edgeLength(child) {
			return parent, child, sf_len
		}
		fmt.Println(sf_len, edgeLength(child))
		parent = child
		l += sf_len
	}

	//this case should not be reached
	fmt.Println("DISASTER")
	return nil, nil, -1
}

// return how many chars we match
func suffix(v *Node, y_part string) int {
	suffix := x[v.startIndex:v.endIndex]

	for i, char := range []byte(suffix) {

		if i == len(y_part) {
			return i
		}
		if char != y_part[i] {
			return i
		}
		fmt.Println(string(char), string(suffix[i]))
		fmt.Println(y_part, suffix)

	}
	return len(suffix)
}

func buildSuffixTree(x string) *Node {
	root := newNode(0, 0, nil)

	for i := range x {
		cur := x[i:]
		fmt.Println(cur, len(x)-i, " kekk")
		parent, child, val := search(root, cur)
		if child == nil {
			fmt.Print("HUHUHUHUHUHUHUHU")
			fmt.Println(parent)
			fmt.Println(parent.endIndex, i)
			newNode(parent.endIndex+i, len(x), parent)
		} else {
			//case where we end on an edge
			splitEdge(parent, child, val, i)
		}
		BfOrder(root)
	}
	return root
}

func edgeLength(node *Node) int {
	return node.endIndex - node.startIndex
}

func newNode(startIndex int, endIndex int, parent *Node) *Node {
	node := Node{startIndex: startIndex, endIndex: endIndex}
	node.Children = make(map[rune]*Node)
	fmt.Println("HOLDO")
	//root does not have a parent
	if parent != nil {
		fmt.Println(startIndex, "help")
		parent.Children[rune(x[startIndex])] = &node
	}
	return &node
}

func splitEdge(parent *Node, child *Node, splitIndex int, start_idx int) {
	mismatch := parent.endIndex + splitIndex + start_idx

	//delete current child because we want to squeeze a new node in between parent and child
	delete(parent.Children, rune(child.startIndex))

	//build new internal node on split, let parent be parent
	new_internal := newNode(child.startIndex, child.startIndex+splitIndex, parent)

	//build the new edge (head), let parent be new_internal
	newNode(mismatch, len(x), new_internal)

	//update length of child
	child.startIndex = child.startIndex + splitIndex

	//add child as a child to new_internal
	new_internal.Children[rune(child.startIndex)] = child

}

type int_tuple struct {
	start int
	end   int
}

//some wrapper for queue inspired from TA feedback
type TreeQueue []*Node

func (t *TreeQueue) push(v *Node) {
	*t = append(*t, v)
}
func (t *TreeQueue) popOrNil() *Node {
	if len(*t) == 0 {
		return nil
	}
	v := (*t)[0]
	*t = (*t)[1:]
	return v
}

// Do a breadth-order traversal of v and output the
// values in the tree.
func BfOrder(v *Node) []int_tuple {
	queue := make(TreeQueue, 0)
	result := []int_tuple{}

	if v == nil {
		return result
	}

	queue.push(v)
	for len(queue) > 0 {
		//dequeue
		v = queue.popOrNil()
		fmt.Println(v)
		res := int_tuple{start: v.startIndex, end: v.endIndex}
		result = append(result, res)

		for _, child := range v.Children {
			queue.push(child)
		}
	}
	return result
}

func main() {
	/*if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}*/
	x = "aababab$"
	/*
		genome := os.Args[1]
		reads := os.Args[2]

		fmt.Println(shared.Todo(genome, reads))
	*/
	fmt.Println(buildSuffixTree(x))
}
