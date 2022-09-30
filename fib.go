package main


import "fmt"

func main(){
	//fib()
	var a, b int
	n:=10
	a = 0
	b = 1
	fmt.Print( a, b," ")
	printFibonacii(a, b, n)



	//fmt.Println("fib for non rec" ,fib(n))

}

//func fib(n int) []int {
//	s := make([]int,2)
//	s[0] =0
//	s[1] =1
//	for i:=2;i<=n;i++ {
//		s = append(s, s[i-2] + s[i-1])
//		if s[i] >= 100{
//			s =s[:i]
//			break
//		}
//
//	}
//	return s
//
//}


func printFibonacii(a int, b int, n int) {
	sum := 0
	if n > 2 {
		sum = a + b
		fmt.Printf("%d ", sum)
		a,b = b,sum
		printFibonacii(a, b, n-1)
	}
}



