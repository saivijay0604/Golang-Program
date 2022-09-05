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
	ID       int  `json:"id"`
	Name     string `json:"name"`
	//Grade string `json:"grade"`
	Result     int  `json:"result"`
}


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Apr@1995"
	dbname   = "Training"
)

type Env struct{
	DB *sql.DB
}

var err error

var env = new(Env)
//To check the DB connection
func CheckDB() (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &sql.DB{}, err
	}
	return db, nil
}

func Connection(c *gin.Context){
	env.DB,err = CheckDB()
	c.JSON(http.StatusOK,"connected")
}

func makeGinResponse(c *gin.Context, statusCode int, value string) {
	c.JSON(statusCode, gin.H{
		"message":    value,
		"statusCode": statusCode,
	})
}
