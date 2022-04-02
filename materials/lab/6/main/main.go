package main

// main.go HAS FOUR TODOS - TODO_1 - TODO_4

import (	
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"scrape/scrape"
)

//TODO_1: Logging right now just happens, create a global constant integer LOG_LEVEL 
//TODO_1: When LOG_LEVEL = 0 DO NOT LOG anything
//TODO_1: When LOG_LEVEL = 1 LOG API details only 
//TODO_1: When LOG_LEVEL = 2 LOG API details and file matches (e.g., everything)

const LOG_LEVEL = 2

func main() {
	if LOG_LEVEL == 2{	
		log.Println("starting API server")
		//create a new router
		router := mux.NewRouter()
		log.Println("creating routes")
		//specify endpoints
		router.HandleFunc("/", scrape.MainPage).Methods("GET")

		router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

		router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
		router.HandleFunc("/search", scrape.FindFile).Methods("GET")		
		router.HandleFunc("/addsearch/{regex}", scrape.ADDREG).Methods("GET")
		router.HandleFunc("/clear", scrape.CLEARARRAY).Methods("GET")
		router.HandleFunc("/reset", scrape.RESETARRAY).Methods("GET")
		http.Handle("/", router)

		//start and listen to requests
		http.ListenAndServe(":8080", router)
		//} else if LOG_LEVEL == 2{
			//log.Println("starting API server")			
	} else if LOG_LEVEL ==1{
		log.Println("starting API server")	
		router := mux.NewRouter()
		log.Println("creating routes")
		//specify endpoints
		router.HandleFunc("/", scrape.MainPage).Methods("GET")

		router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

		router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
		router.HandleFunc("/search", scrape.FindFile).Methods("GET")	
		http.Handle("/", router)
		//start and listen to requests
		http.ListenAndServe(":8080", router)
	} else if LOG_LEVEL ==0{
		fmt.Println("No logging")		
		}
}