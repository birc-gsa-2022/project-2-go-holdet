[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-c66648af7eb3fe8bc4f294546bfd86ef473780cde1dea487d3c4ff354943c9ae.svg)](https://classroom.github.com/online_ide?assignment_repo_id=8749050&assignment_repo_type=AssignmentRepo)
# Project 2: Suffix tree construction

You should implement a suffix tree construction algorithm. You can choose to implement the naive O(n²)-time construction algorithm as discussed in class or McCreight’s O(n) construction algorithm. After that, implement a search algorithm (similar to slow-scan) for finding all occurrences of a pattern. This algorithm should run in O(m+z) where m is the length of the pattern and z the number of occurrences.

Write a program, `st` using the suffix tree exact pattern search algorithm (similar to slow-scan) to report all indices in a string where a given pattern occurs. 

The program should take the same options as in project 1: `st genome.fa reads.fq`. The program should output (almost) the same SAM file. Because a search in a suffix tree is not done from the start to the end of the string the output might be in a different order, but if you sort the output from the previous project and for this program, they should be identical.

The code is set up to build the tool `st` once you provide the details, and then you can install it at the root with

```bash
> GOBIN=$PWD go install ./...
```

## Evaluation

Implement the tool `st` that does exact pattern matching using a suffix tree. Test it to the best of your abilities, and then fill out the report below.

# Report

## Specify if you have used a linear time or quadratic time algorithm.

We did not implement McCreight's Algorithm. We only had time to implement the quadratic time algorithm.


## Insights you may have had while implementing and comparing the algorithms.

We saw that the Search function could be implemented in a way where it could be used both for building the suffix tree and for searching for patterns later on.

It was also interesting to to implement our own data structure and see how this structure did computations notably slower than the algorithms in the previous assignment (Even when having the same time complexity)


## Problems encountered if any.

We had a some problems with getting the algorithm to work in the beginning. The issues primarily snug into the code when building the suffix tree and when we inserted nodes or splitted edges and had to add and change pointers, which we did wrong for some time. 
It was not too

## Correctness

In order to verify the correctness of our suffix tree implementation we verified with our naive border array algorithm from the previous project.
We tested our implementation on some selected input (files ![](./progs/st/testdata/genome.fa), ![](./progs/st/testdata/reads.fq) ) as well as random data generated from different size alphabets (A, AB, ACGT, English). Sam files generated from the suffix tree are not recieved in a specified order, which meant that the two Sam files first were sorted before we could compare that the files were identical. The tests can be found in the ![](./progs/st/main_test.go) file.

## Running time

Our implementation of the build suffix tree operation runs in O(n²) time.
We have conducted an exeriment that shows this. The worst case behaivour can be found by using the alphabet A*, since we guarantee maximal comparisons per inserted suffix - we always have to compare the entire suffix we insert.



*Describe experiments that verifies that your implementation of `st` uses no more time than O(n) or O(n²) (depending on the algorithm) for constructing the suffix tree and no more than O(m) for searching for a given read in it. Remember to explain your choice of test data. What are “best” and “worst” case inputs?*

*If you have graphs that show the running time--you probably should have--you can embed them here like we did in the previous project.*

