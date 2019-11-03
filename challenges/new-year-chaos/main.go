package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
// Time: O(n)
// Space: O(1)
// I couldn't get it; had to look up the solution. Look at the "I DON'T GET IT" section when I figured it out.
func minimumBribes(q []int32) {
	var bribes int32
	for i, num := range q {
		pos := int32(i + 1)
		// If the number is more than too positions to the left of its number, they bribed more than two times.
		if num > pos+2 {
			fmt.Println("Too chaotic")
			return
		}
		// For every other number, we check if the two previous ones are greater. For every greater, there must
		// have been a bribe. We don't check further than 2 places, because the previous condition ensures this
		// doesn't happen.
		//
		// I DON'T GET IT! IS THIS DOUBLE-COUNTING MAYBE?!?!
		// No. Imagine the police going person by person, asking who bribed them. Each person who has n
		// people ahead of them must have been bribed n times, so forget about the swapping; the swapping idea
		// is very confusing. Just go through every person, and check who is ahead of them with a larger number.
		// The catch is that to keep it linear you must not go through the full array every time. But you don't
		// need to; just go up to two places in front of where the current number "should have been", because
		// that's how far ahead a number could have bribed them.
		for j := max(0, int(q[i])-2); j < i; j++ {
			if q[j] > q[i] {
				bribes++
			}
		}
	}
	fmt.Println(bribes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
