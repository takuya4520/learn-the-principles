package main

import (
	"html/template"
	"net/http"
)

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
	tmpl, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		http.Error(w, "faild to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, tasks)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	tasks = append(tasks, title)
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}
