package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the pairs function below.
// Time: O(n^2) because in the worst case all numbers are equal, so they all pair with each other
// Space: O(n) because the map stores all numbers once
func pairs(k int32, arr []int32) int32 {
	m := map[int32][]int{}
	for i, n := range arr {
		absn := abs(n)
		m[absn] = append(m[absn], i)
	}
	var total int32
	for i, n := range arr {
		total += countPairs(m[n-k], n, i)
		if n-k != k+n {
			total += countPairs(m[k+n], n, i)
		}
	}
	return total
}

// Time: O(n), where n == len(indexes)
// Space: O(1)
func countPairs(indexes []int, n int32, i int) int32 {
	var total int32
	for _, j := range indexes {
		if j > i { // the second index must be higher to prevent double counting
			total++
		}
	}
	return total
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := pairs(k, arr)

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
