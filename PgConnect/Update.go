package PgConnect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (env Env) UpdateStudent(c *gin.Context) {
		var toBeUpdatedStudent Student
		if err := c.BindJSON(&toBeUpdatedStudent); err != nil {
			e := fmt.Sprintf("invalid JSON body: %v", err)
			log.Println(e)
			makeGinResponse(c, http.StatusBadRequest, e)
			return
		}

		q := `UPDATE student SET name=$1,result=$2 WHERE id=$3;`
		_, err := env.DB.Exec(q, toBeUpdatedStudent.Name, toBeUpdatedStudent.Result, toBeUpdatedStudent.ID)
		if err != nil {
			e := fmt.Sprintf("error: %v occurred while updating artist record with id: %d", err, toBeUpdatedStudent.ID)
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, e)
			return
		}
		m := "successfully updated the record"
		log.Println(m)
		makeGinResponse(c, http.StatusAccepted, m)

}

