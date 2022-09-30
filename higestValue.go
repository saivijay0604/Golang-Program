package main

import (
	"fmt"
)

func main(){
	m:= map[string]int{
		"Green": 5,
		"Red": 2,
		"Yellow": 8,
		"Blue": 8,

	}
	result:= highValueKey(m)
	fmt.Println(result)
}

func highValueKey(m map[string]int) string{
	var result string
	temp:=0
	for k,v:= range m{
		if temp<v || (v == temp  && result > k){
			temp = v
			result = k
		}
	}
	return result
}
