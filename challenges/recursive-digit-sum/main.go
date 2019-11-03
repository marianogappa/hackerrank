package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the superDigit function below.
//
// Time: O(n)
// Space: O(1)
//
// The k component of this problem can be simplified by just
// multiplying each digit by k.
//
// On the first iteration we have a string, so we convert
// each digit and get the sum of digits (multiplied by k).
//
// Then, doSuperDigit sums the digits until the result is
// smaller than 10, and then returns it.
func superDigit(n string, k int32) int32 {
	var total int
	for i := 0; i < len(n); i++ {
		total += int(n[i]-'0') * int(k)
	}
	return int32(doSuperDigit(total))
}

// Time: O(logn) Log in base 10
// Space: O(1)
func doSuperDigit(n int) int {
	if n < 10 {
		return n
	}
	return doSuperDigit(sumDigits(n))
}

// Time: O(logn) Log in base 10
// Space: O(1)
func sumDigits(n int) int {
	var mod, acc int
	for n > 10 {
		n, mod = n/10, n%10
		acc += mod
	}
	return acc + n
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	n := nk[0]

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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
