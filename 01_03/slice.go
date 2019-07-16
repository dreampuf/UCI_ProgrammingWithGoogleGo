package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var si = make([]int, 0)
	var s string
	var i int
	for {
		fmt.Scanf("%s", &s)
		if s == "X" {
			break
		} else {
			i, _ = strconv.Atoi(s)
		}
		si = append(si, i)
	}
	sort.Ints(si)
	fmt.Print(si)
}
