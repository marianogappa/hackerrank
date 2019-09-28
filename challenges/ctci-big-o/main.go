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
// Space: O(n) because there's a slice of size n+1 to keep track of already calculated visited numbers.
//
// This exercise is filled with off-by-one potential errors.
//
// This algorithm can be simplified greatly by just checking if the modulus of every number leading up to √n is == 0,
// but modulus is much more expensive than multiplication, because division cannot be efficiently parallelised and
// requires bounds checking.
// func primality(n int32) string {
// 	if n == 1 {
// 		return "Not prime"
// 	}
// 	if n <= 3 {
// 		return "Prime"
// 	}
// 	visited := make([]bool, n+1)
// 	multiplier := 2
// 	// A multiplier whose square is greater than n will only visit already visited numbers < n. There's no intuitive
// 	// reason for this other than the fact that I learned this by experience while learning Sieve of Eratosthenes.
// 	for multiplier*multiplier <= int(n) {
// 		// This runs Eratosthenes on the current multiplier. Only needs to run from the square of every multiplier,
// 		// because smaller numbers will already be visited by earlier multipliers.
// 		for i := multiplier * multiplier; i <= int(n); i += multiplier {
// 			if i == int(n) { // If we visit n as a multiple of the multiplier, it's not a prime
// 				return "Not prime"
// 			}
// 			visited[i] = true
// 		}

// 		// The current multiplier has been visited already, but possibly hasn't been marked visited
// 		multiplier++

// 		// It's not necessary to run Eratosthenes on numbers that have already been visited
// 		for ; multiplier*multiplier <= int(n) && visited[multiplier]; multiplier++ {
// 		}
// 	}
// 	return "Prime"
// }
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
