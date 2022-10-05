package main

import (
	"fmt"
	"os"

	// Directories in the root of the repo can be imported
	// as long as we pretend that they sit relative to the
	// url birc.au.dk/gsa, like this for the example 'shared':
	"birc.au.dk/gsa/shared"
)

//global variable storing the string used to build suffix tree
var x string

//tree structure
type Node struct {
	suffix     string
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

		if sf_len < w.endIndex-w.startIndex {
			return w, sf_len
		}
		v = w
		l += sf_len
	}
	return v, v.endIndex - v.startIndex
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

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	x = "placeholder"
	genome := os.Args[1]
	reads := os.Args[2]
	fmt.Println(shared.Todo(genome, reads))
}
