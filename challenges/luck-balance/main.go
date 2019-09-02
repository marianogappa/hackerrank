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

// Complete the luckBalance function below.
func luckBalance(k int32, contests [][]int32) int32 {
	sort.Slice(contests, func(i, j int) bool { return contests[i][0] > contests[j][0] })
	var luck int32
	for _, contest := range contests {
		luckBalance, important := contest[0], contest[1]
		if important == 0 {
			luck += luckBalance
			continue
		}
		if k <= 0 {
			luck -= luckBalance
			continue
		}
		luck += luckBalance
		k--
	}
	return luck
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	var contests [][]int32
	for i := 0; i < int(n); i++ {
		contestsRowTemp := strings.Split(readLine(reader), " ")

		var contestsRow []int32
		for _, contestsRowItem := range contestsRowTemp {
			contestsItemTemp, err := strconv.ParseInt(contestsRowItem, 10, 64)
			checkError(err)
			contestsItem := int32(contestsItemTemp)
			contestsRow = append(contestsRow, contestsItem)
		}

		if len(contestsRow) != 2 {
			panic("Bad input")
		}

		contests = append(contests, contestsRow)
	}

	result := luckBalance(k, contests)

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
