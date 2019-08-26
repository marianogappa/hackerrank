package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countingValleys function below.
// Time: O(n)
// Space: O(1)
func countingValleys(n int32, s string) int32 {
	valleyCount := 0
	altitude := 0
	inValley := false
	for _, c := range s {
		switch c {
		case 'U':
			altitude++
			if inValley && altitude == 0 {
				inValley = false
			}
		case 'D':
			altitude--
			if !inValley && altitude < 0 {
				inValley = true
				valleyCount++
			}
		default:
		}
	}
	return int32(valleyCount)
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

	result := countingValleys(n, s)

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
