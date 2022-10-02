package main

import "fmt"

func main(){
	inputArray:= []int{6,2,-5,3,7,8,1}
	fmt.Println(productOfNum(inputArray))
}

func productOfNum(inputArray []int) int {

	if len(inputArray) <=1 {
		return inputArray[0]
	}
	product := inputArray[0]*inputArray[1]
	for i:=1;i<len(inputArray)-1;i++{
		if product < inputArray[i]*inputArray[i+1]{
			product = inputArray[i]*inputArray[i+1]

		}
	}
	return product
}
