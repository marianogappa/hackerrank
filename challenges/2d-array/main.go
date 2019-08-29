package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hourglassSum function below.
// Space: O(1)
// Time complexity doesn't make sense for fixed array sizes
// But it would be strideX * strideY * maskX * maskY (or len of each)
// Sliding hourglasses share values, so this could be optimised somewhat.
func hourglassSum(arr [][]int32) int32 {
	hourglassMask := [][]int32{
		{1, 1, 1},
		{0, 1, 0},
		{1, 1, 1},
	}
	maxHourglassSum := int32(-1 << 31)
	for strideY := 0; strideY <= 3; strideY++ {
		for strideX := 0; strideX <= 3; strideX++ {
			curHourglassSum := int32(0)
			for maskY := 0; maskY <= 2; maskY++ {
				for maskX := 0; maskX <= 2; maskX++ {
					curHourglassSum += arr[strideY+maskY][strideX+maskX] * hourglassMask[maskY][maskX]
				}
			}
			if curHourglassSum > maxHourglassSum {
				maxHourglassSum = curHourglassSum
			}
		}
	}
	return maxHourglassSum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
