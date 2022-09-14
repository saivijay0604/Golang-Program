package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Student struct {
	ID       int  `json:"id"`
	Name     string `json:"name"`
	Result     int  `json:"result"`
}
func main(){
	studentData:= Student{
		ID:9,
		Name:"RAJA",
		Result: 67}
	studentDataJson,err:=json.Marshal(studentData)
	if err!= nil{
		log.Fatal(err)
	}

	url:= "http://localhost:8081/student"
	//METHOD:= "POST"
	client := &http.Client{}

	req,err:= http.NewRequest("POST",url,bytes.NewBuffer(studentDataJson))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	respBody,readErr := ioutil.ReadAll(resp.Body)
	if readErr!= nil{
		log.Fatal(readErr)
	}
	fmt.Println(string(respBody))
}
