package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	url := "http://localhost:8080/pop"
	client := &http.Client{}
	req,err:= http.NewRequest("DELETE",url,nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	respBody, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(respBody))
}


