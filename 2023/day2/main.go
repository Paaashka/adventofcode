package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID   int
	Sets []GameSet
}

type GameSet struct {
	Blue  int
	Red   int
	Green int
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

	res := 0
	for sc.Scan() {
		g := parseRow(sc.Text())

		isCorrect := true
		for i := 0; i < len(g.Sets); i++ {
			if g.Sets[i].Blue > 14 || g.Sets[i].Green > 13 || g.Sets[i].Red > 12 {
				isCorrect = false
			}
		}

		if isCorrect {
			res += g.ID
		}
	}

	fmt.Println("Result:", res)
}

func part2() {
	f, sc := getScanner()
	defer f.Close()

	res := 0
	for sc.Scan() {
		game := parseRow(sc.Text())

		var r, g, b int
		for i := 0; i < len(game.Sets); i++ {
			if game.Sets[i].Blue > b {
				b = game.Sets[i].Blue
			}
			if game.Sets[i].Green > g {
				g = game.Sets[i].Green
			}
			if game.Sets[i].Red > r {
				r = game.Sets[i].Red
			}
		}

		res += r * g * b
	}

	fmt.Println("Result:", res)
}

var gamePrefixRegExp = regexp.MustCompile("Game (\\d*):(.*)")

func parseRow(input string) (out Game) {
	matches := gamePrefixRegExp.FindStringSubmatch(input)
	out.ID, _ = strconv.Atoi(matches[1])

	sets := strings.Split(matches[2], ";")
	for i := 0; i < len(sets); i++ {
		ns := GameSet{}
		colors := strings.Split(strings.TrimSpace(sets[i]), ",")

		for j := 0; j < len(colors); j++ {
			p := strings.Split(strings.TrimSpace(colors[j]), " ")
			switch p[1] {
			case "blue":
				ns.Blue, _ = strconv.Atoi(p[0])
			case "red":
				ns.Red, _ = strconv.Atoi(p[0])
			case "green":
				ns.Green, _ = strconv.Atoi(p[0])
			}

		}

		out.Sets = append(out.Sets, ns)
	}

	return
}
