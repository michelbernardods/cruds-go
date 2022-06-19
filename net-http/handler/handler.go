package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Crud interface {}
type Todo struct {
	ID   int
	Task string
}

var todos []Todo
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {
		var todo Todo
		json.NewDecoder(r.Body).Decode(&todo)
		todos = append(todos, todo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Method not available"}`))
	}
}

func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		if len(todos) != 0 {
			json.NewEncoder(w).Encode(todos)
		}
		w.Write([]byte(`{"error": "Todo empty"}`))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Method not available"}`))
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	if r.Method == "PUT" {
		query := r.URL.Query()
		id, _ := strconv.Atoi(query.Get("id"))

		for index, todo := range todos {
			if todo.ID == id {
				json.NewDecoder(r.Body).Decode(&todo)
				todos[index].ID = todo.ID
				todos[index].Task = todo.Task
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message": "Success to update todo"}`))
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Method not available"}`))
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	if r.Method == "DELETE" {
		query := r.URL.Query()
		id, _ := strconv.Atoi(query.Get("id"))

		for index, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:index], todos[index+1:]...)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message": "Success to delete todo"}`))
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Method not available"}`))
	}
}
