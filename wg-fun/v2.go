package main

import (
	"fmt"
	"sync"
)

var n = [][]int{
	{2, 6, 9, 24},
	{7, 3, 94, 3, 0},
	{4, 2, 8, 35},
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < len(n); i++ {
		wg.Add(1)
		go func(j int) {
			res := sum(n[j])
			fmt.Printf("Slice %d returned : %d\n", n[j], res)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func sum(arr []int) int {
	res := 0
	for _, val := range arr {
		res += val
	}
	return res
}
