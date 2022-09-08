package main

import (
	_ "github.com/dbAPI/PgConnect"
	"github.com/dbAPI/Router"
)


func main(){
	Router.ConnectRouter()

}

