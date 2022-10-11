package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

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
	res := shared.Handin1_ba(genomes, reads)

	shared.ToFile(res, "./testdata/h1_naive_results.txt")

	shared.SortFile("./testdata/test_result.txt")
	shared.SortFile("./testdata/h1_naive_results.txt")

	if !shared.CmpFiles("./testdata/test_result.txt", "./testdata/h1_naive_results.txt") {
		t.Errorf("files are not identical")
	}
}

func Test_CompareRandomResultsFromBAandSuffixTree(t *testing.T) {
	var rb strings.Builder
	num_of_n := 30
	num_of_m := 5

	genome, reads := shared.BuildSomeFastaAndFastq(num_of_n, num_of_m, 10, shared.AB, 42069)
	parsedGenomes := shared.GeneralParserStub(genome, shared.Fasta, num_of_n*num_of_m+1)
	parsedReads := shared.GeneralParserStub(reads, shared.Fastq, num_of_n*num_of_m+1)

	for _, gen := range parsedGenomes {
		s := gen.Rec
		if s[len(s)-1] != '$' {
			var sb strings.Builder
			sb.WriteString(s)
			sb.WriteRune('$')
			s = sb.String()
		}
		suffixTree := buildSuffixTree(s)
		for _, read := range parsedReads {
			matches := findoccurrences(suffixTree, read.Rec)
			for _, match := range matches {
				rb.WriteString(shared.SamStub(read.Name, gen.Name, match, read.Rec))
			}
		}
	}

	shared.ToFile(rb.String(), "./testdata/test_result_rand.txt")

	//generate (hopefully) identical sam file using last weeks BA method. Both files needs to be sorted
	res := shared.Handin1_ba(parsedGenomes, parsedReads)

	shared.ToFile(res, "./testdata/h1_naive_results_rand.txt")

	shared.SortFile("./testdata/test_result_rand.txt")
	shared.SortFile("./testdata/h1_naive_results_rand.txt")

	if !shared.CmpFiles("./testdata/test_result_rand.txt", "./testdata/h1_naive_results_rand.txt") {
		t.Errorf("files are not identical")
	}
}

func TestMakeDataFixN(t *testing.T) {
	csvFile, err := os.Create("./testdata/fixed_n_data.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)
	_ = csvwriter.Write([]string{"x_size", "naive", "naive_worst"})

	num_of_n := 8
	time_Naive := 0
	time_Naive_worst := 0

	for i := 1; i < 2; i++ {

		num_of_n *= 2
		num_of_m := 1
		genome, _ := shared.BuildSomeFastaAndFastq(num_of_n, num_of_m, 1, shared.English, 78)
		parsedGenomes := shared.GeneralParserStub(genome, shared.Fasta, num_of_n*num_of_m+1)
		//parsedReads := shared.GeneralParserStub(reads, shared.Fastq, num_of_n*num_of_m+1)

		for i := 0; i < 6; i++ {
			for _, gen := range parsedGenomes {
				s := gen.Rec
				if s[len(s)-1] != '$' {
					var sb strings.Builder
					sb.WriteString(s)
					sb.WriteRune('$')
					s = sb.String()
				}
				time_start := time.Now()
				buildSuffixTree(s)
				time_end := int(time.Since(time_start))
				time_Naive += time_end
				/*
					for _, read := range parsedReads {
						findoccurrences(suffixTree, read.Rec)
					}*/
				fmt.Println("NAIVE", int((time_Naive)))
				csvwriter.Flush()

			}

			genome2, _ := shared.BuildSomeFastaAndFastq(num_of_n, num_of_m, 1, shared.A, 78)
			parsedGenomes2 := shared.GeneralParserStub(genome2, shared.Fasta, num_of_n*num_of_m+1)

			for _, gen := range parsedGenomes2 {
				s := gen.Rec
				if s[len(s)-1] != '$' {
					var sb strings.Builder
					sb.WriteString(s)
					sb.WriteRune('$')
					s = sb.String()
				}
				time_start := time.Now()
				buildSuffixTree(s)
				time_end := int(time.Since(time_start))
				time_Naive_worst += time_end
				/*
					for _, read := range parsedReads {
						findoccurrences(suffixTree, read.Rec)
					}*/
				fmt.Println("NAIVE2", int((time_Naive_worst)))

			}
			_ = csvwriter.Write([]string{strconv.Itoa(num_of_n), strconv.Itoa(time_Naive), strconv.Itoa(time_Naive_worst)})
			csvwriter.Flush()
			time_Naive_worst = 0
			time_Naive = 0
		}

	}

}

func TestMakeDataSearchtime(t *testing.T) {
	csvFile, err := os.Create("./testdata/time_search.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)
	_ = csvwriter.Write([]string{"x_size", "naive4", "naive8"})

	num_of_n := 4
	time_Naive4 := 0
	time_Naive8 := 0
	num_of_m := 1

	for i := 1; i < 2; i++ {
		fmt.Println(num_of_n, num_of_m)
		num_of_n *= 2
		num_of_m *= 2
		genome, reads := shared.BuildSomeFastaAndFastq(num_of_n, num_of_m, 1, shared.A, 78)
		parsedGenomes := shared.GeneralParserStub(genome, shared.Fasta, num_of_n*num_of_m+1)
		parsedReads := shared.GeneralParserStub(reads, shared.Fastq, num_of_n*num_of_m+1)

		genome2, reads2 := shared.BuildSomeFastaAndFastq(num_of_n*2, num_of_m, 1, shared.A, 78)
		parsedGenomes2 := shared.GeneralParserStub(genome2, shared.Fasta, num_of_n*num_of_m+1)
		parsedReads2 := shared.GeneralParserStub(reads2, shared.Fastq, num_of_n*num_of_m+1)

		for i := 0; i < 10; i++ {
			for _, gen := range parsedGenomes {
				s := gen.Rec
				if s[len(s)-1] != '$' {
					var sb strings.Builder
					sb.WriteString(s)
					sb.WriteRune('$')
					s = sb.String()
				}
				suffixTree := buildSuffixTree(s)

				for _, read := range parsedReads {
					time_start := time.Now()
					findoccurrences(suffixTree, read.Rec)
					time_end := int(time.Since(time_start))
					time_Naive4 += time_end
				}
				fmt.Println("NAIVE", int((time_Naive4)))
				csvwriter.Flush()

			}

			for _, gen := range parsedGenomes2 {
				s := gen.Rec
				if s[len(s)-1] != '$' {
					var sb strings.Builder
					sb.WriteString(s)
					sb.WriteRune('$')
					s = sb.String()
				}
				suffixTree := buildSuffixTree(s)

				for _, read := range parsedReads2 {
					time_start := time.Now()
					findoccurrences(suffixTree, read.Rec)
					time_end := int(time.Since(time_start))
					time_Naive8 += time_end
				}
				fmt.Println("NAIVE", int((time_Naive8)))
				csvwriter.Flush()

			}
			_ = csvwriter.Write([]string{strconv.Itoa(num_of_m), strconv.Itoa(time_Naive4), strconv.Itoa(time_Naive8)})
			csvwriter.Flush()
			time_Naive4 = 0
			time_Naive8 = 0
		}
	}
}
