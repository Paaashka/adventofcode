package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("\\\\x([0-9a-f][0-9a-f])")

func main() {
	second()
}

func first() {
	file, err := os.Open("./day8/input")
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
	codeChars := 0
	memoryChars := 0
	for scanner.Scan() {
		codeChars += len(scanner.Text())

		s := re.ReplaceAllString(scanner.Text(), "1")
		s = strings.ReplaceAll(s, "\\\"", "\"")
		s = strings.ReplaceAll(s, "\\\\", "\\")

		memoryChars += len(s) - 2
	}

	fmt.Println(codeChars - memoryChars)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func second() {
	file, err := os.Open("./day8/input")
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
	codeChars := 0
	encodedChars := 0
	for scanner.Scan() {
		codeChars += len(scanner.Text())

		//s := re.ReplaceAllString(scanner.Text(), "1")
		s := scanner.Text()
		s = strings.ReplaceAll(s, "\\", "\\\\")
		s = strings.ReplaceAll(s, "\"", "\\\"")

		encodedChars += len(s) + 2
	}

	fmt.Println(encodedChars - codeChars)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
