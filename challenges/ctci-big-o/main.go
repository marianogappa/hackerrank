package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the primality function below.
//
// Time: O(√n) because loop constraints the incrementor to be smaller than square root of n.
// Space: O(1)
func primality(n int32) string {
	if n == 1 {
		return "Not prime"
	}
	for i := 2; i*i <= int(n); i++ {
		if int(n)%i == 0 {
			return "Not prime"
		}
	}
	return "Prime"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	for pItr := 0; pItr < int(p); pItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		result := primality(n)

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
