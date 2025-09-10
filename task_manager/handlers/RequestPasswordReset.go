package handlers

import (
	"encoding/json"
	"net/http"
	"taskManager/database"
	"taskManager/util"
	"time"
)

func RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Email string `json:"email"`
	}
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	// Check if email exists
	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE email=?", req.Email).Scan(&userID)
	if err != nil {
		http.Error(w, "If email exists, a reset link will be sent", http.StatusOK)
		return
	}

	// Generate token and expiry
	token := util.GenerateRandomString(32)
	expiry := time.Now().Add(1 * time.Hour)

	_, err = database.DB.Exec("UPDATE users SET reset_token=?, reset_token_expiry=? WHERE id=?", token, expiry, userID)
	if err != nil {
		http.Error(w, "Unable to generate reset link", http.StatusInternalServerError)
		return
	}

	// Send email with token
	util.SendResetEmail(req.Email, token)

	util.SendData(w, map[string]string{"message": "If email exists, a reset link will be sent"}, http.StatusOK)
}
