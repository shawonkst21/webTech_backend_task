package handlers

import (
	"database/sql"
	"net/http"
	"taskManager/database"
	"taskManager/util"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      int    `json:"user_id"`
}

var TaskList []Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("user").(*util.Claims)
	userID := claims.UserID
	role := claims.Role

	var rows *sql.Rows
	var err error
	query := "SELECT id, title, description, status, user_id FROM tasks"
	if role != "admin" {
		query += " WHERE user_id=?"
		rows, err = database.DB.Query(query, userID)
	} else {
		rows, err = database.DB.Query(query)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	util.SendData(w, tasks, http.StatusOK)
}

