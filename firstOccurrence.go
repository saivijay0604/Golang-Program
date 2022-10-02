package main

import (
	"fmt"
	"sort"
)
func main() {
	arr:= []int{1,3,4,4,4,6,7,8}
	target:= 4
	res := firstOccurrence(arr, target)
	fmt.Println(res)
}

func firstOccurrence(arr []int, target int) int {
	sort.Ints(arr)
	left := 0
	right := len(arr) - 1
	firstIndex:=-1
	for left <= right {
		mid := left + right
		if arr[mid] == target {
			firstIndex = mid
			right =mid-1

		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return firstIndex
}

