package main

import (
	"fmt"
	"strings"
)

func main(){
	str:= "abc(xyz)abc"
	result:= reversWordInBrackets(str)
	fmt.Println(result)
}

func reversWordInBrackets(str string)string{
	result:=""
	temp:=""
	var spStr []string
	for _,v:=range str{
		if string(v) == "("{
			spStr= strings.Split(str,"(")
			result =spStr[0]
			fmt.Println(result)
			temp =spStr[1]
		}else if string(v) == ")" {
			spStr = strings.Split(temp, ")")
			temp = spStr[0]
		}
	}
	result =result+reverseStr(temp)+spStr[1]


	return result
}

func reverseStr(s string) string {
	fmt.Println(s)
	var reverse string
	for _, v := range s {
		reverse = string(v) + reverse
	}
	return reverse
}