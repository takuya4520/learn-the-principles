package main

import "net/http"

func main() {
	http.HandleFunc("GET /{$}", homeHandler)
	http.HandleFunc("GET /tasks", tasksHandler)
	http.HandleFunc("POST /tasks", createTaskHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	println(r.Method)
	w.Write([]byte("Welcome to Learn the Principles!"))
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tasks"))
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Task created"))
}
