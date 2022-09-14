package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	url:= "http://localhost:8081/student/1"
	req,err:= http.Get(url)
	//resp,err= http.
	if err!= nil{
		log.Fatal(err)
	}
	resp,readErr := ioutil.ReadAll(req.Body)
	if readErr!= nil{
		log.Fatal(readErr)
	}
	fmt.Println(string(resp))
}
