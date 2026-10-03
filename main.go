package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var tasks []string

var taskTemplate = template.Must(
	template.New("task").Parse("<li>{{.}}</li>"),
)

func main() {
  db, err := sql.Open(
    "mysql",
    "root@tcp(127.0.0.1:3306)/learn_the_principles",
  )
  if err != nil{
    panic(err)
  }
  defer db.Close()

  if err := db.Ping(); err != nil {
    panic(err)
  }

  println("connected to MySQL")

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

	taskTemplate.Execute(w, title)
}
