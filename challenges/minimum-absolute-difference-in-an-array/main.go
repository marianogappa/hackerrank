package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the minimumAbsoluteDifference function below.
// Time: O(n*logn)
// Space: O(1)
func minimumAbsoluteDifference(arr []int32) int32 {
	// Time: O(n*logn)
	// Space: O(1)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// Time: O(n)
	// Space: O(1)
	minimumDifference := int32(1<<31 - 1)
	for i := 1; i < len(arr); i++ {
		curDiff := absDiff(arr[i-1], arr[i])
		if curDiff < minimumDifference {
			minimumDifference = curDiff
		}
	}
	return minimumDifference
}

func absDiff(a, b int32) int32 {
	if a-b < 0 {
		return b - a
	}
	return a - b
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

	result := minimumAbsoluteDifference(arr)

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
