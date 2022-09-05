package main

import (
	"github.com/dbAPI/Router"

	_ "github.com/dbAPI/PgConnect"
)

func main(){
	Router.ConnectRouter()

}

