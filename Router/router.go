package Router

import (
	"github.com/dbAPI/PgConnect"
	"github.com/gin-gonic/gin"

)
var env1 = new(PgConnect.Env)

func ConnectRouter(){

	router := gin.Default()
	// connect to DB first
	router.GET("pgConnection", PgConnect.Connection) //check the DB connection
	router.GET("pgConnection/:str", PgConnect.Create) //Create Table
	router.POST("pgConnection/:str", env1.InsertStudent) //Insert data
	router.PUT("pgConnection/:str", env1.UpdateStudent) //Update Date
	router.DELETE("pgConnection/:id", env1.DeleteStudentByID) //Delete Data


	router.Run("localhost:8081")

}


