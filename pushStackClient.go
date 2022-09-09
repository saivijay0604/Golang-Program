package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type structNum struct{
	num int `json:"num"`
}
func main(){
	stackNum :=structNum{
		num:1,
	}

	stackData,_:= json.Marshal(stackNum.num)
	url:= "http://localhost:8080/push"

	client := &http.Client{}

	req,err:= http.NewRequest("POST",url,bytes.NewBuffer(stackData))
	req.Header.Set("Content-Type", "application/json")

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
