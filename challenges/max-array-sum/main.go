package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the maxSubsetSum function below.
//
// Space: O(n) which will be the max height of the call stack
//
// The interesting complication of this exercise is that the non-adjacent sets need to be of minimum size two.
// The dynamic programming concept is to divide the problem of doMaxSubsetSum (that is, to solve for no minimum size)
// into simpler subproblems.
// The outer dynamic problem (that is, to solve for minimum size two) must be solved in a way that reuses the memo.
//
// I struggled a lot with this exercise because first I didn't think of memoisation, and then didn't apply it to the
// outer problem.
func maxSubsetSum(arr []int32) int32 {
	if len(arr) < 3 {
		return 0
	}
	var mx int32
	memo := map[int]int32{}
	for i := 0; i < len(arr)-2; i++ {
		mx = max(mx, arr[i]+doMaxSubsetSum(arr, i+2, memo))
	}
	return mx
}

func doMaxSubsetSum(arr []int32, i int, memo map[int]int32) int32 {
	if i >= len(arr) {
		return 0
	}
	if _, ok := memo[i]; !ok {
		memo[i] = max(arr[i]+doMaxSubsetSum(arr, i+2, memo), doMaxSubsetSum(arr, i+1, memo))
	}
	return memo[i]
}

// Time: O(n)
// Space: O(1)
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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := maxSubsetSum(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
