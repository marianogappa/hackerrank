package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the isBalanced function below.
// Time: O(n) with n == len(s), because we go through the string once
// Space: O(n), because in the worst case all characters are opening braces, so we store them all
func isBalanced(s string) string {
	var (
		queue   = make([]byte, 0)
		mapping = map[byte]byte{')': '(', ']': '[', '}': '{'}
	)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			queue = append(queue, s[i])
		case ')', ']', '}':
			if len(queue) == 0 {
				return "NO"
			}
			if queue[len(queue)-1] != mapping[s[i]] {
				return "NO"
			}
			queue = queue[:len(queue)-1]
		}
	}
	if len(queue) != 0 {
		return "NO"
	}
	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
