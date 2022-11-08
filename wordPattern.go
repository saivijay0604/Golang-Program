package main

import (
	"fmt"
	"strings"
)

func main(){
	pattern:= "abba"
	s:= "cat dog cat dog"
	fmt.Println(wordPattern(pattern,s))

}

func wordPattern(pattern string, s string) bool {

	str:= strings.Split(s, " ")
	if len(str) != len(pattern) {
		return false
	}
	m1 := make(map[byte]string, 0)
	m2 := make(map[string]byte, 0)
	for i := 0; i < len(pattern); i++ {
		v, ok := m1[pattern[i]]
		if !ok {
			if _, ok := m2[str[i]]; ok {
				return false
			}
			m1[pattern[i]] = str[i]
			m2[str[i]] = pattern[i]
			continue
		}
		if v != str[i] {
			return false
		}
	}
	return true
}