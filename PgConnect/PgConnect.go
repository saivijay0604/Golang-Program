package PgConnect

import (
	"database/sql"
	"fmt"
	gin "github.com/gin-gonic/gin"
	"net/http"

	_ "github.com/lib/pq"
)
//type Student struct {
//	ID       int64  `json:"id"`
//	Name     string `json:"name"`
//	Grade string `json:"location"`
//	Re      int64  `json:"age"`
//}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Apr@1995"
	dbname   = "Training"
)
//global variables
var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
var db, err = sql.Open("postgres", psqlInfo)

//To check the DB connection
func CheckDB() {
	go CheckError(err)

}
func CheckError(err error){
	if err != nil {
		panic(err)
	}
}


func Connection(c *gin.Context){
	go CheckDB()
	defer db.Close()
	err = db.Ping()
	c.JSON(http.StatusOK,"connected")

}
 //To create the table
func Crud(c *gin.Context){
	go CheckDB()
	str:=c.Param("str")
	if str == "create" {
		table := `create table student (ID int, Name Varchar(100))`
		_, e := db.Exec(table)
		CheckError(e)
	}
	c.JSON(http.StatusOK, "Table Created")
	//To insert the data
	if str == "insert"{
		insertData:= `insert into student (id,"name") values (3,'venkat')`
		_, e := db.Exec(insertData)
		CheckError(e)

	}
	c.JSON(http.StatusOK,"Data inserted into the student table")

	// To update the data
		if str == "update"{
			updateData:= `update student set id=1,name='Vijay' where id =1`
			_, e := db.Exec(updateData)
			CheckError(e)

		}
	c.JSON(http.StatusOK,"Successfully updated")

		//To delete the data
	if str == "delete"{
		deleteData := `delete from student where id=$1`
		_, e := db.Exec(deleteData, 1)
		go CheckError(e)

	}
	c.JSON(http.StatusOK,"Successfully Deleted")
	defer db.Close()
	err = db.Ping()

}


