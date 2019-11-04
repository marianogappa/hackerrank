package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the maxCircle function below.
//
// Time: O(n) where n is the length of distinct friends
// Space: O(n)
//
// I started with a solution that had to traverse a map[int32]int32 based graph to figure out
// which one was the largest circle on every iteration, and got timed out.
//
// My next solution creates "circles" (*map[int32]struct{}), and maps friends to circles.
//
// O(1) <= Add two friends that we haven't seen before
// O(1) <= Add a new friend to a friend on an existing circle
// O(1) <= Acknowledge the friendship of two friends that already exist
// O(1) <= Calculate the maximum group size after any query
// O(n) <= Merge the circles, when two friends from different circles become friends. Linear to the smallest circle.
func maxCircle(queries [][]int32) []int32 {
	var (
		results      = []int32{}
		friends      = map[int32]*map[int32]struct{}{}
		maxGroupSize = 0
	)
	for _, query := range queries {
		var (
			friendA, friendB       = query[0], query[1]
			circleA, circleAExists = friends[friendA]
			circleB, circleBExists = friends[friendB]
		)
		switch {
		case !circleAExists && !circleBExists:
			// None exists; create new circle with them.
			newCircle := map[int32]struct{}{friendA: struct{}{}, friendB: struct{}{}}
			friends[friendA] = &newCircle
			friends[friendB] = &newCircle
			maxGroupSize = max(maxGroupSize, 2)
		case circleAExists && !circleBExists:
			// A exists: add B to A's circle.
			(*circleA)[friendB] = struct{}{}
			friends[friendB] = circleA
			maxGroupSize = max(maxGroupSize, len(*circleA))
		case !circleAExists && circleBExists:
			// B exists: add A to B's circle.
			(*circleB)[friendA] = struct{}{}
			friends[friendA] = circleB
			maxGroupSize = max(maxGroupSize, len(*circleB))
		case circleAExists && circleBExists:
			// Both already exist. If they belong to the same circle, nothing to do. Otherwise, merge circles!
			if circleA != circleB {
				larger, smaller := orderByLengthDesc(circleA, circleB)
				for fs := range *smaller {
					friends[fs] = larger
					(*larger)[fs] = struct{}{}
					delete(*smaller, fs)
				}
				smaller = nil
				maxGroupSize = max(maxGroupSize, len(*larger))
			}
		}
		results = append(results, int32(maxGroupSize))
	}
	return results
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func orderByLengthDesc(a, b *map[int32]struct{}) (*map[int32]struct{}, *map[int32]struct{}) {
	if len(*a) > len(*b) {
		return a, b
	}
	return b, a
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

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := maxCircle(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
