package main

import (
	"github.com/michelbernardods/rest-net-http/handler"
	"net/http"
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
	http.HandleFunc("/Create/todo", createHandler)
	http.HandleFunc("/Read/todo", readHandler)
	http.HandleFunc("/Update/todo", updateHandler)
	http.HandleFunc("/Delete/todo", deleteHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}