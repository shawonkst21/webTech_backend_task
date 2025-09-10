package handlers

import (
	"encoding/json"
	"net/http"
	"taskManager/database"
	"taskManager/util"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&creds)

	var userID int
	var storedPassword ,role string
	err := database.DB.QueryRow("SELECT id,password,role FROM users WHERE email=?", creds.Email).Scan(&userID, &storedPassword,&role)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare hashed password with the password provided by user
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password))
	if err != nil {
		// Password does not match
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := util.GenerateJWT(userID,role)
	util.SendData(w, map[string]string{"token": token}, http.StatusOK)
}
