package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the largestRectangle function below.
// Time: average(nlogn) because we divide by half logn times, each time going through n/2 numbers
// Time: O(n^2) because we divide n times, each time going through k-1 numbers
// Space: O(n) worst-case call stack height
//
// The insight is that the total is either:
// 1) the (minimum height)*(length of slice)
// 2) solving the problem to the left of the minimum
// 3) solving the problem to the right of the minimum
// Because if the solution includes the minimum, then it's not higher than (1).
// Thus, it becomes a divide-and-conquer problem, which on average (if heights are random) should
// solve in n*logn time (assuming every iteration halves the slice), using only a call stack of logn space.
// In the worst case (a monotonically increasing height), every iteration decreases length by 1,
// so there's n iterations.
func largestRectangle(h []int32) int64 {
	if len(h) == 0 {
		return 0
	}
	if len(h) == 1 {
		return int64(h[0])
	}
	minimum, index := minInSlice(h)
	fmt.Println(minimum, index, int64(len(h))*int64(minimum))
	return max(int64(len(h))*int64(minimum), largestRectangle(h[:index]), largestRectangle(h[index+1:]))
}

// Time: O(n)
// Space: O(1)
func minInSlice(h []int32) (int32, int) {
	var minimum int32 = math.MaxInt32
	var index int
	for i, num := range h {
		if num < minimum {
			minimum = num
			index = i
		}
	}
	return minimum, index
}

// Time: O(1)
// Space: O(1)
func max(a, b, c int64) int64 {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
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

	hTemp := strings.Split(readLine(reader), " ")

	var h []int32

	for i := 0; i < int(n); i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	result := largestRectangle(h)

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
