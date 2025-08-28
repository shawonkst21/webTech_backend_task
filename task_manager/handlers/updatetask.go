package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"taskManager/database"
	"taskManager/util"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "please me a valid id", 400)
		return
	}

	var updatedTask database.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	found := false
	for i, t := range database.TaskList {
		if t.ID == id {
			updatedTask.ID = t.ID
			database.TaskList[i] = updatedTask
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	util.SendData(w, updatedTask, http.StatusOK)
}
