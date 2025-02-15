package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-app/db"
	"todo-app/models"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, description, is_completed,created_at FROM tasks")
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt); err != nil {
			http.Error(w, "Ошибка при обработке данных", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}
	err := db.DB.QueryRow(
		"INSERT INTO  tasks (title,description) VALUES ($1,$2) RETURNING id, created_at", task.Title, task.Description).Scan(&task.ID, &task.CreatedAt)
	if err != nil {
		http.Error(w, "Ошибка при добавлении задачи", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Неккоректный ввод:", http.StatusBadRequest)
		return
	}
	result, err := db.DB.Exec("DELETE FROM tasks WHERE id =$1", id)
	if err != nil {
		fmt.Println("Ошибка удаления: ", err)
		http.Error(w, "Ошибка при удалении", http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Задача удалена"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Неккоректный id", http.StatusBadRequest)
		return
	}
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Неккоректный JSON", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("UPDATE  tasks SET title=$1, description=$2,is_completed=$3 WHERE id=$4", task.Title, task.Description, task.IsCompleted, id)
	if err != nil {
		http.Error(w, "Ошибка запроса", http.StatusBadRequest)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Задача обновлена"))

}
func GetTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Ошибка ввода", http.StatusBadRequest)
		return
	}
	err = db.DB.QueryRow("SELECT id,title,description,is_completed,created_at FROM tasks WHERE id=$1", id).Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Ошибка запроса", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
	w.WriteHeader(200)

}
