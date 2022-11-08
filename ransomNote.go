package main

import "fmt"

func main(){
	ransomNote:= "aa"
	magazine:= "aab"
	fmt.Println(canConstruct(ransomNote,magazine))
}

/* Given two strings ransomNote and magazine, return true if ransomNote can be
constructed by using the letters from magazine and false otherwise. */
func canConstruct(ransomNote string, magazine string) bool {
	mm := make(map[byte]int)

	for i := range magazine {
		mm[magazine[i]] += 1
	}

	for i := range ransomNote {
		if value, ok := mm[ransomNote[i]]; ok && value >= 1 {
			mm[ransomNote[i]] -= 1
			continue
		}

		return false
	}

	return true

}
