package main

import (
	"fmt"
	"strconv"
)

func main(){

	num1:="89623"
	num2:="923459"
	result,_:=strconv.Atoi(addStrings(num1,num2))
	fmt.Println(result)
}

func addStrings(num1 string, num2 string) string {
	carry,sum:=0,0

	result:=""
	i:=len(num1)-1
	j:=len(num2)-1
	for i>=0 ||j>=0{
		var x,y int
		if i>=0{
			x,_=strconv.Atoi(string(num1[i]))
		}else{x=0}

		if j>=0{
			y,_=strconv.Atoi(string(num2[j]))
		}else{y=0}

		sum =x+y+carry
		carry =sum/10
		result=strconv.Itoa(sum%10)+result
		i--
		j--

	}
	if carry!=0{
		return "1"+result
	}
	return result
}

