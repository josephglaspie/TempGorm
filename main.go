package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HelloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeletUser).Methods("DELETE")
	myRouter.HandleFunc("user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("GORM tutorial")
	InitialMigration()

	handleRequests()
}
