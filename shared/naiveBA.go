package shared

import "strings"

func Handin1_ba(genome string, reads string) string {
	var rb strings.Builder
	parsedGenomes := GeneralParser(genome, Fasta)

	parsedReads := GeneralParser(reads, Fastq)

	for _, read := range parsedReads {
		for _, gen := range parsedGenomes {
			matches := naive(gen.Rec, read.Rec)
			for _, match := range matches {
				rb.WriteString(SamStub(read.Name, gen.Name, match, read.Rec))
			}
		}
	}
	return rb.String()
}

func naive(x string, p string) (matches []int) {
	if p == "" {
		return matches
	}
outer_loop:
	for i := 0; i < len(x)-len(p)+1; i++ {
		for j, char := range []byte(p) {
			if char != x[i+j] {
				continue outer_loop
			}
		}
		matches = append(matches, i)
	}

	return matches
}
