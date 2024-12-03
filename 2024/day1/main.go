package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var a, b int

	var left, right []int

	for sc.Scan() {
		row := sc.Text()

		a, _ = strconv.Atoi(strings.Split(row, "   ")[0])
		b, _ = strconv.Atoi(strings.Split(row, "   ")[1])

		left = append(left, a)
		right = append(right, b)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i, l := range left {
		if l < right[i] {
			result += right[i] - l
		} else {
			result += l - right[i]
		}
	}

	fmt.Println("Result:", result)
}

func part2() {
	f, sc := getScanner()
	defer f.Close()

	var result int
	var a, b int

	var left []int
	right := make(map[int]int)

	for sc.Scan() {
		row := sc.Text()

		a, _ = strconv.Atoi(strings.Split(row, "   ")[0])
		left = append(left, a)

		b, _ = strconv.Atoi(strings.Split(row, "   ")[1])
		if _, ok := right[b]; !ok {
			right[b] = 0
		}

		right[b]++
	}

	for _, l := range left {
		result += l * right[l]
	}

	fmt.Println("Result:", result)
}
