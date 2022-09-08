package PgConnect

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
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
func ConnectDB() (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &sql.DB{}, err
	}
	return db, nil
}

func makeGinResponse(c *gin.Context, statusCode int, value interface{}) {
	c.JSON(statusCode, gin.H{
		"response":   value,
		"statusCode": statusCode,
	})
}
