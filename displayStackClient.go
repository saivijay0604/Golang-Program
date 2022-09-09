package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){

	url := "http://localhost:8080/display"
	req,err:=http.Get(url)
	if err!= nil{
		fmt.Println(err)
		return
	}
	resp,readErr:= ioutil.ReadAll(req.Body)
	if readErr!= nil{
		fmt.Println(readErr)
		return
	}
	fmt.Println(string(resp))

}
