package main

import (
	"fmt"
	"os"
	"strings"

	"birc.au.dk/gsa/shared"
)

// global variable storing the string used to build suffix tree
var x string

// tree structure
type Node struct {
	Label      int
	Children   map[rune]*Node
	startIndex int
	endIndex   int
}

/* should take root as first parameter, string to search for as second param.
Returns parent, child, and how far it got in total. Lastly we also return
how many steps on the current edge.
*/
func search(v *Node, y string) (*Node, *Node, int, int) {
	i := 0
	parent := v

	for i < len(y) {
		child := parent.Children[rune(y[i])]

		//case where we ended in a node and now have a mismatch
		if child == nil {
			return parent, child, i, edgeLength(parent)
		}
		sf_len := slow_scan(child, y[i:])

		//case where we have a mismatch on some edge, we need both parent and child
		i += sf_len
		if sf_len < edgeLength(child) {
			return parent, child, i, sf_len
		}

		//case for when we match on patterns
		if i == len(y) {
			return parent, child, i, sf_len
		}

		parent = child

	}
	return nil, nil, -1, -1
}

// return how many chars we match
func slow_scan(v *Node, y_part string) int {
	suffix := x[v.startIndex:v.endIndex]

	for i, char := range []byte(suffix) {

		if i == len(y_part) {
			return i
		}
		if char != y_part[i] {
			return i
		}
	}
	return len(suffix)
}

/*Builds a tree from some string */
func buildSuffixTree(x string) *Node {
	root := newNode(0, 0, nil, 0)

	for i := range x {
		cur := x[i:]
		parent, child, total, val := search(root, cur)
		if parent == nil {
			fmt.Println("this case is not good")
		}
		if child == nil {
			newNode(i+total, len(x), parent, i)
		} else {
			//case where we end on an edge
			splitEdge(parent, child, val, i+total, i)
		}
		BfOrder(root)
	}
	return root
}

func edgeLength(node *Node) int {
	return node.endIndex - node.startIndex
}

func newNode(startIndex int, endIndex int, parent *Node, label int) *Node {
	node := Node{startIndex: startIndex, endIndex: endIndex, Children: make(map[rune]*Node), Label: label}
	//root does not have a parent
	if parent != nil {
		parent.Children[rune(x[startIndex])] = &node
	}
	return &node
}

/* creates a new internal node between parent and child.
it also creates a new node branching (head)*/
func splitEdge(parent *Node, child *Node, splitIndex int, start_idx int, label int) {
	delete(parent.Children, rune(x[child.startIndex]))

	new_internal := newNode(child.startIndex, child.startIndex+splitIndex, parent, parent.Label)
	newNode(start_idx, len(x), new_internal, label)

	child.startIndex = child.startIndex + splitIndex

	new_internal.Children[rune(x[child.startIndex])] = child

}

func findoccurrences(root *Node, y string) []int {
	parent, child, l, split := search(root, y)
	if l == len(y) {
		//if we end in a node
		if split == 0 {
			return BfOrder(parent)
		}
		//if we end on an edge
		return BfOrder(child)
	}
	//no match
	return []int{}
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
func BfOrder(v *Node) []int {
	queue := make(TreeQueue, 0)
	result := make([]int, 0)

	if v == nil {
		return result
	}

	queue.push(v)
	for len(queue) > 0 {
		//dequeue
		v = queue.popOrNil()
		//we only want to leaves as results
		if len(v.Children) == 0 {
			result = append(result, v.Label)
		} else {
			for _, child := range v.Children {
				queue.push(child)
			}
		}
	}
	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	genomes := shared.GeneralParser(os.Args[1], shared.Fasta)
	reads := shared.GeneralParser(os.Args[2], shared.Fastq)

	for _, gen := range genomes {
		s := gen.Rec
		if s[len(s)-1] != '$' {
			var sb strings.Builder
			sb.WriteString(s)
			sb.WriteRune('$')
			x = sb.String()
		}
		suffixTree := buildSuffixTree(s)
		for _, read := range reads {
			matches := findoccurrences(suffixTree, read.Rec)
			for _, match := range matches {
				shared.Sam(read.Name, gen.Name, match, read.Rec)
			}
		}
	}

}
