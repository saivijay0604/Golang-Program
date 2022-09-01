package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"api/stak"
)
//type personInfo struct {
//
//	Name string `json:"name"`
//	Age int `json:"age"`
//	Profession string `json: profession`
//	Created bool `json:created`
//
//}
//
// var personDetails = []personInfo{
//	{Name:  "Vijay", Age: 29, Profession: "Software", Created: true},
//	{Name: "Vimal", Age: 5, Profession: "Student", Created: true},
//	{Name: "sai", Age: 20, Profession: " ", Created: false},
//}
 var ss = make([]int,0)
var str = stak.StackStruck{}

func getStack(context *gin.Context){
	context.IndentedJSON(http.StatusOK, ss)
}

func main(){
	ss = append(ss,1)
	ss = append(ss,2)
	ss = append(ss,5)
	router := gin.Default()
	router.GET("/ss", getStack)
	router.POST("/push",str.Push)
	router.DELETE("/pop", str.Pop)
	router.GET("display", str.Display)
	router.GET("/sum",str.AddStackElements)
	router.Run("localhost:8080")

}

