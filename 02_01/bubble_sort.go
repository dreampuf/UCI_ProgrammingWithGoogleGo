package main

import "fmt"

func main() {
	slice := []int{}
	d := -1

	for {
		if _, err := fmt.Scanf("%d", &d); err != nil {
			break
		}
		slice = append(slice, d)
	}

	BubbleSort(slice)

	for _, i := range slice {
		fmt.Println(i)
	}


}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice) - 1; i ++ {
		for j := 0; j < len(slice) - 1; j ++ {
			//fmt.Println(slice[j], " -- ", slice[j+1])
			if slice[j] > slice[j+1] {
				//fmt.Println("swap: ", j)
				//fmt.Println("Before:", slice)
				Swap(slice, j)
				//fmt.Println("After:", slice)
			}
		}
	}
}

func Swap(slice []int, pos int) {
	slice[pos], slice[pos+1] = slice[pos+1], slice[pos]
}
