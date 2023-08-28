package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var input = "move 1 from 7 to 6\nmove 1 from 8 to 5\nmove 3 from 7 to 4\nmove 5 from 9 to 6\nmove 3 from 7 to 9\nmove 2 from 5 to 7\nmove 10 from 6 to 8\nmove 2 from 2 to 3\nmove 2 from 9 to 1\nmove 6 from 8 to 2\nmove 5 from 3 to 8\nmove 4 from 5 to 9\nmove 3 from 4 to 5\nmove 2 from 1 to 8\nmove 3 from 1 to 7\nmove 1 from 7 to 1\nmove 4 from 7 to 8\nmove 1 from 5 to 6\nmove 1 from 9 to 3\nmove 8 from 2 to 4\nmove 1 from 5 to 8\nmove 1 from 5 to 3\nmove 2 from 1 to 8\nmove 4 from 3 to 4\nmove 1 from 3 to 4\nmove 1 from 1 to 7\nmove 1 from 7 to 8\nmove 1 from 7 to 4\nmove 5 from 9 to 1\nmove 2 from 6 to 7\nmove 3 from 2 to 1\nmove 12 from 8 to 7\nmove 8 from 7 to 3\nmove 1 from 2 to 8\nmove 6 from 7 to 1\nmove 1 from 6 to 3\nmove 8 from 4 to 3\nmove 5 from 3 to 6\nmove 6 from 1 to 8\nmove 2 from 1 to 2\nmove 2 from 3 to 1\nmove 4 from 4 to 5\nmove 1 from 5 to 7\nmove 1 from 6 to 9\nmove 1 from 4 to 9\nmove 8 from 1 to 4\nmove 10 from 3 to 5\nmove 2 from 4 to 5\nmove 2 from 2 to 6\nmove 2 from 1 to 6\nmove 11 from 4 to 7\nmove 9 from 6 to 5\nmove 16 from 8 to 3\nmove 15 from 5 to 6\nmove 10 from 3 to 6\nmove 24 from 6 to 5\nmove 5 from 7 to 5\nmove 1 from 6 to 3\nmove 1 from 7 to 2\nmove 2 from 7 to 6\nmove 3 from 3 to 6\nmove 8 from 5 to 1\nmove 3 from 9 to 8\nmove 3 from 8 to 4\nmove 1 from 7 to 1\nmove 1 from 2 to 9\nmove 1 from 9 to 2\nmove 2 from 3 to 1\nmove 2 from 4 to 2\nmove 5 from 6 to 8\nmove 3 from 7 to 1\nmove 1 from 4 to 2\nmove 26 from 5 to 9\nmove 1 from 3 to 6\nmove 7 from 1 to 9\nmove 1 from 3 to 5\nmove 1 from 6 to 5\nmove 1 from 5 to 4\nmove 5 from 5 to 6\nmove 1 from 4 to 9\nmove 3 from 9 to 3\nmove 4 from 8 to 5\nmove 2 from 5 to 2\nmove 1 from 1 to 6\nmove 1 from 8 to 9\nmove 2 from 2 to 4\nmove 2 from 3 to 7\nmove 1 from 7 to 6\nmove 7 from 6 to 7\nmove 1 from 4 to 3\nmove 2 from 2 to 4\nmove 28 from 9 to 3\nmove 26 from 3 to 7\nmove 2 from 4 to 3\nmove 2 from 9 to 1\nmove 4 from 3 to 6\nmove 1 from 4 to 5\nmove 1 from 3 to 4\nmove 3 from 1 to 9\nmove 1 from 4 to 7\nmove 1 from 5 to 7\nmove 1 from 6 to 9\nmove 23 from 7 to 1\nmove 4 from 9 to 5\nmove 3 from 9 to 4\nmove 2 from 6 to 3\nmove 1 from 6 to 7\nmove 3 from 3 to 9\nmove 11 from 7 to 2\nmove 4 from 2 to 3\nmove 23 from 1 to 2\nmove 15 from 2 to 4\nmove 2 from 7 to 9\nmove 13 from 2 to 8\nmove 1 from 7 to 5\nmove 1 from 2 to 8\nmove 7 from 4 to 8\nmove 6 from 4 to 3\nmove 1 from 2 to 4\nmove 1 from 2 to 9\nmove 20 from 8 to 5\nmove 1 from 8 to 4\nmove 3 from 4 to 7\nmove 3 from 3 to 9\nmove 1 from 2 to 8\nmove 20 from 5 to 3\nmove 6 from 5 to 3\nmove 26 from 3 to 9\nmove 2 from 7 to 5\nmove 1 from 5 to 4\nmove 1 from 7 to 8\nmove 2 from 8 to 5\nmove 12 from 9 to 4\nmove 2 from 3 to 2\nmove 4 from 1 to 9\nmove 2 from 3 to 1\nmove 4 from 5 to 6\nmove 5 from 9 to 4\nmove 2 from 6 to 3\nmove 2 from 6 to 8\nmove 2 from 8 to 3\nmove 1 from 2 to 7\nmove 21 from 4 to 2\nmove 1 from 4 to 5\nmove 13 from 2 to 4\nmove 4 from 3 to 9\nmove 25 from 9 to 7\nmove 7 from 2 to 4\nmove 18 from 7 to 8\nmove 2 from 1 to 5\nmove 1 from 3 to 9\nmove 2 from 9 to 3\nmove 1 from 1 to 6\nmove 8 from 7 to 6\nmove 4 from 3 to 2\nmove 1 from 4 to 7\nmove 6 from 2 to 5\nmove 1 from 7 to 3\nmove 5 from 6 to 8\nmove 4 from 4 to 1\nmove 9 from 5 to 1\nmove 12 from 4 to 3\nmove 1 from 6 to 5\nmove 1 from 5 to 2\nmove 13 from 3 to 8\nmove 14 from 8 to 6\nmove 2 from 1 to 6\nmove 1 from 2 to 5\nmove 11 from 1 to 3\nmove 1 from 5 to 3\nmove 6 from 6 to 8\nmove 23 from 8 to 5\nmove 1 from 8 to 1\nmove 18 from 5 to 8\nmove 5 from 6 to 8\nmove 10 from 3 to 8\nmove 1 from 1 to 5\nmove 2 from 4 to 8\nmove 1 from 4 to 7\nmove 5 from 5 to 3\nmove 1 from 6 to 1\nmove 6 from 3 to 9\nmove 35 from 8 to 4\nmove 1 from 7 to 6\nmove 2 from 9 to 8\nmove 1 from 1 to 6\nmove 17 from 4 to 7\nmove 1 from 5 to 1\nmove 4 from 9 to 6\nmove 12 from 6 to 4\nmove 29 from 4 to 2\nmove 17 from 7 to 8\nmove 27 from 2 to 7\nmove 2 from 2 to 1\nmove 1 from 3 to 1\nmove 25 from 7 to 4\nmove 25 from 4 to 6\nmove 1 from 4 to 2\nmove 4 from 1 to 6\nmove 1 from 2 to 6\nmove 25 from 6 to 1\nmove 5 from 6 to 8\nmove 15 from 1 to 6\nmove 2 from 7 to 8\nmove 15 from 6 to 2\nmove 14 from 2 to 8\nmove 1 from 2 to 3\nmove 4 from 1 to 4\nmove 4 from 4 to 2\nmove 6 from 1 to 8\nmove 3 from 2 to 5\nmove 3 from 5 to 7\nmove 1 from 2 to 3\nmove 1 from 6 to 8\nmove 8 from 8 to 5\nmove 2 from 7 to 4\nmove 1 from 7 to 9\nmove 3 from 5 to 8\nmove 2 from 4 to 6\nmove 3 from 5 to 8\nmove 2 from 3 to 4\nmove 2 from 6 to 5\nmove 1 from 9 to 8\nmove 48 from 8 to 5\nmove 1 from 8 to 9\nmove 41 from 5 to 4\nmove 4 from 5 to 2\nmove 3 from 2 to 7\nmove 1 from 2 to 7\nmove 1 from 8 to 1\nmove 1 from 9 to 4\nmove 1 from 1 to 3\nmove 7 from 4 to 7\nmove 11 from 7 to 4\nmove 4 from 4 to 1\nmove 37 from 4 to 9\nmove 4 from 4 to 3\nmove 32 from 9 to 3\nmove 5 from 9 to 1\nmove 12 from 3 to 2\nmove 3 from 4 to 1\nmove 3 from 1 to 6\nmove 3 from 1 to 6\nmove 2 from 1 to 5\nmove 9 from 2 to 7\nmove 3 from 7 to 3\nmove 6 from 6 to 5\nmove 4 from 3 to 6\nmove 3 from 6 to 9\nmove 13 from 3 to 8\nmove 3 from 1 to 9\nmove 2 from 3 to 2\nmove 2 from 7 to 8\nmove 1 from 6 to 8\nmove 4 from 2 to 8\nmove 2 from 8 to 3\nmove 1 from 2 to 1\nmove 4 from 7 to 3\nmove 6 from 3 to 5\nmove 3 from 9 to 8\nmove 13 from 8 to 6\nmove 1 from 9 to 2\nmove 2 from 3 to 8\nmove 1 from 1 to 9\nmove 1 from 1 to 3\nmove 10 from 6 to 3\nmove 1 from 2 to 5\nmove 22 from 5 to 7\nmove 1 from 9 to 3\nmove 1 from 8 to 7\nmove 2 from 7 to 8\nmove 6 from 8 to 4\nmove 2 from 9 to 2\nmove 21 from 7 to 6\nmove 4 from 8 to 5\nmove 1 from 8 to 4\nmove 1 from 5 to 7\nmove 12 from 3 to 6\nmove 1 from 2 to 6\nmove 1 from 7 to 9\nmove 1 from 2 to 6\nmove 6 from 3 to 5\nmove 6 from 4 to 2\nmove 1 from 3 to 6\nmove 1 from 9 to 7\nmove 6 from 2 to 7\nmove 22 from 6 to 4\nmove 3 from 6 to 5\nmove 7 from 5 to 7\nmove 3 from 7 to 8\nmove 2 from 5 to 3\nmove 2 from 3 to 7\nmove 13 from 6 to 8\nmove 3 from 7 to 1\nmove 3 from 5 to 9\nmove 16 from 4 to 5\nmove 1 from 5 to 8\nmove 2 from 1 to 6\nmove 1 from 1 to 7\nmove 6 from 4 to 2\nmove 4 from 8 to 7\nmove 13 from 5 to 7\nmove 1 from 6 to 3\nmove 2 from 5 to 6\nmove 10 from 7 to 6\nmove 1 from 3 to 9\nmove 1 from 4 to 3\nmove 1 from 3 to 5\nmove 12 from 7 to 3\nmove 2 from 2 to 1\nmove 1 from 5 to 9\nmove 2 from 9 to 6\nmove 4 from 2 to 7\nmove 7 from 7 to 9\nmove 1 from 7 to 8\nmove 1 from 1 to 9\nmove 11 from 9 to 7\nmove 4 from 8 to 3\nmove 5 from 3 to 5\nmove 2 from 8 to 4\nmove 3 from 5 to 2\nmove 2 from 2 to 8\nmove 1 from 5 to 2\nmove 5 from 8 to 2\nmove 7 from 7 to 2\nmove 4 from 8 to 9\nmove 2 from 7 to 6\nmove 4 from 9 to 7\nmove 6 from 2 to 4\nmove 1 from 5 to 6\nmove 5 from 3 to 5\nmove 1 from 8 to 1\nmove 10 from 6 to 3\nmove 8 from 2 to 8\nmove 1 from 8 to 1\nmove 5 from 3 to 2\nmove 2 from 8 to 7\nmove 6 from 7 to 4\nmove 12 from 4 to 1\nmove 4 from 1 to 2\nmove 1 from 2 to 1\nmove 8 from 2 to 9\nmove 2 from 4 to 8\nmove 5 from 9 to 7\nmove 8 from 3 to 8\nmove 2 from 3 to 1\nmove 6 from 8 to 2\nmove 7 from 7 to 2\nmove 1 from 3 to 5\nmove 2 from 7 to 2\nmove 1 from 9 to 1\nmove 1 from 9 to 7\nmove 1 from 9 to 4\nmove 1 from 6 to 7\nmove 1 from 2 to 3\nmove 1 from 3 to 8\nmove 1 from 4 to 9\nmove 5 from 6 to 1\nmove 7 from 8 to 2\nmove 1 from 7 to 4\nmove 9 from 2 to 8\nmove 7 from 2 to 7\nmove 1 from 4 to 2\nmove 8 from 7 to 5\nmove 4 from 8 to 7\nmove 8 from 8 to 6\nmove 9 from 1 to 4\nmove 1 from 9 to 1\nmove 4 from 7 to 6\nmove 7 from 1 to 7\nmove 6 from 7 to 3\nmove 4 from 1 to 8\nmove 13 from 6 to 3\nmove 6 from 2 to 3\nmove 1 from 3 to 4\nmove 2 from 3 to 7\nmove 1 from 6 to 9\nmove 11 from 5 to 1\nmove 1 from 6 to 3\nmove 8 from 4 to 1\nmove 2 from 5 to 2\nmove 1 from 9 to 5\nmove 2 from 8 to 7\nmove 7 from 1 to 5\nmove 2 from 7 to 3\nmove 8 from 5 to 4\nmove 1 from 8 to 2\nmove 1 from 5 to 7\nmove 3 from 7 to 2\nmove 4 from 4 to 7\nmove 4 from 3 to 4\nmove 20 from 3 to 2\nmove 1 from 8 to 3\nmove 1 from 3 to 8\nmove 4 from 7 to 2\nmove 1 from 8 to 6\nmove 1 from 7 to 5\nmove 1 from 3 to 1\nmove 1 from 4 to 2\nmove 5 from 1 to 4\nmove 14 from 4 to 1\nmove 1 from 6 to 5\nmove 1 from 2 to 3\nmove 1 from 5 to 1\nmove 11 from 2 to 9\nmove 18 from 1 to 2\nmove 4 from 1 to 3\nmove 12 from 2 to 5\nmove 5 from 2 to 4\nmove 7 from 5 to 1\nmove 1 from 2 to 9\nmove 9 from 1 to 9\nmove 1 from 3 to 6\nmove 2 from 3 to 9\nmove 1 from 6 to 1\nmove 1 from 4 to 8\nmove 1 from 3 to 4\nmove 1 from 3 to 8\nmove 16 from 9 to 5\nmove 2 from 2 to 7\nmove 14 from 5 to 8\nmove 16 from 8 to 5\nmove 1 from 7 to 9\nmove 1 from 7 to 6\nmove 4 from 9 to 5\nmove 11 from 5 to 6\nmove 12 from 2 to 4\nmove 16 from 5 to 7\nmove 4 from 7 to 2\nmove 1 from 5 to 6\nmove 3 from 9 to 1\nmove 4 from 7 to 9\nmove 3 from 6 to 4\nmove 9 from 2 to 9\nmove 3 from 1 to 8\nmove 2 from 8 to 1\nmove 1 from 8 to 2\nmove 5 from 6 to 1\nmove 7 from 7 to 1\nmove 1 from 7 to 6\nmove 8 from 4 to 5\nmove 1 from 2 to 6\nmove 12 from 9 to 2\nmove 3 from 2 to 9\nmove 8 from 5 to 8\nmove 12 from 4 to 5\nmove 1 from 2 to 9\nmove 1 from 5 to 6\nmove 2 from 1 to 7\nmove 4 from 5 to 2\nmove 6 from 5 to 1\nmove 2 from 7 to 6\nmove 1 from 5 to 1\nmove 1 from 8 to 5\nmove 7 from 6 to 9\nmove 2 from 9 to 4\nmove 16 from 1 to 8\nmove 1 from 5 to 8\nmove 7 from 2 to 8\nmove 3 from 6 to 2\nmove 1 from 4 to 8\nmove 28 from 8 to 3\nmove 1 from 4 to 2\nmove 4 from 1 to 2\nmove 11 from 2 to 7\nmove 9 from 7 to 8\nmove 7 from 9 to 5\nmove 4 from 8 to 1\nmove 2 from 9 to 1\nmove 2 from 1 to 5\nmove 1 from 7 to 9\nmove 1 from 1 to 9\nmove 6 from 5 to 3\nmove 3 from 5 to 1\nmove 2 from 2 to 8\nmove 7 from 8 to 3\nmove 7 from 3 to 7\nmove 4 from 1 to 9\nmove 1 from 8 to 9\nmove 2 from 8 to 1\nmove 1 from 8 to 1\nmove 6 from 7 to 6\nmove 6 from 6 to 5\nmove 17 from 3 to 6\nmove 2 from 9 to 2\nmove 2 from 1 to 4\nmove 12 from 3 to 8\nmove 6 from 6 to 5\nmove 2 from 2 to 1\nmove 4 from 9 to 7\nmove 2 from 7 to 3\nmove 1 from 1 to 5\nmove 10 from 8 to 6\nmove 2 from 3 to 9\nmove 9 from 5 to 2\nmove 7 from 2 to 8\nmove 1 from 4 to 8\nmove 1 from 4 to 6\nmove 7 from 8 to 7\nmove 3 from 9 to 7\nmove 4 from 3 to 4"

