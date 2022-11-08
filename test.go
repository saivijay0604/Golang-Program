package main

import "fmt"

func main(){
	str:="hello world"
	result:= duplicate(str)
	for k,v:= range result {
		fmt.Println(k,v)
	}

}
func duplicate(str string)map[string]int{
	result:= make(map[string]int)
	for _,v:=range str{

		ch:= string(v)
		if ch ==" "{
			continue
		}
		result[ch]=result[ch]+1
	}
	return result
}