package main

import "fmt"

func main() {
	a, v, s := float64(0), float64(0), float64(0)
	if _, err := fmt.Scanf("%f %f %f", &a, &v, &s); err != nil {
		fmt.Println("Error on: ", err)
	}
	fn := GenDisplaceFn(a, v, s)
	for {
		t := float64(-1)
		if _, err := fmt.Scanf("%f", &t); err != nil {
			break
		}
		fmt.Println(fn(t))
	}
}

func GenDisplaceFn(a, v, s float64) func (float64) (float64) {
	return func(t float64) (float64) {
		return a * t * t / 2 + v * t + s
	}
}
