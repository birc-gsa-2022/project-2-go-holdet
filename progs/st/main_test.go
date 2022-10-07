package main

import (
	"strings"
	"testing"

	"birc.au.dk/gsa/shared"
)

/* test that compares results from our exact pattern matching with BA from handin 1 with
our Suffix Tree approach. Files generated are not in the same order, so the results are
both compared the same way before comparison */
func Test_CompareResultsFromBAandSuffixTree(t *testing.T) {
	var rb strings.Builder

	genomes := shared.GeneralParser("./testdata/genome.fa", shared.Fasta)
	reads := shared.GeneralParser("./testdata/reads.fq", shared.Fastq)

	for _, gen := range genomes {
		s := gen.Rec
		if s[len(s)-1] != '$' {
			var sb strings.Builder
			sb.WriteString(s)
			sb.WriteRune('$')
			s = sb.String()
		}
		suffixTree := buildSuffixTree(s)
		for _, read := range reads {
			matches := findoccurrences(suffixTree, read.Rec)
			for _, match := range matches {
				rb.WriteString(shared.SamStub(read.Name, gen.Name, match, read.Rec))
			}

		}
	}

	shared.ToFile(rb.String(), "./testdata/test_result.txt")

	//generate (hopefully) identical sam file using last weeks BA method. Both files needs to be sorted
	res := shared.Handin1_ba("./testdata/genome.fa", "./testdata/reads.fq")

	shared.ToFile(res, "./testdata/h1_naive_results.txt")

	shared.SortFile("./testdata/test_result.txt")
	shared.SortFile("./testdata/h1_naive_results.txt")

	if !shared.CmpFiles("./testdata/test_result.txt", "./testdata/h1_naive_results.txt") {
		t.Errorf("files are not identical")
	}
}
