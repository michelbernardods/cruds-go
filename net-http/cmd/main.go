package main

import (
	"net/http"

	"github.com/michelbernardods/rest-net-http/handler"
)

func main() {
	//handlers
	healthHandler := &handler.Health{}
	createHandler := handler.Create
	readHandler := handler.Read
	updateHandler := handler.Update
	deleteHandler := handler.Delete

	//routers
	http.HandleFunc("/health", healthHandler.Handle().ServeHTTP)
	http.HandleFunc("/create/todo", createHandler)
	http.HandleFunc("/read/todo", readHandler)
	http.HandleFunc("/update/todo", updateHandler)
	http.HandleFunc("/delete/todo", deleteHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
