package main

import (

	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type Student struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Class string  `json:"class"`
	Address string `json:"address"`
}

var students []Student
func main(){
	router:=mux.NewRouter()

	students =append(students ,Student{ID:1,Name:"vaibhav",Class:"tybsccs",Address:"abc"},
	Student{ID:2,Name:"Raj",Class:"tybscit",Address:"adc"},
	Student{ID:3,Name:"prince",Class:"tybsccs",Address:"xyz"},
	Student{ID:4,Name:"rahul",Class:"tyb.com",Address:"hsj"})



	router.HandleFunc("/students", getStudents).Methods("GET")
	router.HandleFunc("/students/{id}", getStudent).Methods("GET")
	router.HandleFunc("/students", addStudent).Methods("POST")
	router.HandleFunc("/students", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", removeStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":7002",router))
}

func getStudents(w http.ResponseWriter,r *http.Request){

json.NewEncoder(w).Encode(students)

}

func getStudent(w http.ResponseWriter,r *http.Request){

	params :=mux.Vars(r)

i,_:=strconv.Atoi(params["id"])

for _,student:=range students {
	
	if student.ID == i {
		
		json.NewEncoder(w).Encode(&student)
	}
}
}

func addStudent(w http.ResponseWriter,r *http.Request){

	var student Student
	
	_ =json.NewDecoder(r.Body).Decode(&students)
	
	students = append(students,student)
	
	json.NewEncoder(w).Encode(students)


}

func updateStudent(w http.ResponseWriter,r *http.Request){

	var student Student
json.NewDecoder(r.Body).Decode(&student)

for i,item:=range students{
	if item.ID == student.ID{
	students[i] = student
	
}
}

json.NewEncoder(w).Encode(students)

}

func removeStudent(w http.ResponseWriter,r *http.Request){

	params :=mux.Vars(r)
	
	i,_:=strconv.Atoi(params["id"])
	
	for _,student:=range students {
		if student.ID == i {
			students =append(students[:i],students[i+1:]...)
}
	}
	json.NewEncoder(w).Encode(students)
}