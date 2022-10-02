package main

import "fmt"

func main(){
	n:=5
	fmt.Println(shape(n))
}

func shape(n int) int {
	num := 1
	for i := 1; i <= n; i++{
		num = num + (4 * (i-1))
		fmt.Println(num)
	}
	return num

}
