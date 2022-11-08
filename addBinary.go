package main

import (
	"fmt"
	"math/big"
	"strconv"
)


func main(){
	binaryNum1:="1101"
	binaryNum2:="01"
	result:= addBinary(binaryNum1,binaryNum2)
	num, _ := new(big.Int).SetString(result, 2)
	fmt.Print("sum of binary code in number ",num," and Binary code ",result)

}
func addBinary(a string, b string) string {
	carry, result := 0, ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		if sum >= 2 {
			carry = 1
			sum -= 2
		} else {
			carry = 0
		}
		result = strconv.Itoa(sum) + result
	}
	return result
}
