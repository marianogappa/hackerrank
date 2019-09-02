package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	ma, mb := mapOf(a), mapOf(b)
	for letterB, countB := range mb {
		minCount := min(ma[letterB], countB)
		if minCount == 0 {
			continue
		}
		ma[letterB] -= minCount
		mb[letterB] -= minCount
	}
	total := 0
	for _, countA := range ma {
		total += countA
	}
	for _, countB := range mb {
		total += countB
	}
	return int32(total)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mapOf(a string) map[byte]int {
	var m = make(map[byte]int, len(a))
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}
	return m
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

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
