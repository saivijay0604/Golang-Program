package main

import (
	"fmt"
	"strconv"
)


func main(){
	n:=7
	result:=isHappy(n)
	fmt.Println(result)
}
func isHappy(n int) bool {
	str := strconv.Itoa(n)
	sum:=0
	for i:=0;i<len(str);i++ {
		x, _ := strconv.Atoi(string(str[i]))
		product:= x*x
		sum+= product

	}
	fmt.Println(sum)
	if sum ==1{
		return false
	}


	isHappy(sum)

	return true
}


//for memory management

func isHappy1(n int) bool  {
	for n != 1 && n !=4{
		sum := 0
		for n != 0{
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}
	return n == 1
}