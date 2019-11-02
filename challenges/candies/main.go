package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the candies function below.
// Time: O(n)
// Space: O(n)
// I couldn't figure it out after thinking for hours. I looked up this solution.
// I see it works but I don't find it very intuitive.
func candies(n int32, arr []int32) int64 {
	if len(arr) <= 1 {
		return int64(len(arr))
	}

	var candies = make([]int32, len(arr))
	// Start with minimal candy
	candies[0] = 1
	// Traverse left to right, from the second
	for i := 1; i < len(arr); i++ {
		// If left neighbour has lower score, gimme one more candy than them
		if arr[i] > arr[i-1] {
			candies[i] = candies[i-1] + 1
		} else { // Otherwise, give me minimal candy
			candies[i] = 1
		}
	}

	// Traverse right to left, from the one before last
	for i := len(arr) - 2; i >= 0; i-- {
		// If right neighbour has lower score and not less candy, gimme 1 more candy than them
		if arr[i] > arr[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	// Sum the allocated candy
	var sum int64
	for _, n := range candies {
		sum += int64(n)
	}
	return sum
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

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := candies(n, arr)

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
