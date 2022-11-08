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



//is two words are anagram or not

func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}

	sm:=make(map[string]int)
	tm:=make(map[string]int)
	for _,v:=range s{
		sm[string(v)]+=1
	}


	for _,v:=range t{
		tm[string(v)]+=1
	}

	for k1, v1 := range sm {
		if v2, ok := tm[k1]; ok {
			if v2 != v1 {
				return false
			}
			continue
		}
		return false
	}

	return true
}



