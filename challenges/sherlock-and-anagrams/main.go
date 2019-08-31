package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
// What's the Time/Space complexity of this? With n == len(s)
// Space: at worst we have n maps at a time, so Space: O(n) ?
// Time: (n problems) * (n substrings) * ((n iterations for mapFor) + ((n map checks) * (n iterations to check map))
// Looks like O(n^4) but the worst cases of each n are never simultaneous, so it seems cheaper.
// Also wondering if the innermost for-loop could be removed somehow.
func sherlockAndAnagrams(s string) int32 {
	if len(s) < 2 {
		return 0
	}
	var (
		minLength = 1
		maxLength = len(s) - 1 // Len(s) can't have an associated pair
		total     = 0
	)
	// Can be simplified into independent subproblems by length
	for i := minLength; i <= maxLength; i++ {
		total += sherlockAndAnagramsForLength(s, i)
	}
	return int32(total)
}

func sherlockAndAnagramsForLength(s string, length int) int {
	maps := []map[byte]int{}
	duplicates := 0 // Anagram pairs happen when a "map" of a string has already been seen (i.e. a duplicate)
	// For each substring of len _length_ in _s_, get its "map" and append it to _maps_ unless already seen
	for i := 0; i < len(s)-length+1; i++ {
		curMap := mapFor(s[i : i+length])
		// Note: if we could make a map of maps we wouldn't need a loop here, but keys of maps cannot be maps
		for _, m := range maps {
			if reflect.DeepEqual(curMap, m) { // If already seen, don't append and increment the list of duplicates
				duplicates++
				continue
			}
		}
		maps = append(maps, curMap)
	}
	return duplicates
}

// A "map" of a string is a set of tuples: (character, appearances in string)
// Two strings are anagram pairs if their "maps" are equal
func mapFor(s string) map[byte]int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	return m
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
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

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
