package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type EmployerDetails struct {
	EmployerId int `json:"id"`
	EmployerName string `json:"name"`
	Role string `json:"role"`
	Salary int `json:"salary"`
}

func NewEmployer(eId int, eName string, role string,salary int) EmployerDetails {
	return EmployerDetails{eId, eName, role,salary}
}
var eId int
var eName string
var role string
var salary int

func (env Env) SelectEmployer(c *gin.Context) {

	q := `select * from employer`
	rows, err := env.DB.Query(q)
	var emp = make([]EmployerDetails, 0)
	for rows.Next() {
		err = rows.Scan(&eId, &eName, &role,&salary)
		if err != nil {
			log.Printf("error occurred while reading the database rows: %v", err)
			makeGinResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		emp = append(emp, NewEmployer(eId, eName, role,salary))
	}
	makeGinResponse(c, http.StatusOK,emp)

}



func (env Env) SelectEmployerByID(c *gin.Context){
	eid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		er := fmt.Sprintf("received invalid id path param which is not string: %v", c.Param("eId"))
		log.Println(er)
		makeGinResponse(c, http.StatusNotFound, er)
		return
	}

	var e = make([]EmployerDetails, 0)
	q := `select * from employer where employerId=$1`
	res, err := env.DB.Query(q, eid)
	for res.Next() {
		err = res.Scan(&eId, &eName, &role,&salary)
		e = append(e, NewEmployer(eId, eName, role,salary))
		if err != nil {
			err := fmt.Sprintf("error occurred while selecting the record with id: %d and error is: %v", id, err)
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, err)
			return
		}

	}
	makeGinResponse(c, http.StatusOK,e)
}



func (env Env) InsertEmployer(c *gin.Context) {
	var newEmployer EmployerDetails
	//fmt.Println(newStudent)
	if err := c.BindJSON(&newEmployer); err != nil {
		log.Printf("invalid JSON body: %v", err)
		makeGinResponse(c, http.StatusNotFound, err.Error())
		return
	}

	q := `INSERT INTO employer(employerId ,employerName ,role ,salary) VALUES($1,$2,$3,$4)`
	_, err := env.DB.Exec(q, newEmployer.EmployerId,newEmployer.EmployerName,newEmployer.Role,newEmployer.Salary)
	if err != nil {
		log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
		makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	m := "successfully created the record"
	log.Println(m)
	makeGinResponse(c, http.StatusCreated, m)
}


func (env Env) UpdateEmployer(c *gin.Context){
	var updateEmployer EmployerDetails
	if err := c.BindJSON(&updateEmployer); err != nil {
		e := fmt.Sprintf("invalid JSON body: %v", err)
		log.Println(e)
		makeGinResponse(c, http.StatusBadRequest, e)
		return
	}
	q := `UPDATE employer SET employerName=$1, role=$2 ,salary=$3 WHERE employerId =$4;`
	_, err := env.DB.Exec(q,updateEmployer.EmployerName,updateEmployer.Role,updateEmployer.Salary, updateEmployer.EmployerId)
	if err != nil {
		log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
		makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	m := "successfully created the record"
	log.Println(m)
	makeGinResponse(c, http.StatusCreated, m)

}

func (env Env) DeleteEmployerByID(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e := fmt.Sprintf("received invalid id path param which is not string: %v", c.Param("id"))
		log.Println(e)
		makeGinResponse(c, http.StatusNotFound, e)
		return
	}
	q := `delete from employer where id = $1;`
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

