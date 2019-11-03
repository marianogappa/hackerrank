package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the maxMin function below.
//
// Time: O(n*logn)
// Space: O(1)
//
// Finding the minimum "polar distance" is least computationally expensive
// if the slice is sorted.
//
// Once it's sorted, one just traverses the slice in subslices of size k,
// storing the minimum distance between the first and last elements of
// each subslice.
func maxMin(k int32, arr []int32) int32 {
	// Time: O(n*logn)
	// Space: O(1)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// Time: O(n)
	// Space: O(1)
	// Find the minimum "polar distance".
	var m int32 = math.MaxInt32
	for i := 0; i < len(arr)-int(k)+1; i++ {
		m = min(m, arr[i+int(k)-1]-arr[i])
		// Quit early if we already achieved the minimum possible distance
		if m == 0 {
			break
		}
	}
	return m
}

func min(a, b int32) int32 {
	if a < b {
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

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := maxMin(k, arr)

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
