package main

import (
	"github.com/dbAPI/PgConnect"

	"github.com/gin-gonic/gin"
)


//var stud =PgConnect.Student{}

func main(){



	router := gin.Default()

	//DB
	// connect to DB first

	router.GET("postgreSQL", PgConnect.Connection) //check the DB connection
	router.GET("postgreSQL/:str", PgConnect.Crud)  //inseret
	////router.POST("postgreSQL/:str", PgConnect.Update) //create
	//router.POST("postgreSQL/:str", PgConnect.Delete) //create

	//Local host
	router.Run("localhost:8080")

}

