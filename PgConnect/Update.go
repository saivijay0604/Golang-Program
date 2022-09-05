package PgConnect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (env Env) UpdateStudent(c *gin.Context) {
	// Call BindJSON to bind the received JSON to
	// toBeUpdatedAlbum.
	str:=c.Param("str")
	if str == "update" {
		CheckDB()
		//c.JSON(http.StatusOK,"connected")
		var toBeUpdatedStudent Student
		if err := c.BindJSON(&toBeUpdatedStudent); err != nil {
			e := fmt.Sprintf("invalid JSON body: %v", err)
			log.Println(e)
			makeGinResponse(c, http.StatusBadRequest, e)
			return
		}

		q := `UPDATE student SET name=$1,result=$2 WHERE id=$3;`
		result, err := env.DB.Exec(q, toBeUpdatedStudent.Name, toBeUpdatedStudent.Result, toBeUpdatedStudent.ID)
		if err != nil {
			e := fmt.Sprintf("error: %v occurred while updating artist record with id: %d", err, toBeUpdatedStudent.ID)
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, e)
			return
		}

		// checking the number of rows affected
		n, err := result.RowsAffected()
		if err != nil {
			e := fmt.Sprintf("error occurred while checking the returned result from database after updation: %v", err)
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, e)
		}

		// if no record was updated, let us say client has failed
		if n == 0 {
			e := "could not update the record, please try again after sometime"
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, e)
			return
		}

		m := "successfully updated the record"
		log.Println(m)
		makeGinResponse(c, http.StatusOK, m)
	}
}
