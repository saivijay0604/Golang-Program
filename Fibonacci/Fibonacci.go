package Fibonacci

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func  FibSeq(c *gin.Context){
	fSeries:= make([]int,2)
	fSeries[0]= 0
	fSeries[1]= 1
	n:= c.Param("n")
	num, _:= strconv.Atoi(n)

	for i:=2;i<num;i++ {
		fSeries = append(fSeries, fSeries[i-1] + fSeries[i-2])
	}

	c.JSON(http.StatusOK,fSeries)
}
