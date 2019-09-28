package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the stepPerms function below.
//
// Space: O(1)
// Time: would be O(3^n) because the maximum branching is 3 and the max height of the tree is n, if we always choose 1,
// but due to the memoization step it becomes O(n), because branches of already visited steps won't do recursion.
func stepPerms(n int32) int32 {
	return doStepPerms(n, map[int32]int32{0: 0, 1: 1, 2: 2, 3: 4})
}

func doStepPerms(n int32, cache map[int32]int32) int32 {
	if v, ok := cache[n]; ok {
		return v
	}
	cache[n] = doStepPerms(n-1, cache) + doStepPerms(n-2, cache) + doStepPerms(n-3, cache)
	return cache[n]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	s := int32(sTemp)

	for sItr := 0; sItr < int(s); sItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		res := stepPerms(n)

		fmt.Fprintf(writer, "%d\n", res)
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
