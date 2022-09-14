package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type Student1 struct {
	ID       int  `json:"id"`
	Name     string `json:"name"`
	//Grade string `json:"grade"`
	Result     int  `json:"result"`
}
func main(){
	studentData := Student1{
		Name:"RAJA",
		Result: 50,
		ID:12}
	studentDataJson,_:= json.Marshal(studentData)


	url:= "http://localhost:8081/student"

	client := &http.Client{}

	req,err:= http.NewRequest("PUT",url,bytes.NewBuffer(studentDataJson))
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
