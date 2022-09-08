package Router

import (
	"github.com/dbAPI/PgConnect"
	"github.com/gin-gonic/gin"
	"log"
)
var env = new(PgConnect.Env)
var err error


func ConnectRouter(){
	//connect to Postgresql
	env.DB,err = PgConnect.ConnectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	defer env.DB.Close()

	//router's
	router := gin.Default()
	//CRUD operations

	router.GET("student", env.SelectStudent)//select data
	router.GET("student/:id", env.SelectStudentByID) //select student by ID
	router.POST("student", env.InsertStudent) //Insert data
	router.PUT("student", env.UpdateStudent) //Update Date
	router.DELETE("student/:id", env.DeleteStudentByID) //Delete Data


	router.Run("localhost:8081")

}


