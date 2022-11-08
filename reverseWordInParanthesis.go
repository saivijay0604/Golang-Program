package main

import (
	"fmt"
	"strings"
)

//abc(xyz)abc --> abczyxabc
//abc(abc(xyz))abc --> abc (abc zyx) abc --> abc xyz cba abc

func main() {
	fmt.Println(reverse("abc(abc(xyz))abc"))
	//str := "abc(abc)abc"

}

func reverse(s string) string {
	fmt.Println("input:", s)
	if !strings.Contains(s, "(") {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		fmt.Println("***", string(runes))
		return string(runes)
	}
	i, j := -1, -1

	for idx, r := range s {
		if string(r) == "(" && i == -1 {
			i = idx
		}
		if string(r) == ")" {
		//if string(r) == ")" && j == -1{
			j = idx
		}
	}
	rev := reverse(s[i+1 : j])
	fmt.Println("output", s[:i]+rev+s[j+1:])
	//return s[:i] + rev + s[j+1:]

	newStr := s[:i] + rev + s[j+1:]
	//if strings.Contains(newStr, "(") {
	//	return reverse(newStr)
	//}
	return newStr

}


