package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Person struct{
	PersonID int `json:"id"`
	PersonName string `json"name"`
	PersonAddress string `json"address"`
}


var personId int
var personName string
var address string

func (env Env) SelectPerson(c *gin.Context){
	q:=`select * from Person`
	rows,err:= env.DB.Query(q)
	var per = make([]Person,0)
	for rows.Next() {
		err = rows.Scan(&personId,&personName,&address)
		if err != nil {
			log.Printf("error occurred while reading the database rows: %v", err)
			makeGinResponse(c, http.StatusInternalServerError, err.Error())
			return

		}
		per =append(per,Person{PersonID: personId,PersonName: personName,PersonAddress: address})
	}
	makeGinResponse(c, http.StatusOK,per)
}

func (env Env) SelectPersonByID(c *gin.Context){
	 pId,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e := fmt.Sprintf("received invalid id path param which is not string: %v", c.Param("id"))
		log.Println(e)
		makeGinResponse(c, http.StatusNotFound, e)
		return
	}
	q:=`select * from Person WHERE personId =$1`
	rows,err:= env.DB.Query(q,pId)
	var per = make([]Person,0)
	for rows.Next() {
	err = rows.Scan(&personId,&personName,&address)
		per =append(per,Person{PersonID: personId,PersonName: personName,PersonAddress: address})
	if err != nil {
	log.Printf("error occurred while reading the database rows: %v", err)
	makeGinResponse(c, http.StatusInternalServerError, err.Error())
	return

	}

	}
	makeGinResponse(c, http.StatusOK,per)
}

func (env Env) InsertPerson(c *gin.Context){
	var newPersonData Person
	if err := c.BindJSON(&newPersonData); err != nil {
		log.Printf("invalid JSON body: %v", err)
		makeGinResponse(c, http.StatusNotFound, err.Error())
		return
	}
	q:=`INSERT INTO Person (personId ,personName ,address) VALUES ($1,$2,$3)`
	_,err:= env.DB.Exec(q,newPersonData.PersonID,newPersonData.PersonName,newPersonData.PersonAddress)
	if err != nil {
		log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
		makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	m := "successfully created the record"
	log.Println(m)
	makeGinResponse(c, http.StatusCreated, m)


}
func (env Env) UpdatePerson(c *gin.Context){
	var updatePersonData Person
	if err := c.BindJSON(&updatePersonData); err != nil {
		log.Printf("invalid JSON body: %v", err)
		makeGinResponse(c, http.StatusNotFound, err.Error())
		return
	}
	q:=`UPDATE Person SET personName=$1 ,address=$2 WHERE personId =$3`
	_,err:= env.DB.Exec(q,updatePersonData.PersonName,updatePersonData.PersonAddress,updatePersonData.PersonID)
	if err != nil {
		log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
		makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	m := "successfully updated the record"
	log.Println(m)
	makeGinResponse(c, http.StatusCreated, m)
}
func (env Env) DeletePerson(c *gin.Context){
	pId,err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		e := fmt.Sprintf("received invalid id path param which is not string: %v", c.Param("id"))
		log.Println(e)
		makeGinResponse(c, http.StatusNotFound, e)
		return
	}
	q:=`DELETE FROM Person WHERE personId =$1`
	_,err = env.DB.Exec(q,pId)
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
