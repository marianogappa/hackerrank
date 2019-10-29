package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the commonChild function below.
// Time: len(s1)*len(s2) -> O(n*m)
// Space: O(n*m)
func commonChild(s1 string, s2 string) int32 {
	// This exercise is known as longest common subsequence (LCS).
	// It can be solved by:
	// 1. if the last char on both is equal, add it to the len of the LCS.
	// 2. if the last char is not equal, then len(LCS) is the max of branching on each.
	// Different branches will end up on same substrings, so memoization is needed.
	// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
	//
	// I tried the normal DP approach with memoization, but the problem is that the memo keys
	// are too large (s1xs2 for every node). So I looked around and there's another approach
	// that doesn't need keys, and the complexity is more clear. The Wiki page has it.
	//
	// GOTCHA: longest common subsequence cannot be solved with traditional DP because memo
	// keys are too large. https://en.wikipedia.org/wiki/Longest_common_subsequence_problem

	memo := make([][]int32, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int32, len(s2)+1)
	}

	for i1 := 0; i1 < len(s1); i1++ {
		for i2 := 0; i2 < len(s2); i2++ {
			if s1[i1] == s2[i2] {
				memo[i1+1][i2+1] = 1 + memo[i1][i2]
			} else {
				memo[i1+1][i2+1] = max(memo[i1][i2+1], memo[i1+1][i2])
			}
		}
	}

	return memo[len(s1)][len(s2)]
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
