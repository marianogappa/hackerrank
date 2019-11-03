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

// Complete the getMinimumCost function below.
// Time: O(n*logn)
// Space: O(1)
//
// First sort the prices descending, because it's cheaper to buy the most expensive
// flowers without multipliers.
//
// The k friends will buy flowers in i iterations of k purchases (first for-loop).
// The second for-loop calculates the sum of the batch.
//
// Careful not to exceed the len of the slice in the second for-loop.
func getMinimumCost(k int32, c []int32) int32 {
	// Time: O(n*logn)
	// Space: O(1)
	sort.Slice(c, func(i, j int) bool { return c[i] > c[j] })

	// Time: O(n)
	// Space: O(1)
	var total int32
	for i := 1; i <= int(math.Ceil(float64(len(c))/float64(k))); i++ {
		var acc int32
		for j := (i - 1) * int(k); j < min(i*int(k), len(c)); j++ {
			acc += c[j]
		}
		total += acc * int32(i)
	}
	return total
}

func min(a, b int) int {
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

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	minimumCost := getMinimumCost(k, c)

	fmt.Fprintf(writer, "%d\n", minimumCost)

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
