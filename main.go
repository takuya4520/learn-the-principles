package main

import (
	"database/sql"
	"html/template"
	"net/http"
  "os"

	_ "github.com/go-sql-driver/mysql"
)

var taskTemplate = template.Must(
	template.New("task").Parse("<li>{{.}}</li>"),
)

func main() {
  dsn := os.Getenv("DATABASE_URL")
  db, err := sql.Open("mysql", dsn)
	
  if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	println("connected to MySQL")

	http.HandleFunc("GET /{$}", homeHandler)
	http.HandleFunc("GET /tasks", tasksHandler(db))
	http.HandleFunc("POST /tasks", createTaskHandler(db))
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	println(r.Method)
	w.Write([]byte("Welcome to Learn the Principles!"))
}

func tasksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tasks []string

		rows, err := db.Query("SELECT title FROM tasks ORDER BY id")

		if err != nil {
			http.Error(w, "failed to query tasks", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var title string

			if err := rows.Scan(&title); err != nil {
				http.Error(w, "failed to scan task", http.StatusInternalServerError)
				return
			}

			tasks = append(tasks, title)
		}
		tmpl, err := template.ParseFiles("templates/tasks.html")
		if err != nil {
			http.Error(w, "faild to load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, tasks)
	}
}

func createTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")

		_, err := db.Exec(
			"INSERT INTO tasks (title, created_at, updated_at) VALUES (?, NOW(), NOW())",
			title,
		)
		if err != nil {
			http.Error(w, "failed to create task", http.StatusInternalServerError)
			return
		}
		taskTemplate.Execute(w, title)
	}
}
