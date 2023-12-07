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
	part1()
	part2()
}

func getScanner() (*os.File, *bufio.Scanner) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	return file, bufio.NewScanner(file)
}

func part1() {

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

		if i, err := strconv.ParseInt(rowString, 10, 32); err == nil {
			result = result + i
		}
	}

	fmt.Println("Result: ", result)
}

func part2() {

	f, sc := getScanner()
	defer f.Close()

	// optionally, resize scanner's capacity for lines over 64K, see next example
	var result int64
	for sc.Scan() {
		n := fetchNumberFromString(sc.Text())

		if n == 0 {
			panic("null!")
		}

		result += n
	}

	fmt.Println("Result: ", result)
}

func fetchNumberFromString(str string) int64 {
	firstDigit := getFirstDigit(str, replacements)
	secondDigit := getFirstDigit(reverse(str), revReplacements)

	return firstDigit*10 + secondDigit
}

func getFirstDigit(str string, words map[string]string) int64 {
	var found string

	for i := 0; i < len(str); i++ {
		symbol := string(str[i])

		if isInt(symbol) {
			found = symbol
			break
		}

		for old, n := range words {
			if strings.Index(str[i:], old) == 0 {
				found = n
				break
			}
		}

		if found != "" {
			break
		}
	}

	if i, err := strconv.ParseInt(found, 10, 32); err == nil {
		return i
	}

	return 0
}

// taken from SO
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

func isInt(in string) bool {
	_, ok := numbers[in]

	return ok
}
