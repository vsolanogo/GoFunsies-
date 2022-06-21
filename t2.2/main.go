package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fields := strings.Fields("1 9 3 4 -5")
	length := len(fields)
	var arr = make([]int, length, length)

	for i := 0; i <= length-1; i++ {
		arr[i], _ = strconv.Atoi(fields[i])
	}

	for i := 0; i < length-1; i++ {
		if arr[i] > arr[i+1] {
			var temp = arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}
	var maxValue = arr[length-1]

	var minIndex = 0
	for j := 1; j < length; j++ {
		if arr[minIndex] > arr[j] {
			minIndex = j
		}
	}
	var minValue = arr[minIndex]

	res := fmt.Sprintf("%d %d", maxValue, minValue)
	fmt.Println(res)
}
