package main

import "fmt"

func main() {
	a := map[string]int64{}

	a["foo"] = 0

	v, _ := a["foo"]

	v += 1
	a["foo"] = v

	fmt.Println(a)
	fmt.Println(v)
}
