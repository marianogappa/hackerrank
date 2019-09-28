package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the crosswordPuzzle function below.
func crosswordPuzzle(crossword []string, words string) []string {
	// Exercise assumes there will always be a successful result
	result, _ := doCrosswordPuzzle(crossword, analyzeCrossword(crossword), strings.Split(words, ";"))
	return result
}

func doCrosswordPuzzle(crossword []string, placeholders map[int][]placeholder, words []string) ([]string, bool) {
	if len(words) == 0 {
		return crossword, true
	}
	word := words[0]
	possiblePlaceholders := placeholders[len(word)] // Exercise assumes this will always exist
	for _, place := range possiblePlaceholders {
		undos, placeOk := tryPlaceWord(crossword, place, word)
		if !placeOk {
			undoPlaceWord(crossword, undos) // Clean crossword
			continue
		}
		result, finishOk := doCrosswordPuzzle(crossword, placeholders, words[1:])
		if !finishOk {
			undoPlaceWord(crossword, undos) // Clean crossword
			continue
		}
		return result, true
	}
	return crossword, false
}

type undo struct {
	y, x int
	c    byte
}

func tryPlaceWord(crossword []string, place placeholder, word string) ([]undo, bool) {
	undos := []undo{}
	var dx, dy int
	if place.isHorizontal {
		dx = 1
	} else {
		dy = 1
	}
	for i := 0; i < place.length; i++ {
		undos = append(undos, undo{place.y + i*dy, place.x + i*dx, crossword[place.y+i*dy][place.x+i*dx]})
		c := crossword[place.y+i*dy][place.x+i*dx]
		if c != '-' && c != word[i] {
			return undos, false
		}
		cw := crossword[place.y+i*dy]
		cw = cw[:place.x+i*dx] + string(word[i]) + cw[place.x+i*dx+1:]
		crossword[place.y+i*dy] = cw
	}
	return undos, true
}

// Time: O(n^2), where n is the number of characters. It's quadratic because each string modification is linear.
// This can be optimised
// Space: O(1)
func undoPlaceWord(crossword []string, undos []undo) {
	for _, undo := range undos {
		cw := crossword[undo.y]
		cw = cw[:undo.x] + string(undo.c) + cw[undo.x+1:]
		crossword[undo.y] = cw
	}
}

type placeholder struct {
	y, x         int
	isHorizontal bool
	length       int
}

// Time: O(n) where n is ∑len(crowssword)
// Space: O(n), because at most we'll have as many placeholders as letters in the crossword, minus mandatory spaces
func analyzeCrossword(crossword []string) map[int][]placeholder {
	result := map[int][]placeholder{}

	isHorizontal := true
	place := placeholder{isHorizontal: isHorizontal}
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			switch crossword[y][x] {
			case '-':
				if place.length == 0 {
					place.x = x
					place.y = y
				}
				place.length++
			default:
				if place.length > 1 {
					result[place.length] = append(result[place.length], place)
				}
				place = placeholder{isHorizontal: isHorizontal, length: 0}
			}
		}
		if place.length > 1 {
			result[place.length] = append(result[place.length], place)
		}
		place = placeholder{isHorizontal: isHorizontal, length: 0}
	}

	isHorizontal = false
	place = placeholder{isHorizontal: isHorizontal}
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			switch crossword[y][x] {
			case '-':
				if place.length == 0 {
					place.x = x
					place.y = y
				}
				place.length++
			default:
				if place.length > 1 {
					result[place.length] = append(result[place.length], place)
				}
				place = placeholder{isHorizontal: isHorizontal, length: 0}
			}
		}
		if place.length > 1 {
			result[place.length] = append(result[place.length], place)
		}
		place = placeholder{isHorizontal: isHorizontal, length: 0}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var crossword []string

	for i := 0; i < 10; i++ {
		crosswordItem := readLine(reader)
		crossword = append(crossword, crosswordItem)
	}

	words := readLine(reader)

	result := crosswordPuzzle(crossword, words)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
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
