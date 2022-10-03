package main

import "fmt"

func main(){
	str:=[]string{"vijay","vijju","vimal"}
	result:=longestCommonPrefix(str)
	fmt.Println(result)
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := []byte{}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return string(result)
			}
		}
		result = append(result, strs[0][i])
	}
	return string(result)
}
