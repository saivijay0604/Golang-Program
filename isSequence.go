package main

import "fmt"

func main(){
	s:="abc"
	t:="ambhc"
	result:=isSubsequence(s,t)
	fmt.Println(result)
}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i != len(s) && j != len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == len(s)
}
