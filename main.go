package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/db"
	"todo-app/handlers"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	db.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id:[0-9]+}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id:[0-9]+}", handlers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id:[0-9]+}", handlers.UpdateTask).Methods("PUT")
	log.Println("Cервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	defer db.DB.Close()
	query := `CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        is_completed BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

	_, err := db.DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Таблица успешно создана!")

}
