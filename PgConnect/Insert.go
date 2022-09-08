package PgConnect

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (env Env) InsertStudent(c *gin.Context) {


		var newStudent Student
		//fmt.Println(newStudent)
		if err := c.BindJSON(&newStudent); err != nil {
			log.Printf("invalid JSON body: %v", err)
			makeGinResponse(c, http.StatusNotFound, err.Error())
			return
		}

		q := `INSERT INTO student(id,name,result) VALUES($1,$2,$3)`
		_, err := env.DB.Exec(q, newStudent.ID,newStudent.Name,newStudent.Result)
			if err != nil {
				log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
				makeGinResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			m := "successfully created the record"
			log.Println(m)
			makeGinResponse(c, http.StatusCreated, m)

}

