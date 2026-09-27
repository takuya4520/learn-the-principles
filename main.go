package main

import "net/http"

var tasks []string

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
	for _, task := range tasks {
		w.Write([]byte(task + "\n"))
	}
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, 100)
	n, _ := r.Body.Read(buf)

	title := string(buf[:n])
	println(title)
	tasks = append(tasks, title)

	w.Write([]byte("Task created"))
}
