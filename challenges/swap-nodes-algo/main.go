package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value       int32
	left, right *node
}

func (n *node) String() string {
	if n == nil {
		return "ø"
	}
	return fmt.Sprintf("%v->(%v,%v)", n.value, n.left.String(), n.right.String())
}

// Time: O(n) with n == len(indexes)
// Space: O(n)
// The values in indexes are sorted in BFS traversal order.
func bfsPush(root *node, indexes [][]int32) {
	if len(indexes) == 0 || root == nil {
		return
	}
	queue := []*node{root}
	for _, curNodeValues := range indexes {
		leftValue, rightValue := curNodeValues[0], curNodeValues[1]
		if leftValue != -1 {
			root.left = &node{leftValue, nil, nil}
			queue = append(queue, root.left)
		}
		if rightValue != -1 {
			root.right = &node{rightValue, nil, nil}
			queue = append(queue, root.right)
		}
		queue = queue[1:]
		if len(queue) == 0 {
			return
		}
		root = queue[0]
	}
}

// Time: O(n) because at worst all nodes are traversed
// Space: O(logn) on a balanced tree, O(n) unbalanced, because the call stack at worst will be the tree height
func swapAtDepth(root *node, depth int32, current int32) {
	if root == nil {
		return
	}
	if current%depth == 0 {
		root.left, root.right = root.right, root.left
	}
	swapAtDepth(root.left, depth, current+1)
	swapAtDepth(root.right, depth, current+1)
}

// Time: O(n) because at worst all nodes are traversed
// Space: O(logn) on a balanced tree, O(n) unbalanced, because the call stack at worst will be the tree height
func inOrderTraversal(root *node) []int32 {
	if root == nil {
		return []int32{}
	}
	result := []int32{}
	result = append(result, inOrderTraversal(root.left)...)
	result = append(result, root.value)
	result = append(result, inOrderTraversal(root.right)...)
	return result
}

/*
 * Complete the swapNodes function below.
 */
// Time: O(n)
// Space: O(n)
func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	/*
	 * Write your code here.
	 */
	result := [][]int32{}
	root := &node{1, nil, nil}
	bfsPush(root, indexes)
	for _, swapDepth := range queries {
		swapAtDepth(root, swapDepth, 1)
		result = append(result, inOrderTraversal(root))
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var indexes [][]int32
	for indexesRowItr := 0; indexesRowItr < int(n); indexesRowItr++ {
		indexesRowTemp := strings.Split(readLine(reader), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != int(2) {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int32

	for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for resultRowItr, rowItem := range result {
		for resultColumnItr, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if resultColumnItr != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if resultRowItr != len(result)-1 {
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
