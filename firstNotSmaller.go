package main

import "fmt"

func main() {
	arr:= []int{1,3,4,6,7,8}
	target:= 6
	res := firstNotSmaller(arr, target)
	fmt.Println(res)
}

func firstNotSmaller(arr []int, target int) int {
	left := 0
	right :=  len(arr)-1
	boundaryIndex := -1
	for	left <= right {
		mid := left + right // 2
		if arr[mid] >= target {
			boundaryIndex = mid
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return boundaryIndex
}