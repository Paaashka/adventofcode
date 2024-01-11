package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	dest   int64
	source int64
	len    int64
}

func main() {
	part1()
	part2()
}

func readFile() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(file)
}

func parseMap(in string) (res []Map) {
	lines := strings.Split(in, "\n")
	for i := 1; i < len(lines); i++ {
		ints := splitToInts(lines[i])
		res = append(res, Map{
			dest:   ints[0],
			source: ints[1],
			len:    ints[2],
		})
	}

	return res
}

func applyMaps(in int64, ms []Map) int64 {
	for _, m := range ms {
		if in >= m.source && in < m.source+m.len {
			in += m.dest - m.source
			break
		}
	}

	return in
}

func part1() {
	blocks := strings.Split(readFile(), "\n\n")

	// Parsing file
	seeds := splitToInts(strings.Replace(blocks[0], "seeds: ", "", 1))

	var maps [][]Map
	for i := 1; i < len(blocks); i++ {
		maps = append(maps, parseMap(blocks[i]))
	}

	// Applying maps
	var result int64
	for _, s := range seeds {
		for _, group := range maps {
			s = applyMaps(s, group)
		}

		if result == 0 || s < result {
			result = s
		}
	}

	fmt.Println("Result:", result)

	return
}

func part2() {
	blocks := strings.Split(readFile(), "\n\n")

	// Parsing file
	seeds := splitToInts(strings.Replace(blocks[0], "seeds: ", "", 1))

	var maps [][]Map
	for i := 1; i < len(blocks); i++ {
		maps = append(maps, parseMap(blocks[i]))
	}

	// Applying maps
	var result int64
	for i := 0; i < len(seeds); i += 2 {
		for sourceSeed := seeds[i]; sourceSeed < seeds[i]+seeds[i+1]; sourceSeed++ {
			//fmt.Printf("Seed %d ", sourceSeed)

			s := sourceSeed
			for _, group := range maps {
				s = applyMaps(s, group)
			}

			//fmt.Printf("became %d\n", s)

			if result == 0 || s < result {
				result = s
			}
		}
	}

	fmt.Println("Result:", result)

	return
}

func splitToInts(in string) (out []int64) {
	parts := strings.Split(in, " ")

	for _, el := range parts {
		i, _ := strconv.ParseInt(el, 10, 64)
		out = append(out, i)
	}

	return out
}
