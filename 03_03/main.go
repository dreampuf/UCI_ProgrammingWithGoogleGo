package main

import (
	"fmt"
	"os"
	"sort"
	"sync"
)

func sortSubArray(ls []int) {
	fmt.Printf("%v\n", ls)
	sort.Ints(ls)
}

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func main() {
	c := 4
	s := []int{}
	var d int
	for {
		if _, err := fmt.Scanf("%d", &d); err != nil {
			break
		}
		s = append(s, d)
	}
	fmt.Printf("%v\n", s)
	if len(s) % 4 != 0 {
		fmt.Println("please make sure the length of input array must be multiple of four: current is ", len(s))
		os.Exit(1)
	}
	//s := []int{3, 6, 2, 5, 9, 4, 2, 4, 6, 9, 2, 4}
	n := len(s) / c
	l := sync.Mutex{}
	subs := [][]int{}
	wg := sync.WaitGroup{}
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func(i, n int) {
			subSlice := s[i*n : (i+1)*n]
			sortSubArray(subSlice)

			l.Lock()
			subs = append(subs, subSlice)
			l.Unlock()
			wg.Done()
		}(i, n)
	}
	wg.Wait()

	t1 := merge(subs[0], subs[1])
	t2 := merge(subs[2], subs[3])
	result := merge(t1, t2)
	fmt.Printf("%v\n", result)
}
