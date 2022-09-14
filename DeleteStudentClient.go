package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//studentData := Student{
	//	ID:12}
	//studentDataJson,_:= json.Marshal(studentData)
	url := "http://localhost:8081/student/12"


	client := &http.Client{}
	req,err:= http.NewRequest("DELETE",url,nil)
	//req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	respBody, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(respBody))
}
