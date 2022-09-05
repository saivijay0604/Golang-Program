package main

import (
	"github.com/dbAPI/PgConnect"
	_ "github.com/dbAPI/PgConnect"
	"github.com/dbAPI/Router"
)
var env = new(PgConnect.Env)


func main(){


	defer env.DB.Close()
	Router.ConnectRouter()

}

