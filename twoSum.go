package main

import "fmt"

func main(){
	nums:=[]int{3,6,4,8,1,10}
	target:=9
	result:=twoSum(nums,target)
	fmt.Println(result)
}
func twoSum(nums []int, target int) []int {
	result:= make([]int,0)
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if target== nums[i]+nums[j]{
				result=append(result,i)
				result=append(result,j)
				fmt.Println(i,j)

			}
		}

	}
	return result

}
