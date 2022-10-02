package main

import "fmt"

func main(){
	num:=20
	result:=isPerfectSquare(num)
	fmt.Println(result)
}

func isPerfectSquare(num int) bool {
	if num==1 {
		return true
	}else{
		for i:=1;i<num;i++{
			if i*i==num{
				return true
			}
		}
	}
	return false;
}
