package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Row []string
type Matrix []Row
type XY [2]int

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

func getMatrix() Matrix {
	f, sc := getScanner()
	defer f.Close()

	m := Matrix{}
	for sc.Scan() {
		row := sc.Text()
		r := Row{}

		for i := 0; i < len(row); i++ {
			r = append(r, string(row[i]))
		}
		m = append(m, r)
	}

	return m
}

func part1() {
	matrix := getMatrix()

	var result int
	for x := 0; x < len(matrix); x++ {

		var buffer int
		hasSymbols := false
		for y := 0; y < len(matrix[x]); y++ {
			n, err := strconv.Atoi(matrix[x][y])
			if err == nil {
				// is number
				buffer = buffer*10 + n
				if hasSymbols == false && hasNeighborSymbol(matrix, x, y) {
					hasSymbols = true
				}

				if y != len(matrix[x])-1 {
					continue
				}
			}

			// number has ended
			if buffer != 0 {
				// process buffer
				//fmt.Println("err", err, "buffer", buffer, "hasSymbols", hasSymbols)
				if hasSymbols {
					result += buffer
				}

				buffer = 0
				hasSymbols = false
			}
		}
	}

	fmt.Println("Result:", result)
}

func part2() {
	matrix := getMatrix()

	var result int

	loneStars := map[XY]int{}

	for x := 0; x < len(matrix); x++ {

		var buffer int
		currentNumberHasStar := false
		var numberStars []XY
		for y := 0; y < len(matrix[x]); y++ {
			n, err := strconv.Atoi(matrix[x][y])
			if err == nil {
				// is number
				buffer = buffer*10 + n

				sc := getNeighborStars(matrix, x, y)
				if len(sc) > 0 {
					numberStars = mergeXYs(numberStars, sc)
					currentNumberHasStar = true
				}

				if y != len(matrix[x])-1 {
					continue
				}
			}

			// number has ended
			if buffer != 0 {
				// process buffer
				fmt.Println("buffer", buffer, "currentNumberHasStar", currentNumberHasStar)

				for i := 0; i < len(numberStars); i++ {
					pair, ok := loneStars[numberStars[i]]

					if ok {
						fmt.Println("Found lone star", numberStars[i], "pair", pair)
						result += pair * buffer
						delete(loneStars, numberStars[i])
					} else {
						loneStars[numberStars[i]] = buffer
					}
				}

				buffer = 0
				currentNumberHasStar = false
				numberStars = []XY{}
			}
		}
	}

	fmt.Println("Result:", result)
}

func getNeighbors(x, y int) []XY {
	var ns []XY

	ns = append(
		ns,
		[2]int{x - 1, y},     // n
		[2]int{x + 1, y},     // s
		[2]int{x, y - 1},     // w
		[2]int{x, y + 1},     // e
		[2]int{x - 1, y + 1}, // ne
		[2]int{x + 1, y + 1}, // se
		[2]int{x + 1, y - 1}, // sw
		[2]int{x - 1, y - 1}, // nw
	)

	return ns
}

func hasNeighborSymbol(matrix Matrix, x, y int) bool {
	ns := getNeighbors(x, y)

	for i := 0; i < len(ns); i++ {
		if isSymbol(getValue(matrix, ns[i][0], ns[i][1])) {
			return true
		}
	}

	return false
}

func getNeighborStars(matrix Matrix, x, y int) []XY {
	ns := getNeighbors(x, y)

	var res []XY
	for i := 0; i < len(ns); i++ {
		if getValue(matrix, ns[i][0], ns[i][1]) == "*" {
			res = append(res, XY{ns[i][0], ns[i][1]})
		}
	}

	return res
}

func getValue(matrix Matrix, x, y int) string {
	maxX := len(matrix) - 1 // can be optimized
	maxY := len(matrix[0]) - 1

	if x < 0 || x > maxX || y < 0 || y > maxY {
		return ""
	}

	return matrix[x][y]
}

func isSymbol(s string) bool {
	if s == "" || s == "." {
		return false
	}

	if _, err := strconv.Atoi(s); err == nil {
		return false
	}

	return true
}

// not optimal but okay version
func mergeXYs(target []XY, source []XY) []XY {
	for i := 0; i < len(source); i++ {
		sx, sy := source[i][0], source[i][1]

		exists := false
		for j := 0; j < len(target); j++ {
			tx, ty := target[j][0], target[j][1]

			if sx == tx && sy == ty {
				exists = true
			}
		}

		if !exists {
			target = append(target, XY{sx, sy})
		}
	}

	return target
}
