package main

import (
	"fmt"
	"sort"
)

func main(){
	nums:= []int{7,3,1,8,4}
	n:=60
	a,b,sum:= maxSum(nums,n)
	if sum==n {
		fmt.Println("the sum of index ", a, "+", b, "=", sum)
	}
}

func maxSum(nums []int,n int)(int,int,int){
	var a,b int
	sum:=0
	sort.Ints(nums)
	fmt.Println(nums)
	for i,j:= 0,len(nums)-1;i<j-1;{
		a=i
		b=j
		sum = nums[i]+nums[j]
		if sum == n{
			break
		}
		if sum>n {
			j--
			b = j
		}else {
			i++
			a = i
		}
	}
	if sum!=n{
		fmt.Println("Sum of Array don't have the given number",n)
	}
	return a,b,sum
}
