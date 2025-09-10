package handlers

import (
	"net/http"
	"strconv"
	"taskManager/database"
	"taskManager/util"
)
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "please me a valid id", 400)
		return
	}
	index := -1
	for i, t := range database.TaskList {
		if t.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	database.TaskList = append(database.TaskList[:index], database.TaskList[index+1:]...)
	util.SendData(w, map[string]string{"message": "Task deleted"}, http.StatusOK)
}
