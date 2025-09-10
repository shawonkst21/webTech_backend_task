package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"taskManager/database"
	"taskManager/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var userID int
	var dbToken string
	var expiry sql.NullTime

	// Fetch user with token + expiry
	err := database.DB.QueryRow(
		"SELECT id, reset_token, reset_token_expiry FROM users WHERE reset_token=?",
		req.Token,
	).Scan(&userID, &dbToken, &expiry)

	fmt.Println("Incoming token:", req.Token)
	fmt.Println("DB token:", dbToken)
	fmt.Println("Expiry valid:", expiry.Valid)
	if expiry.Valid {
		fmt.Println("Expiry time:", expiry.Time)
	}
	fmt.Println("Now:", time.Now().UTC())

	if err != nil || dbToken == "" || !expiry.Valid {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	// Expired?
	if time.Now().UTC().After(expiry.Time.UTC()) {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Update password and clear reset token
	_, err = database.DB.Exec(
		"UPDATE users SET password=?, reset_token=NULL, reset_token_expiry=NULL WHERE id=?",
		string(hashedPassword), userID,
	)
	if err != nil {
		http.Error(w, "Unable to reset password", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{"message": "Password reset successfully"}, http.StatusOK)
}
