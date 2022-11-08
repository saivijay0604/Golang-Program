package main

import (
	"fmt"
)

func main() {
	fmt.Println(bin(32))
}

func bin(n int) []int {
	//n := 10
	//fmt.Println(strconv.FormatInt(n, 2))
	//1 -0,1 												0,1
	//2 - 10,11, 												1,2
	//3 - 100,101,110,111 											1,2,2,3
	//4 - 1000,1001,1010,1011,1100,1101,1110,1111 								1,2,2,3,2,3,3,4
	//5 - 10000,10001,10010,10011,10100,10101,10110,10111,11000,11001,11010,11011,11100,11101,11110,11111   1,2,2,3,2,3,3,4,2,3,3,4,3,4,4,5

	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	ret := []int{0, 1}
	sub := []int{1}
	for {

		sub = double(sub)
		//fmt.Println(sub)
		ret = append(ret, sub...)
		if len(ret) > n {
			return ret[:n]
		}
	}

}

func double(sub []int) []int {
	ret := sub
	for i := 0; i < len(sub); i++ {
		ret = append(ret, sub[i]+1)
	}
	//fmt.Println(ret)
	return ret
}

//Other Method

//func countBits(n int) []int {
//	ret := []int{}
//	for i:=0;i<=n;i++{
//		ret = append(ret,increment(i))
//	}
//	return ret
//}
//
//func increment(n int)int{
//	binary := strconv.FormatInt(int64(n),2)
//	c := strings.Count(binary,"1")
//	return c
//
//}
