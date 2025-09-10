package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	var updatedTask map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if len(updatedTask)==0{
		http.Error(w, "No fields provided to update", http.StatusBadRequest)
		return
	}
	setparts:=[]string{}
	args:=[]interface{}{}

	for col,val:=range updatedTask{
		setparts=append(setparts, fmt.Sprintf("%s=?",col))
		args=append(args, val)
	}

	query:=fmt.Sprintf(`update tasks set %s where id=?`,strings.Join(setparts,","))
	args = append(args, id)

	result,er:=database.DB.Exec(query,args...)

	if er != nil {
		http.Error(w, "Database update error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	util.SendData(w,  map[string]string{"message": "Task updated successfully"}, http.StatusOK)
}
