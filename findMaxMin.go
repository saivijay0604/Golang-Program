package main

import (
	"fmt"
	"sort"
)

func main(){
	s:=[]int{1,2,3,4,5,5}
	max,min,sum:=MaxMin(s)
	fmt.Println("max number in the array: ",max)
	fmt.Println("Min number in an array: ",min)
	fmt.Println("sum of max and min: ",sum)

	maxS,minS:=MaxMinSort(s)
	fmt.Println("max number by using sorted package: ",maxS)
	fmt.Println("Min number by using sorted package: ",minS)


}
func MaxMin(nums []int) (int,int,int){

		max :=nums[0]
		min:=nums[0]
	for _,v:= range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	sum:= max+min

	return max,min,sum
}

func MaxMinSort(nums []int) (int,int){
	sort.Ints(nums)
	return nums[len(nums)-1],nums[0]

}
