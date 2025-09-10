package handlers

import (
	"encoding/json"
	"net/http"
	"taskManager/database"
	"taskManager/util"
)

// func CreateTask(w http.ResponseWriter, r *http.Request) {
	
// 	var newTask database.Task
// 	err := json.NewDecoder(r.Body).Decode(&newTask)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	result,err:=database.DB.Exec(
// 		`insert into tasks(title,description,status,user_id)
// 		values(?,?,?,?)
// 		`,
// 		newTask.Title,newTask.Description,newTask.Status,newTask.UserID,
// 	)
// 	if err!=nil{
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	id,_:=result.LastInsertId()
// 	newTask.ID=int(id)
// 	util.SendData(w, newTask, http.StatusCreated)
// }
func CreateTask(w http.ResponseWriter, r *http.Request) {
	
	claims := r.Context().Value("user").(*util.Claims)
	userID := claims.UserID

	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("INSERT INTO tasks (title, description, status, user_id) VALUES (?, ?, ?, ?)",
		t.Title, t.Description, t.Status, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	t.ID = int(id)
	t.UserID = userID

	util.SendData(w, t, http.StatusCreated)
}