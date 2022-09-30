package main

import (
	"fmt"
	"sort"
)

func main(){
	s:= []int{1}
	buy,sell,profit:= stockPrice(s)
	fmt.Println("Buy the stock on day ",buy,"and sell on day ",sell,".You will get profit of $",profit)
}

func stockPrice(s []int)(int,int,int){
	//var indexBuy,priceBuy int
	tempBuy:=0
	tempSell:=0
	profit:=0

	for i:=0;i<len(s)-1;i++{
		for j:=i+1;j<len(s);j++ {
			if s[i] < s[j] && profit <s[j]-s[i]{
					profit = s[j]-s[i]
					tempBuy =i
					tempSell =j
				}
			}
		}

	return tempBuy, tempSell,profit
}

func stockProfit(s []int)int{
	profit:=0
	for i:=0;i<len(s)-1;i++{
		tempProfit := s[i] - max(s[i+1:])
		if tempProfit > profit{
			profit = tempProfit
		}
	}

	return profit
}

func max(n []int)int{
	sort.Ints(n)
	return n[len(n)-1]
}
