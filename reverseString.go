package main

import "fmt"

func main(){
	s:="Hello"
	rs :=reverseString(s)
	fmt.Println(rs)

	rs1 :=reverseString1(s)
	fmt.Println(rs1)
}
 func reverseString(s string) string {
	 var reverse  string
		 for _, v := range s {
				fmt.Println(string(v))
			 reverse = string(v) + reverse

		 }
		 if s == reverse{
		 	fmt.Println("Palindrom")
		 }else{
		 	fmt.Println("Not palindrom")
		 }
		 return reverse
 }

 func reverseString1(str string)string{
 	reverse :=[]rune(str)
 	for i,j:=0,len(str)-1;i<j;i,j=i+1,j-1{
 		reverse[i],reverse[j] = reverse[j],reverse[i]
	 }
	 if str == string(reverse){
		 fmt.Println("Palindrom")
	 }else{
		 fmt.Println("Not palindrom")
	 }

 	return string(reverse)
 }