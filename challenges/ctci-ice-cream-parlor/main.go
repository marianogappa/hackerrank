package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the whatFlavors function below.
//
// Space: O(n)
// Time: O(n)
func whatFlavors(cost []int32, money int32) {
	// Space: O(n)
	// Time: O(n)
	costToIndices := map[int32][]int32{}
	for i, c := range cost {
		costToIndices[c] = append(costToIndices[c], int32(i+1))
	}

	// Space: O(1)
	// Time: O(n)
	for _, c := range cost {
		if c >= money {
			continue
		}
		if _, ok := costToIndices[money-c]; !ok {
			continue
		}
		if money-c == c && len(costToIndices[c]) <= 1 {
			continue
		}
		if money-c == c {
			printSortedAscending(costToIndices[c][0], costToIndices[c][1])
			return
		}
		// Solution guaranteed to exist
		printSortedAscending(costToIndices[c][0], costToIndices[money-c][0])
		return
	}
}

func printSortedAscending(a, b int32) {
	fmt.Printf("%v %v\n", min(a, b), max(a, b))
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
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
		moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
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
