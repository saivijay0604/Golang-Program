package main

import "fmt"

func main(){
	nums:=[]int{1,4,6,8}
	target:=5
	position:= searchInsert(nums,target)
	fmt.Println(position)
}

func searchInsert(nums []int, target int) int {
	var position int
	if nums[len(nums)-1] <target{
		position = len(nums)
	}else{
		for k,_:=range nums{
			if nums[k] < target && nums[k+1]>target  {
				position =k+1
			}
			if nums[k] ==target{
				position =k
			}
		}
	}
	return position
}
