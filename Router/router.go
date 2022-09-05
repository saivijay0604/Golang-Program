package Router

import (
	"github.com/dbAPI/PgConnect"
	"github.com/gin-gonic/gin"
	"log"
)
var env = new(PgConnect.Env)
var err error


func ConnectRouter(){
	env.DB,err = PgConnect.CheckDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}


	router := gin.Default()
	// connect to DB first
	router.GET("pgConnection", PgConnect.Connection) //check the DB connection
	router.GET("pgConnection/:str", PgConnect.Create) //Create Table
	router.GET("pgConnection/select", env.SelectStudent) //select data
	router.POST("pgConnection/:str", env.InsertStudent) //Insert data
	router.PUT("pgConnection/:str", env.UpdateStudent) //Update Date
	router.DELETE("pgConnection/:id", env.DeleteStudentByID) //Delete Data


	router.Run("localhost:8081")

}


