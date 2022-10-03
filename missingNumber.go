package main

import (
	"fmt"
)

func main(){
	nums:=[]int{3,0,1}
	result:= missingNumber(nums)
	fmt.Println(result)
}

func missingNumber(nums []int) int {

	n := len(nums)
	sum := n * (n+1) / 2
	for _, v := range nums {
		sum -= v
	}

	return sum

}
