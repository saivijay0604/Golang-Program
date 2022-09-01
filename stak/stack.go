package stak

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"fmt"
	"net/http"
)

type StackStruck struct {
	num []int
}


//
//func main(){
//
//	stack := StackStruck{}
//	stack.Push(1)
//	stack.Push(2)
//	stack.Display()
//	stack.Pop()
//	stack.Display()
//}

func(ss *StackStruck) Push(c *gin.Context){
	var a int
	if err := c.BindJSON(&a);err!= nil {
		return
	}
	ss.num = append(ss.num, a)
	c.JSON(http.StatusCreated,ss.num)
}

func (ss *StackStruck) Pop(c *gin.Context){

	ss.num = (ss.num)[ :len(ss.num)-1]
	c.JSON(http.StatusOK,ss.num)
}

func (ss *StackStruck) Display(c *gin.Context){
	 c.JSON(http.StatusOK,ss.num)
}

func (ss *StackStruck) AddStackElements(c *gin.Context){

	sum:=0
	for _,v := range ss.num{
		sum =sum+ v
	}
	fmt.Println(sum)
	c.JSON(http.StatusOK,sum)
}
