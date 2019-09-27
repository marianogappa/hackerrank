package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the flippingBits function below.
// Space: O(1) because the bits are flipped in-place
// Time: O(n^2) if n is the number of bits to flip, but irrelevant as n is a small number
func flippingBits(n int64) int64 {
	for i := 0; i < 32; i++ {
		if (n>>uint(i))&1 == 1 {
			n &= allOnesExceptOn(i)
		} else {
			n |= (int64(1) << uint(i))
		}
	}
	return n
}

// func show(n int64) {
// 	for i := 32; i >= 0; i-- {
// 		if (n>>uint(i))&1 == 1 {
// 			fmt.Print("1")
// 		} else {
// 			fmt.Print("0")
// 		}
// 	}
// 	fmt.Println("")
// }

func allOnesExceptOn(j int) int64 {
	var num int64
	for i := 0; i < 32; i++ {
		if j == i {
			continue
		}
		num |= (int64(1) << uint(i))
	}
	return num
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		n, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		result := flippingBits(n)

		fmt.Fprintf(writer, "%d\n", result)
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