type stack []string

func (s stack) Push(i []string) stack {
	r := s
	for _, is := range i {
		r = append(r, is)
	}

	return r
}

func (s stack) Pop(q int) (stack, []string) {
	l := len(s)

	return s[0 : l-q], s[l-q : l]
}

func (s stack) Peek() string {
	l := len(s)

	return s[l-1]
}

var re = regexp.MustCompile("move (\\d*) from (\\d*) to (\\d*)")

func main() {
	stacks := initStacks()

	chunks := strings.Split(input, "\n")

	for _, line := range chunks {
		matches := re.FindStringSubmatch(line)
		qty, _ := strconv.ParseInt(matches[1], 10, 32)
		from := matches[2]
		to := matches[3]

		newStack, crates := stacks[from].Pop(int(qty))
		stacks[from] = newStack

		stacks[to] = stacks[to].Push(crates)
	}

	res := []string{
		stacks["1"].Peek(),
		stacks["2"].Peek(),
		stacks["3"].Peek(),
		stacks["4"].Peek(),
		stacks["5"].Peek(),
		stacks["6"].Peek(),
		stacks["7"].Peek(),
		stacks["8"].Peek(),
		stacks["9"].Peek(),
	}
	fmt.Println(strings.Join(res, ""))
}

func initStacks() map[string]stack {
	r := make(map[string]stack)

	r["1"] = stack{"D", "L", "V", "T", "M", "H", "F"}
	r["2"] = stack{"H", "Q", "G", "J", "C", "T", "N", "P"}
	r["3"] = stack{"R", "S", "D", "M", "P", "H"}
	r["4"] = stack{"L", "B", "V", "F"}
	r["5"] = stack{"N", "H", "G", "L", "Q"}
	r["6"] = stack{"W", "B", "D", "G", "R", "M", "P"}
	r["7"] = stack{"G", "M", "N", "R", "C", "H", "L", "Q"}
	r["8"] = stack{"C", "L", "W"}
	r["9"] = stack{"R", "D", "L", "Q", "J", "Z", "M", "T"}

	return r
}
