package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"./controllers"
)


func main() {
	// init router
	r := mux.NewRouter()

	// route handlers
	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8000", r))
}