package main

import (
	"fmt"
	"log"
	"net/http"

	"github .com/gorilla/mux"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",helloworld).Methods("GET")
	myRouter.HandleFunc("/users",AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}",DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}",UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081",myRouter))
}

func main() {
    fmt.Println("Go ORM Tutorial")

    // Add the call to our new initialMigration function
    initialMigration()

    handleRequests()
}

func main()
	fmt.Println("Go ORM Tutorial")
	handleRequests()
}	