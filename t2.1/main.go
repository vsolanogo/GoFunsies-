package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 4, 4, 3, 6, 3}
	visited := map[int]bool{}
	res := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		n := arr[i]
		if visited[n] {
			continue
		}

		visited[n] = true
		res = append([]int{n}, res...)
	}

	for i := 0; i < len(res)-1; i++ {
		for j := 0; j < len(res)-i-1; j++ {
			if res[j] > res[j+1] {
				var flag = res[j]
				res[j] = res[j+1]
				res[j+1] = flag
			}
		}
	}

	fmt.Println(res)
}
