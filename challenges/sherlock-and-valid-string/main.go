package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the isValid function below.
// Time: O(n)
// Space: O(n)
func isValid(s string) string {
	// For each letter, how many times it appears (i.e. its frequency)
	letterFreq := map[byte]int{}
	for i := 0; i < len(s); i++ {
		letterFreq[s[i]]++
	}
	// Count all the different frequencies (at most there can be two different freqs)
	freqCounts := map[int]int{}
	for _, freq := range letterFreq {
		freqCounts[freq]++
	}
	fmt.Printf("%#+v\n", freqCounts)
	// If there's only 1 distinct frequency, then s is a valid string
	if len(freqCounts) == 1 {
		return "YES"
	}
	// If there's more than 2 distinct frequencies, then s is not a valid string
	if len(freqCounts) >= 3 {
		return "NO"
	}

	// This horrible section is meant to extract the two frequencies and their counts
	isFirst := true
	var freqA, freqB, countA, countB int
	for freq, count := range freqCounts {
		if isFirst {
			isFirst = false
			freqA, countA = freq, count
		} else {
			freqB, countB = freq, count
		}
	}

	// If one of the two frequencies is just one letter once, we can just remove it and make it valid
	if (countA == 1 && freqA == 1) || (countB == 1 && freqB == 1) {
		return "YES"
	}

	// If one of the freqs is represented by just one letter, and if that freq is exactly one more than the
	// other freq, then by removing a letter from the former freq, all letters will have the same freq
	if (countA == 1 && freqA-freqB == 1) || (countB == 1 && freqB-freqA == 1) {
		return "YES"
	}

	// Otherwise its not valid
	return "NO"
}

func main() {
	fmt.Println(isValid("ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd"))
	os.Exit(0)
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

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
