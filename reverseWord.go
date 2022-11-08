package main

import (
	"fmt"
	"strings"
)
func main(){
	str:= "hello vijay! How are you?"
	fmt.Println(reverseWords(str))
}
func reverseWords(s string) string {
	str:= strings.Split(s, " ")
	var result string
	for i:=0;i<len(str);i++{
		if i==0{
			result=result+reverseStr(str[i])
			continue
		}
		result = result+" "+reverseStr(str[i])
	}
	return result
}


func reverseStr(s string)string{
	rev:=""
	for _,v:=range s{
		rev = string(v)+rev
	}
	return rev
}
