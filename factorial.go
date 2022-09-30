package main

import "fmt"
const LIM = 41
var facts [LIM]uint64
func main(){
	x :=fact(5)
	fmt.Println(x)
	result:=uint64(0)
	for i:=uint64(0); i < LIM; i++ {
		result = factorialMemoization(uint64(i))

	}
	fmt.Printf("The factorial value for 40 is %d\n", uint64(result))

}

func fact(n int)int{
	if n<=1{
		return 1
	}
	return n*fact(n-1)
}

func factorialMemoization(n uint64)(res uint64) {
	if (facts[n] != 0) {
		res = facts[n]
		return res
	}

	if (n > 0) {
		return n * factorialMemoization(n-1)
	}

	return 1
}
