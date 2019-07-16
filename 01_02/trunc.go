package main

import "fmt"

func main() {
	var i float32
	if _, err := fmt.Scanf("%f\n", &i); err != nil {
		fmt.Printf("scanf gets error: %s\n", err)
	}
	fmt.Println(int32(i))
}
