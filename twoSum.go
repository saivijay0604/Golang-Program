package main

import "fmt"

func main(){
	nums:=[]int{1,7,2,4,11,8,6}
	target:=10
	fmt.Println(twoSum(nums,target))
}
//func twoSum(nums []int, target int) (int,int) {
//	for i:=0;i<len(nums);i++{
//		for j:=i+1;j<len(nums);j++{
//			if target== nums[i]+nums[j]{
//				return i,j
//			}
	//		}
	//	}
	//	return -1,-1
	//}



func twoSum(nums []int, target int) (int,int) {
	m:= make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := m[target - nums[i]]; ok {
			return val,i
		}
			m[nums[i]] = i

	}
	return -1,-1
}
