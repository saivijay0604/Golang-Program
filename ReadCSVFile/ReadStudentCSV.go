package ReadCSVFile

import (
	"ReadFile/example/ConnectDB"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type studentData struct{
	ID string
	Name string
	Result string
}
var env ConnectDB.Env
var err error

func ReadStudent(){
	env.DB,err =ConnectDB.ConnectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	defer env.DB.Close()


	csvFile,err := os.Open("C:\\Users\\saivi\\Downloads\\StudentDetails.csv")
	if err!=nil{
		log.Fatalln("Couldn't open the csv file", err)
	}
	//defer csvFile.Close()
	read :=csv.NewReader(csvFile)
	readAllStudents,_:=read.ReadAll()
	//fmt.Println(readAllStudents)
	var studentDetails = []studentData{}
	for _,value:= range readAllStudents{
		studentDetails = append(studentDetails,studentData{ID:value[0],Name:value[1],Result: value[2]})
	}

	for i:=1;i<len(studentDetails);i++ {
		sID,_:=strconv.Atoi(studentDetails[i].ID)
		sResult,_:=strconv.Atoi(studentDetails[i].Result)

		q := `INSERT INTO student(id,name,result) VALUES($1,$2,$3)`
		_, err := env.DB.Exec(q,sID,studentDetails[i].Name,sResult)
		if err != nil {
			log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
		}
	}
	fmt.Println(studentDetails)
	}
