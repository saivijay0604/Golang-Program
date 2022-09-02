package main

import (
	"api/Fibonacci"
	"api/stak"
	"github.com/gin-gonic/gin"
)

var str = stak.StackStruck{}
//var stud =PgConnect.Student{}

func main(){
	router := gin.Default()
	//stack
	router.POST("/push",str.Push)
	router.DELETE("/pop", str.Pop)
	router.GET("display", str.Display)
	//sum of number
	router.GET("/sum",str.AddStackElements)
	//Fib series
	router.GET("/Fib/:n", Fibonacci.FibSeq)


	//Local host
	router.Run("localhost:8080")

}

