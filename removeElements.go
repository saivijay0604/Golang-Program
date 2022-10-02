package main

import "fmt"

func main(){
	nums:=[]int{0,1,2,2,3,0,4,2}
	value:= 2
	result:=removeElement(nums,value)
	fmt.Println(result)
}
func removeElement(nums []int, val int) int {
	result:=0
	for i,_:= range nums{
		if nums[i]!=val{
			nums[result] = nums[i]
			result++

		}
	}

	return result

}

