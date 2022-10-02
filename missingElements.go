package main

import (
	"fmt"
	"sort"
)



func main(){
	status:= []int{10,6,8,3,4}
	fmt.Println(totalMissElements(status))
}
func totalMissElements(statues []int) int {
	sort.Ints(statues)
	cnt:=0
	for i:=0; i<len(statues)-1; i++{
		cnt =cnt+ statues[i+1]-statues[i]-1
	}
	return cnt

}

//
//For statues = [6, 2, 3, 8], the output should be
//solution(statues) = 3.
//
//Ratiorg needs statues of sizes 4, 5 and 7.
