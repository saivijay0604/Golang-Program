package main

import (
	"fmt"
	"sort"
)



func main(){
	str:= []string{"tab","eat","ate","bat","tea","ten","net"}
	anagramWords :=anagarm(str)
	fmt.Println(anagramWords)

}

func anagarm(str []string) [][]string{
	m := make(map[string][]string)

	//sort the word in an array string and
	for _, v := range str {
		bytes := []byte(v)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		s := string(bytes)
		m[s] = append(m[s], v)
	}
	var ss [][]string
	for e := range m {
		ss = append(ss, m[e])
	}
	return ss
}



