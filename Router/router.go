package Router

import (
	"github.com/dbAPI/Controller"
	//"ReadFile/example/ConnectDB"
	"github.com/dbAPI/PgConnect"
	"github.com/gin-gonic/gin"
	_ "github.com/dbAPI/Controller"
	"log"
)
var env = new(Controller.Env)
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
	//CRUD operations for Student
	router.GET("student", env.SelectStudent)//select data
	router.GET("student/:id", env.SelectStudentByID) //select student by ID
	router.POST("student", env.InsertStudent) //Insert data
	router.PUT("student", env.UpdateStudent) //Update Date
	router.DELETE("student/:id", env.DeleteStudentByID) //Delete Data

	//CRUD operations for Employer
	router.GET("employer", env.SelectEmployer)//select data
	router.GET("employer/:id", env.SelectEmployerByID) //select student by ID
	router.POST("employer", env.InsertEmployer) //Insert data
	router.PUT("employer", env.UpdateEmployer) //Update Date
	router.DELETE("employer/:id", env.DeleteEmployerByID) //Delete Data

	router.GET("person", env.SelectPerson)//select data
	router.GET("person/:id", env.SelectPersonByID) //select student by ID
	router.POST("person", env.InsertPerson) //Insert data
	router.PUT("person", env.UpdatePerson) //Update Date
	router.DELETE("person/:id", env.DeletePerson) //Delete Data





	router.Run("localhost:8081")

}


