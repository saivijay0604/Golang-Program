package PgConnect

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (env Env) InsertStudent(c *gin.Context) {
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	str:=c.Param("str")
	if str == "insert"{
		CheckDB()
	//	c.JSON(http.StatusOK,"connected")
		var newStudent Student
		if err := c.BindJSON(&newStudent); err != nil {
			log.Printf("invalid JSON body: %v", err)
			makeGinResponse(c, http.StatusNotFound, err.Error())
			return
		}

		q := `INSERT INTO student(id,name,result) VALUES($1,$2,$3)`
		result, err := env.DB.Exec(q, newStudent.ID,newStudent.Name,newStudent.Result)
		if err != nil {
			log.Printf("error occurred while inserting new record into artist table: %v", err)
			makeGinResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// checking the number of rows affected
		n, err := result.RowsAffected()
		if err != nil {
			log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
			makeGinResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// if no record was inserted, let us say client has failed
		if n == 0 {
			e := "could not insert the record, please try again after sometime"
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, e)
			return
		}
		m := "successfully created the record"
		log.Println(m)
		makeGinResponse(c, http.StatusOK, m)
	}
}
