package main

import (
	"bufio"
	"fmt"
	"os"
)

type People struct {
	firstname, lastname string
}

func main() {
	ppl := []People{}
	if f, err := os.Open("afile.txt"); err != nil {
		fmt.Printf("open file error: %s\n", err)
		os.Exit(1)
	} else {
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			p := People{}
			if _, err := fmt.Sscanf(scanner.Text(), "%s %s", &p.firstname, &p.lastname); err != nil {
				continue
			}
			ppl = append(ppl, p)
		}
	}

	for _, p := range ppl {
		fmt.Printf("%s %s\n", p.firstname, p.lastname)
	}
}
