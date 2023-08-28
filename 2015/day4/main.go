package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

var key = "iwrupvqb"

func main() {
	limit := 10000000
	var hash [16]byte

	for i := 0; i < limit; i = i + 1 {
		hash = md5.Sum([]byte(key + strconv.FormatInt(int64(i), 10)))
		if hex.EncodeToString(hash[:])[0:6] == "000000" {
			fmt.Println(i)
			fmt.Println(hex.EncodeToString(hash[:]))
			break
		}
	}
}
