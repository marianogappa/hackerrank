package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the triplets function below.
// Time: O(nlogn)
// Space: O(n)
func triplets(a []int32, b []int32, c []int32) int64 {
	// Time: O(nlogn)
	// Space: O(1)
	sortInt64Slice(a)
	sortInt64Slice(b)
	sortInt64Slice(c)

	// Time: O(n)
	// Space: O(n)
	a = uniq(a)
	b = uniq(b)
	c = uniq(c)

	// Time: O(n) where n is the size of the largest length of the slices
	// Space: O(1)
	var ia, ic int
	var total int64
	for _, bn := range b {
		for ; ia < len(a) && a[ia] <= bn; ia++ {
		}
		for ; ic < len(c) && c[ic] <= bn; ic++ {
		}
		total += int64(ia * ic)
	}
	return total
}

func sortInt64Slice(is []int32) {
	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })
}

func uniq(is []int32) []int32 {
	u := []int32{}
	var last int32
	for i, n := range is {
		if i == 0 || n != last {
			u = append(u, n)
			last = n
		}
	}
	return u
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	lenaLenbLenc := strings.Split(readLine(reader), " ")

	lenaTemp, err := strconv.ParseInt(lenaLenbLenc[0], 10, 64)
	checkError(err)
	lena := int32(lenaTemp)

	lenbTemp, err := strconv.ParseInt(lenaLenbLenc[1], 10, 64)
	checkError(err)
	lenb := int32(lenbTemp)

	lencTemp, err := strconv.ParseInt(lenaLenbLenc[2], 10, 64)
	checkError(err)
	lenc := int32(lencTemp)

	arraTemp := strings.Split(readLine(reader), " ")

	var arra []int32

	for i := 0; i < int(lena); i++ {
		arraItemTemp, err := strconv.ParseInt(arraTemp[i], 10, 64)
		checkError(err)
		arraItem := int32(arraItemTemp)
		arra = append(arra, arraItem)
	}

	arrbTemp := strings.Split(readLine(reader), " ")

	var arrb []int32

	for i := 0; i < int(lenb); i++ {
		arrbItemTemp, err := strconv.ParseInt(arrbTemp[i], 10, 64)
		checkError(err)
		arrbItem := int32(arrbItemTemp)
		arrb = append(arrb, arrbItem)
	}

	arrcTemp := strings.Split(readLine(reader), " ")

	var arrc []int32

	for i := 0; i < int(lenc); i++ {
		arrcItemTemp, err := strconv.ParseInt(arrcTemp[i], 10, 64)
		checkError(err)
		arrcItem := int32(arrcItemTemp)
		arrc = append(arrc, arrcItem)
	}

	ans := triplets(arra, arrb, arrc)

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
