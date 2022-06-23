package main

import (
	"fmt"
)

func main() {
	input := []int{4, 1, 4, -4, 6, 3, 8, 8}
	m := make(map[int]int)
	var result []int
	for _, number := range input {
		if _, ok := m[number]; ok {
			continue
		}
		m[number]++
		result = append(result, number)
	}

	fmt.Println(result)
}
