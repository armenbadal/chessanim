package main

import (
	"fmt"
	"regexp"
)

var (
	columns = map[byte]int{
		'a': 0,
		'b': 1,
		'c': 2,
	}
	rows = map[byte]int{
		'1': 0,
		'2': 1,
		'3': 2,
	}
)

func indices(cell string) (int, int) {
	pat := regexp.MustCompile("^([a-h])([1-9])$")
	sm := pat.FindStringSubmatch(cell)
	return columns[sm[1][0]], rows[sm[2][0]]
}

func main() {
	s0 := "b3"
	c, r := indices(s0)
	fmt.Printf("(%d, %d)\n", c, r)
}
