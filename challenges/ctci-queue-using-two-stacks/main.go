package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	q []int
}

func (q *queue) pop() {
	q.q = q.q[1:]
}

func (q *queue) push(n int) {
	q.q = append(q.q, n)
}

func (q *queue) peek() {
	if len(q.q) == 0 {
		log.Fatal("You promised this would never happen!")
	}
	fmt.Println(q.q[0])
}

// Time: O(n)
// Space: O(n)
// I think the point of this exercise is missed in this Go implementation, because
// you can efficiently queue and dequeue using slices in Go, so you don't need two stacks.
func twoStacks(queries [][]int) {
	q := &queue{}
	for _, query := range queries {
		switch query[0] {
		case 1:
			q.push(query[1])
		case 2:
			q.pop()
		case 3:
			q.peek()
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	queries := [][]int{}
	for qItr := 0; qItr < int(q); qItr++ {
		intStrings := strings.Split(readLine(reader), " ")
		query := []int{}
		for _, intString := range intStrings {
			n, err := strconv.ParseInt(intString, 10, 64)
			checkError(err)
			query = append(query, int(n))
		}
		queries = append(queries, query)
	}
	twoStacks(queries)

	//Enter your code here. Read input from STDIN. Print output to STDOUT
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
