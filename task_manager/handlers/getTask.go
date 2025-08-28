package handlers

import (
	"net/http"
	"taskManager/database"
	"taskManager/util"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.TaskList, http.StatusOK)
}
