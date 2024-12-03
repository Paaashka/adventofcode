package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//part1()
	//part2()
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
		fmt.Println(row)
	}

	fmt.Println("Result:", result)
}

func part2() {
	f, sc := getScanner()
	defer f.Close()

	var result int

	for sc.Scan() {
		row := sc.Text()
		fmt.Println(row)
	}

	fmt.Println("Result:", result)
}
