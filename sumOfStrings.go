package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str:= "1/4+1/4"
	num1,den1,num2,den2:= splitString(str)
	resultStr:= addFraction(num1, den1, num2, den2)
	fmt.Println(num1,"/",den1,"+",num2,"/",den2, "=",resultStr)
}

func splitString(str string)(int,int,int,int){

	x:=strings.Split(str,"+")
	y:=strings.Split(x[0],"/")
	n1,_:=strconv.Atoi(y[0])
	d1,_:=strconv.Atoi(y[1])
	y=strings.Split(x[1],"/")

	n2,_:=strconv.Atoi(y[0])
	d2,_:=strconv.Atoi(y[1])
	return n1,d1,n2,d2
}

func addFraction(num1,  den1,num2,  den2 int)string{
	num := (num1*den2) + (num2*den1)
	den := den1*den2
	i:=num
	if num> den{
		i=den
	}
	for ;i>=1;i-- {
		if num%i ==0 && den%i==0{
			num = num/i
			den =den/i
		}
	}
	result:= strconv.Itoa(num) + "/" + strconv.Itoa(den)
	return result
}
