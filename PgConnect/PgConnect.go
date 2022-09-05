package PgConnect

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "strconv"

	_ "github.com/lib/pq"
)
type Student struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	//Grade string `json:"grade"`
	Result     int64  `json:"result"`
}

func NewStudent(id int64, name string, result int64) Student {
	return Student{id, name, result}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Apr@1995"
	dbname   = "Training"
)
//global variables
type Env struct{
	DB *sql.DB
}
var db *sql.DB
var err error


//To check the DB connection
func CheckDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db,err= sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		//makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()
	err = db.Ping()

}

func Connection(c *gin.Context){
	CheckDB()
	c.JSON(http.StatusOK,"connected")
}

//To create the table
func Create(c *gin.Context) {

	str := c.Param("str")
	if str == "create" {
		CheckDB()
		q := `create table student (id int,name VARCHAR ( 50 ),result int )`
		result, err := db.Exec(q)
		if err != nil {
			log.Printf("failed to connect to database: %v", err)
			makeGinResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		n, err := result.RowsAffected()
		if err != nil {
			e := fmt.Sprintf("error occurred while checking the returned result from database after deletion: %v", err)
			log.Println(e)
			makeGinResponse(c, http.StatusInternalServerError, e)
			return
		}

		// if no record was deleted, let us inform that there might be no
		// records to delete for this given album ID.
		if n == 0 {
			e := "Could not create the table"
			log.Println(e)
			makeGinResponse(c, http.StatusBadRequest, e)
			return
		}

		m := "successfully deleted the record"
		log.Println(m)
		makeGinResponse(c, http.StatusOK, m)
	}
}




func makeGinResponse(c *gin.Context, statusCode int, value string) {
	c.JSON(statusCode, gin.H{
		"message":    value,
		"statusCode": statusCode,
	})
}

