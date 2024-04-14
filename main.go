package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	initDB()
	r := mux.NewRouter()
	corsRouter := enableCORS(r)

	r.HandleFunc("/todos", logHandler(getTodos)).Methods("GET")
	r.HandleFunc("/todo/{id}", logHandler(getTodo)).Methods("GET")
	r.HandleFunc("/todo", logHandler(createTodo)).Methods("POST")
	r.HandleFunc("/todo/{id}", logHandler(deleteTodo)).Methods("DELETE")

	http.ListenAndServe(":8080", corsRouter)
}
