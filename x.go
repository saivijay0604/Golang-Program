package main

import (
	"fmt"
	"sort"
)

func main(){
	s:=[]int{0,3}
	x:=solution(s)
	fmt.Println(x)
}

func solution(statues []int) int {
sort.Ints(statues)
	cnt:=0
	for i:=0; i<len(statues)-1; i++{
		cnt =cnt+ statues[i+1]-statues[i]-1
	}
	return cnt


}