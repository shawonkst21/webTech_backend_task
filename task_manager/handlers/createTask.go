package handlers

import (
	"encoding/json"
	"net/http"
	"taskManager/database"
	"taskManager/util"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	
	var newTask database.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newTask.ID = len(database.TaskList) + 1
	database.TaskList = append(database.TaskList, newTask)

	util.SendData(w, newTask, http.StatusCreated)
}
