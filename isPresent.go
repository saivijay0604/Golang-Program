package main

import (
	"fmt"
)

func main(){
	s:="SST"
	x:="ST"
	result:=isPresent(s,x)
	fmt.Println(result)
}

func isPresent(s string, x string) int {


	result:= 0
	var str string

	for i:=0;i<len(s); i++{
		//fmt.Println(s[i],s[j])
		b:=false
		if s[i] == x[0] || b==true {
			b= true
			str = str + string(s[i])
			if str == x {
				return result
			}
			continue
		}
		i++
	}
	return -1
}

