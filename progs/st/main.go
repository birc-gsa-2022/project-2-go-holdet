package main

import (
	"fmt"
	"strings"
)

// global variable storing the string used to build suffix tree
var x string

// tree structure
type Node struct {
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
		fmt.Println(i, len(y), "THE NEW IT")

		child := parent.Children[rune(y[i])]

		//case where we ended in a node and now have a mismatch
		if child == nil {
			fmt.Println("oh here boi")
			return parent, child, i, edgeLength(parent)
		}

		sf_len := slow_scan(child, y[i:])

		//case where we have a mismatch on some edge, we need both parent and child
		i += sf_len
		if sf_len < edgeLength(child) {
			return parent, child, i, sf_len
		}

		fmt.Println(sf_len, edgeLength(child))

		parent = child
		//case for when we match on patterns
		if i == len(y) {
			fmt.Println(parent, child, "huhhhh")
			return parent, child, i, sf_len
		}

	}
	return nil, nil, -1, -1
}

// return how many chars we match
func slow_scan(v *Node, y_part string) int {
	fmt.Println(y_part, v, x)
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

/*Builds a tree from some string */
func buildSuffixTree(x string) *Node {
	root := newNode(0, 0, nil)

	fmt.Println(x)
	for i := range x {
		cur := x[i:]
		parent, child, total, val := search(root, cur)
		if parent == nil {
			fmt.Println("oh no")
		}
		if child == nil {
			newNode(i+total, len(x), parent)
		} else {
			//case where we end on an edge
			splitEdge(parent, child, val, i+total)
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
	//root does not have a parent
	if parent != nil {
		parent.Children[rune(x[startIndex])] = &node
	}
	return &node
}

/* creates a new internal node between parent and child.
it also creates a new node branching (head)*/
func splitEdge(parent *Node, child *Node, splitIndex int, start_idx int) {
	fmt.Println("splitting edge")
	delete(parent.Children, rune(x[child.startIndex]))

	new_internal := newNode(child.startIndex, child.startIndex+splitIndex, parent)
	fmt.Println(start_idx+splitIndex, start_idx, splitIndex)
	s := newNode(start_idx, len(x), new_internal)

	child.startIndex = child.startIndex + splitIndex

	new_internal.Children[rune(x[child.startIndex])] = child

	fmt.Println("split:", parent, child, new_internal, s)

}

func findoccurrences(root *Node, y string) []int_tuple {
	parent_tree, _, _, _ := search(root, y)
	fmt.Println(parent_tree, "cringe")
	return BfOrder(parent_tree)

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
		fmt.Println(v, "oka")
		res := int_tuple{start: v.startIndex, end: v.endIndex}

		//we only want to leaves as results
		if len(v.Children) == 0 {
			result = append(result, res)
		} else {
			for _, child := range v.Children {
				queue.push(child)
			}
		}
	}
	return result
}

func main() {
	/*if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}*/
	x = "aaaaaaaaaaaa"
	x = "baa"
	x = "aababab"
	x = "abbabbafdsfdshacxzczdsffdsdfsd"
	x = "abbabaaabbabababxababxababaadsadsasafdsadsadshadbgsadasdbahdsahajhdbashjfbsahjdbashjdsabhdjhasbjshdabashjaaxxbaaabsaxxasxabx"
	x = "mississippi"
	x = "aaaaaaaaaaaa"
	if x[len(x)-1] != '$' {
		var sb strings.Builder
		sb.WriteString(x)
		sb.WriteRune('$')
		x = sb.String()
	}
	/*
		genome := os.Args[1]
		reads := os.Args[2]

		fmt.Println(shared.Todo(genome, reads))
	*/
	s_tree := buildSuffixTree(x)
	fmt.Println("lets go")
	matches := findoccurrences(s_tree, "a")

	fmt.Println(matches)
}
