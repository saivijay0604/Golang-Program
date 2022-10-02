package main

import "fmt"

func main() {
	arr:= []bool{false,false,true,true}
	result := findBoundary(arr)
	fmt.Println(result)
}
func findBoundary(arr []bool) int {
	left := 0
	right := len(arr) - 1
	boundaryIndex := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] {
			boundaryIndex = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return boundaryIndex
}



