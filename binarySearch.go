package main

import (
	"fmt"
)
func main() {
	arr:= []int{1,2,3,4,6,7,8}
	target:= 2
	res := binarySearch(arr, target)
	fmt.Println(res)
}

func binarySearch(arr []int, target int) int {
	//sort.Ints(arr)
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
