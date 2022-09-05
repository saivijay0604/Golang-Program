package PgConnect

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
func NewStudent(id int, name string, result int) Student {
	return Student{id, name, result}
}
func (env Env) SelectStudent(c *gin.Context) {
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	str := c.Param("str")
	if str == "select" {
		env.DB,err = CheckDB()
		//	c.JSON(http.StatusOK,"connected")
		var newStudent Student
		if err := c.BindJSON(&newStudent); err != nil {
			log.Printf("invalid JSON body: %v", err)
			makeGinResponse(c, http.StatusNotFound, err.Error())
			return
		}

		a:= make([]Student,0)

		q := `select * from student`
		_,err := env.DB.Exec(q)
		//err = rows.Scan(newStudent.ID,newStudent.Name, newStudent.Result)
		//defer env.DB.QueryRow(q).Close()
		for i:=0;;i++ {
			var id int
			var name  string
			var result int
			//err = rows.Scan(&id, &name, result)
			if err != nil {
				log.Printf("error occurred while reading the database rows: %v", err)
				//rowReadErr = true
				break
			}
			a = append(a, NewStudent(id, name, result))
		}
		c.JSON(http.StatusOK, a)

}}

