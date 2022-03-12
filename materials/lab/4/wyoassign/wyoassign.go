package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Response struct{
	Classes []Class `json:"classes"`
}

type Class struct {
	Id string `json:"id"`
	Name string `json:"name`
	Instructor string `json:"instructor`
	Description string `json:"desc"`
	Number int `json:"number"`	
}

var Classes []Class
const Valkey string = "FooKey"

func InitClasses(){
	var clas Class
	clas.Id = "65A"
	clas.Name = "Physics "
	clas.Instructor = "Albert Einstein"
	clas.Description = "Learn Physics"
	clas.Number = 30	
	Classes = append(Classes, clas)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetClasses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Classes = Classes

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func GetClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, class := range Classes {
		if class.Id == params["id"]{
			json.NewEncoder(w).Encode(class)
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, class := range Classes {
			if class.Id == params["id"]{
				Classes = append(Classes[:index], Classes[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	
	var response Response
	var clas Class
	response.Classes = Classes
	params := mux.Vars(r)

	for index, class := range Classes {
		if class.Id == params["id"]{
			clas.Id =  r.FormValue("id")
			clas.Name =  r.FormValue("name")
			clas.Instructor = r.FormValue("instructor")
			clas.Description =  r.FormValue("desc")
			clas.Number, _ =  strconv.Atoi(r.FormValue("number"))
			Classes = append(Classes, clas)
			Classes = append(Classes[:index], Classes[index+1:]...)
		}
	}
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var clas Class
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		clas.Id =  r.FormValue("id")
		clas.Name =  r.FormValue("name")
		clas.Instructor = r.FormValue("instructor")
		clas.Description =  r.FormValue("desc")
		clas.Number, _ =  strconv.Atoi(r.FormValue("number"))
		Classes = append(Classes, clas)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)
}