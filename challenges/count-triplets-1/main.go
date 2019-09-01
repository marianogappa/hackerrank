package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
// Space: O(n)
// Time: O(n)
// I failed twice to solve this on my own and looked at a solution online.
// The idea is: for each number in _arr_, calculate _num*r_, store it in _m1_ and wait for that to appear.
// If it shows up, calculate the num*r, store it in _m2_ and wait for that to appear.
// If it shows up, increment the total count.
// Numbers may appear more than once, so keep count of appearances. The way in which the summing of appearances is done
// handles the combinatorics; that's very unintuitive to me.
func countTriplets(arr []int64, r int64) int64 {
	m1 := make(map[int64]int64, len(arr))
	m2 := make(map[int64]int64, len(arr))

	var count int64
	for _, num := range arr {
		count += m2[num]
		if cnt, ok := m1[num]; ok {
			m2[num*r] += cnt
		}
		m1[num*r]++
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
