package PgConnect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Create(c *gin.Context) {

	str := c.Param("str")
	if str == "create" {
		env.DB, err = CheckDB()
		q := `create table student (id int,name VARCHAR ( 50 ),result int )`
		result, err := env.DB.Exec(q)

		//n, err := result.RowsAffected()
		if err != nil {
			e := fmt.Sprintf("error occurred while checking the returned result from database after creating the table: %v", err)
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, e)
			return
		}

		_, err = result.RowsAffected()
		if err != nil {
			log.Printf("error occurred while checking the returned result from database after creating the table: %v", err)
			makeGinResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		m := "successfully created the table"
		log.Println(m)
		makeGinResponse(c, http.StatusOK, m)
	}
}

