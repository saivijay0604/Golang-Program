package Controller

import (
	"database/sql"
	"fmt"
	"github.com/dbAPI/PgConnect"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"strconv"
)
type Student struct {

	ID       int  `json:"id"`
	Name     string `json:"name"`
	Result     int  `json:"result"`
//	tableName string `json:"-"`
}
type Env struct{
	DB *sql.DB
}




//Select Student

func NewStudent(id int, name string, result int) Student {
	return Student{id, name, result}
}
var id int
var name string
var result int

func (s *Student)Select() error{
	var c *gin.Context
	return SelectStudent(c)
}

func SelectStudent(c *gin.Context) error {
	db,err := PgConnect.ConnectDB()
	if err != nil{
		return err
	}
	defer db.Close()

	q := `select * from student`
	rows, err := db.Query(q)
	var a = make([]Student, 0)
	//fmt.Println(a)
	for rows.Next() {
		err = rows.Scan(&id, &name, &result)
		if err != nil {
			log.Printf("error occurred while reading the database rows: %v", err)
			makeGinResponse(c, http.StatusInternalServerError, err.Error())
			return err
		}
		a = append(a, NewStudent(id, name, result))
	}
	makeGinResponse(c, http.StatusOK,a)
	return  nil

}
 //Select by id
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
//Insert Student details
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


//Update Student details
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

//Delete Student by ID
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

func makeGinResponse(c *gin.Context, statusCode int, value interface{}) {
	fmt.Println(value)
	c.JSON(statusCode, gin.H{
		"response":   value,
		"statusCode": statusCode,
	})

}


