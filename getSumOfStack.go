package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	url := "http://localhost:8080/sum"
	req,err:= http.Get(url)
	if err!= nil{
		fmt.Println(err)
	}
	resp,readErr:= ioutil.ReadAll(req.Body)
	if readErr!=nil{
		fmt.Println(readErr)
	}
	fmt.Println(string(resp))
}
