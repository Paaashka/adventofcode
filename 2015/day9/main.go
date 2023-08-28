package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile("^([a-zA-Z]*) to ([a-zA-Z]*) = (\\d*)$")

//var regOp = regexp.MustCompile("([a-z0-9]*) (AND|OR|LSHIFT|RSHIFT) ([a-z0-9]*) -> ([a-z]*)")

// AlphaCentauri to Norrath = 147

type Roads map[string]map[string]int64

var roads = Roads{}

var cities = []string{
	"Straylight",
	"Arbre",
	"Tristram",
	"AlphaCentauri",
	"Snowdin",
	"Tambi",
	"Faerun",
	"Norrath",
}

//var cities = []string{
//	"London",
//	"Dublin",
//	"Belfast",
//}

func main() {
	first()
}

func first() {
	allCities := map[string]struct{}{}

	file, err := os.Open("./day9/input")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		matches := regex.FindStringSubmatch(scanner.Text())

		from, to := matches[1], matches[2]
		length, _ := strconv.ParseInt(matches[3], 10, 32)

		//fmt.Println(matches)
		//fmt.Println(scanner.Text())

		if _, ok := roads[from]; !ok {
			roads[from] = map[string]int64{}
		}

		if _, ok := roads[from][to]; !ok {
			roads[from][to] = length
		}

		if _, ok := roads[to]; !ok {
			roads[to] = map[string]int64{}
		}

		if _, ok := roads[to][from]; !ok {
			roads[to][from] = length
		}

		allCities[from] = struct{}{}
		allCities[to] = struct{}{}
	}

	var m int64
	for _, v := range cities {
		//c := lengthFrom(v, map[string]struct{}{})
		c := lengthFrom(v, allCities)
		if m == 0 || c > m {
			m = c
		}
	}

	fmt.Println(m)
}

func cloneMap(source map[string]struct{}) map[string]struct{} {
	target := map[string]struct{}{}

	for k := range source {
		target[k] = struct{}{}
	}

	return target
}

func lengthFrom(from string, cities map[string]struct{}) int64 {
	//fmt.Println(cities)
	newCities := cloneMap(cities)
	delete(newCities, from)

	var m int64
	for to := range newCities {
		//fmt.Print(from, " -> ", to)

		//if _, ok := newVisited[v]; ok || from == v {
		//	fmt.Println("...visited")
		//
		//	continue
		//}

		//fmt.Println(from, v)
		//fmt.Println("...ok")

		curr := roads[from][to] + lengthFrom(to, newCities)
		if m == 0 || curr > m {
			m = curr
		}
	}

	//fmt.Println("end of route")
	return m
}

//func lengthFrom(from string, visited map[string]struct{}) int64 {
//	newVisited := cloneMap(visited)
//	newVisited[from] = struct{}{}
//	//length += roads
//	fmt.Println(newVisited)
//	var m int64
//	for _, v := range cities {
//		if from == v {
//			continue
//		}
//
//		fmt.Print(from, " -> ", v)
//
//		if _, ok := newVisited[v]; ok || from == v {
//			fmt.Println("...visited")
//
//			continue
//		}
//
//		//fmt.Println(from, v)
//		fmt.Println("...ok")
//
//		if m == 0 || roads[from][v] < m {
//			m = roads[from][v] + lengthFrom(v, newVisited)
//		}
//	}
//
//	fmt.Println("end of route")
//	return m
//}
