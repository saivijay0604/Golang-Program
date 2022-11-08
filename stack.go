package main

import "fmt"

type StackStruck struct {
	num []int
}



func main(){
	stack := StackStruck{}
	stack.Push(1)
	stack.Push(2)
	stack.Display()
	stack.Pop()
	stack.Display()
}

func(ss *StackStruck) Push(a int){
	ss.num = append(ss.num, a)
}

func (ss *StackStruck) Pop(){
	ss.num = (ss.num)[ :len(ss.num)-1]
}

func (ss StackStruck) Display(){
	fmt.Println(ss.num)
}