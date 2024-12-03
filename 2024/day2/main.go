package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part1()
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

	var result int
	//var report []int

	for sc.Scan() {
		row := sc.Text()
		fmt.Println(row)

		raw := strings.Split(row, " ")
		var report []int

		for _, el := range raw {
			i, _ := strconv.Atoi(el)
			report = append(report, i)
		}

		if rep(report) {
			result++
		}
	}

	fmt.Println("Result:", result)
}

func part2() {
	f, sc := getScanner()
	defer f.Close()

	var result int
	//var report []int

	for sc.Scan() {
		row := sc.Text()
		fmt.Println(row)

		raw := strings.Split(row, " ")
		var report []int

		for _, el := range raw {
			i, _ := strconv.Atoi(el)
			report = append(report, i)
		}

		ok := false

		for i, _ := range report {
			var curReport []int

			if i == 0 {
				curReport = report[1:]
			} else if i == len(report)-1 {
				curReport = report[:i]
			} else {
				curReport = append(curReport, report[:i]...)
				curReport = append(curReport, report[i+1:]...)
			}

			if rep(curReport) {
				ok = true
				break
			}
		}

		if ok {
			result++
		}

		fmt.Println(report, ok)
	}

	fmt.Println("Result:", result)
}

/*
The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.

7 8 6 5 4
*/
func rep(in []int) bool {
	var incr *bool

	t, f := true, false

	if in[0] < in[1] {
		incr = &t
	} else {
		incr = &f
	}

	for i, _ := range in {
		if i == len(in)-1 {
			break
		}

		if in[i] < in[i+1] {
			if in[i+1]-in[i] < 1 || in[i+1]-in[i] > 3 {
				return false
			}
		} else {
			if in[i]-in[i+1] < 1 || in[i]-in[i+1] > 3 {
				return false
			}
		}

		if in[i] > in[i+1] && *incr {
			return false
		}

		if in[i] < in[i+1] && !*incr {
			return false
		}
	}

	return true
}

func rep2(in []int) bool {
	nIncr, nDecr := 0, 0
	errors := 0

	for i, _ := range in {
		if i == len(in)-1 {
			break
		}

		if in[i] < in[i+1] {
			nIncr++
			if in[i+1]-in[i] < 1 || in[i+1]-in[i] > 3 {
				errors++
			}
		} else {
			nDecr++
			if in[i]-in[i+1] < 1 || in[i]-in[i+1] > 3 {
				errors++
			}
		}
	}

	if errors > 1 {
		return false
	}

	if nIncr > 1 && nDecr > 1 {
		return false
	}

	return true
}
