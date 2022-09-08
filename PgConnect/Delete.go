package PgConnect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (env Env) DeleteStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e := fmt.Sprintf("received invalid id path param which is not string: %v", c.Param("id"))
		log.Println(e)
		makeGinResponse(c, http.StatusNotFound, e)
		return
	}
	q := `delete from student where id = $1;`
	_, err = env.DB.Exec(q,id)
		if err != nil {
			e := fmt.Sprintf("error occurred while deleting artist record with id: %d and error is: %v", id, err)
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, e)
			return
		}
		m := "successfully deleted the record"
		log.Println(m)
		makeGinResponse(c, http.StatusOK, m)
}
	//c.JSON(http.StatusOK,"deleted")

