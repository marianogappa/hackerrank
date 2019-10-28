package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the substrCount function below.
type segment struct {
	b     byte
	count int
}

// Time: O(n)
// Space: O(n)
func substrCount(n int32, s string) int64 {
	var total int64
	// Split the string into segments of equal letters, with counts
	// Time: O(n)
	// Space: O(n)
	segments := []segment{}
	for i := 0; i < len(s); i++ {
		if len(segments) == 0 || s[i] != segments[len(segments)-1].b {
			segments = append(segments, segment{s[i], 1})
			continue
		}
		segments[len(segments)-1].count++
	}

	// Calculate how many "special strings" can be found on each segment
	// Time: O(n)
	// Space: O(1)
	for _, segment := range segments {
		// len(ss("a")) = len([]string{"a"}) = 1
		// len(ss("aa")) = len([]string{"a", "a", "aa"}) = 3
		// len(ss("aaa")) = len([]string{"a", "a", "a", "aa", "aa", "aaa"}) = 6
		// len(ss("aaaa")) = len([]string{"a", "a", "a", "a", "aa", "aa", "aa", "aaa", "aaa", "aaaa"}) = 10
		total += int64(segment.count) * int64(segment.count+1) / 2
	}

	// With a "cursor" in the middle of any 3 segments, calculate how many
	// "special" strings can be found by this pattern (a+)b(a+)
	// Time: O(n)
	// Space: O(1)
	for i := 1; i < len(segments)-1; i++ {
		if segments[i].count != 1 || segments[i-1].b != segments[i+1].b {
			continue
		}
		total += min(int64(segments[i-1].count), int64(segments[i+1].count))
	}

	return total
}

func min(a, b int64) int64 {
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

	s := readLine(reader)

	result := substrCount(n, s)

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
