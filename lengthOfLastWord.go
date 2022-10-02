package main

import "fmt"

func main(){
	str:= "Hello world"
	length:=lengthOfLastWord(str)
	fmt.Println(length)
}

func lengthOfLastWord(s string) int {
	length:=0
	for v := range s {
		if s[v] != ' ' {
			if v > 0 && s[v-1] == ' '{
				length = 0
			}
			length++
		}
	}
	return length

}
