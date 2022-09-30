package main

import "fmt"

func main(){
	var str string
	fmt.Print("Enter the Word:")
	fmt.Scanln(&str)
	result:= palindrome(str)
	fmt.Println(str,":",result)
}
func palindrome(str string)string{
	m:= make(map[string]int)
	for i:=0;i<len(str);i++ {
		x:= string(str[i])
		if _,Ok:=m[x];!Ok{
			m[x]++
			continue
		}
		m[x] =m[x]+1
	}
	cnt:= 0
	for _,v:= range m{
		if v%2==0 {
			cnt++
		}
	}
	if cnt==len(m)-1{
		return "word can form palindrome"
	}
	return "word can't form palindrome"
}
