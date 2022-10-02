package main

import "fmt"

func main(){
	year:= 1904
	fmt.Println(centureYear(year))
}
func centureYear(year int) int {
	if year <= 100{
		return 1
	} else if year % 100 == 0 {
		return year / 100
	} else{
		return year / 100 + 1
	}
}
