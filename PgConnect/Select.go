package PgConnect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)
func NewStudent(id int, name string, result int) Student {
	return Student{id, name, result}
}
var id int
var name string
var result int


func (env Env) SelectStudent(c *gin.Context) {

	q := `select * from student`
	rows, err := env.DB.Query(q)
	var a = make([]Student, 0)
	//fmt.Println(a)
	for rows.Next() {
		err = rows.Scan(&id, &name, &result)
		if err != nil {
			log.Printf("error occurred while reading the database rows: %v", err)
			makeGinResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		a = append(a, NewStudent(id, name, result))
	}
	makeGinResponse(c, http.StatusOK,a)

}

func (env Env) SelectStudentByID(c *gin.Context){
	//var newStudent Student
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e := fmt.Sprintf("received invalid id path param which is not string: %v", c.Param("id"))
		log.Println(e)
		makeGinResponse(c, http.StatusNotFound, e)
		return
	}

	var s = make([]Student, 0)
		q := `select * from student where id=$1`
		res, err := env.DB.Query(q, id)
		for res.Next() {
			err = res.Scan(&id, &name, &result)
			s = append(s, NewStudent(id, name, result))
				if err != nil {
					e := fmt.Sprintf("error occurred while selecting the record with id: %d and error is: %v", id, err)
					log.Println(e)
					makeGinResponse(c, http.StatusInternalServerError, e)
					return
				}

		}
	makeGinResponse(c, http.StatusOK,s)
}