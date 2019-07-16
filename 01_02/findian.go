package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	if len(s) >= 3 && s[0] == 'i' && s[len(s)-1] == 'n' && s[1:len(s)-2] == strings.Repeat("a", len(s)-3) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
