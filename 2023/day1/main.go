package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers = map[string]struct{}{
	"0": {},
	"1": {},
	"2": {},
	"3": {},
	"4": {},
	"5": {},
	"6": {},
	"7": {},
	"8": {},
	"9": {},
}

var words = [][2]string{
	{"one", "1"},
	{"two", "2"},
	{"six", "6"},
	{"four", "4"},
	{"five", "5"},
	{"three", "3"},
	{"nine", "9"},
	{"seven", "7"},
	{"eight", "8"},
}

var replacements = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var revReplacements = map[string]string{
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

func main() {
	fmt.Println("hello world")

	//task1()
	task2()
}

func getScanner() (*os.File, *bufio.Scanner) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	return file, bufio.NewScanner(file)
}

func task1() {

	f, sc := getScanner()
	defer f.Close()

	// optionally, resize scanner's capacity for lines over 64K, see next example
	var row string
	var result int64
	for sc.Scan() {
		row = sc.Text()
		rowLen := len(row)
		var rowString string

		for i := 0; i < rowLen; i++ {
			symbol := string(row[i])

			if _, ok := numbers[symbol]; ok {
				rowString = symbol
				break
			}
		}

		for i := rowLen - 1; i >= 0; i-- {
			symbol := string(row[i])

			if _, ok := numbers[symbol]; ok {
				rowString += symbol
				break
			}
		}

		fmt.Println(rowString)

		if i, err := strconv.ParseInt(rowString, 10, 32); err == nil {
			result = result + i
		}
	}

	fmt.Println("Result: ", result)
}

func task2() {
	//
	//words := map[int64][]string{
	//	1: {"o", "n", "e"},
	//	2: {"t", "w", "o"},
	//	3: {"t", "h", "r", "e", "e"},
	//	4: {"f", "o", "u", "r"},
	//	5: {"f", "i", "v", "e"},
	//	6: {"s", "i", "x"},
	//	7: {"s", "e", "v", "e", "n"},
	//	8: {"e", "i", "g", "h", "t"},
	//	9: {"n", "i", "n", "e"},
	//}

	f, sc := getScanner()
	defer f.Close()

	// optionally, resize scanner's capacity for lines over 64K, see next example
	var result int64
	for sc.Scan() {
		result += fetchNumberFromString(sc.Text())
	}

	fmt.Println("Result: ", result)
}

func fetchNumberFromString(str string) int64 {
	oldStr := str

	str = ReplaceWords(str)

	/*	for _, pair := range words {
		row = strings.ReplaceAll(row, pair[0], pair[1])
	}*/

	rowLen := len(str)
	var rowString string

	for i := 0; i < rowLen; i++ {
		symbol := string(str[i])

		if _, ok := numbers[symbol]; ok {
			rowString = symbol
			break
		}
	}

	for i := rowLen - 1; i >= 0; i-- {
		symbol := string(str[i])

		if _, ok := numbers[symbol]; ok {
			rowString += symbol
			break
		}
	}

	if i, err := strconv.ParseInt(rowString, 10, 32); err == nil {
		fmt.Println(oldStr, "->", str, " = ", rowString, i)
		return i
	}

	return 0
}

func replaceFirstWord(input string, repls map[string]string) string {
	foundIdx := len(input)
	var foundWord string

	for old := range repls {
		i := strings.Index(input, old)
		if i != -1 && i < foundIdx {
			foundIdx = i
			foundWord = old
		}
	}

	if foundWord != "" {
		input = strings.Replace(input, foundWord, repls[foundWord], 1)
	}

	return input
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func ReplaceWords(input string) string {
	result := input

	result = replaceFirstWord(result, replacements)

	result = reverse(replaceFirstWord(reverse(result), revReplacements))

	return result
}
