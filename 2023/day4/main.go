package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winning []int
	numbers []int
	score   int
}

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

	for sc.Scan() {
		row := sc.Text()
		parts := strings.Split(strings.Split(row, ": ")[1], "|")

		winning := split(parts[0][:len(parts[0])-1])
		numbers := split(parts[1][1:])

		winningMap := make(map[int]struct{})

		for i := 0; i < len(winning); i++ {
			winningMap[winning[i]] = struct{}{}
		}

		wins := getWins(winning, numbers)
		if wins > 0 {
			result += int(math.Pow(2, float64(wins-1)))
		}
	}

	fmt.Println("Result:", result)
}

func part2() {
	f, sc := getScanner()
	defer f.Close()

	var cards []Card

	for sc.Scan() {
		row := sc.Text()
		parts := strings.Split(strings.Split(row, ": ")[1], "|")

		winning := split(parts[0][:len(parts[0])-1])
		numbers := split(parts[1][1:])

		cards = append(cards, Card{
			winning: winning,
			numbers: numbers,
			score:   getWins(winning, numbers),
		})
	}

	result := 0
	for i := 0; i < len(cards); i++ {
		result += calcRecursive(i, cards)
	}

	fmt.Println("Result:", result)
}

func split(in string) (res []int) {
	var buf string
	for i := 0; i < len(in); i++ {
		if i%3 == 2 {
			continue
		}

		s := string(in[i])
		buf += s

		if i%3 == 1 || i == len(in)-1 {
			n, _ := strconv.Atoi(strings.TrimSpace(buf))
			res = append(res, n)

			buf = ""
		}
	}

	return
}

func getWins(winning, numbers []int) int {
	winningMap := make(map[int]struct{})
	for i := 0; i < len(winning); i++ {
		winningMap[winning[i]] = struct{}{}
	}

	res := 0
	for i := 0; i < len(numbers); i++ {
		if _, ok := winningMap[numbers[i]]; ok {
			res++
		}
	}

	return res
}

func calcRecursive(idx int, cs []Card) int {
	r := 1

	for i := 1; i <= cs[idx].score; i++ {
		newIdx := idx + i
		r += calcRecursive(newIdx, cs)
	}

	return r
}
