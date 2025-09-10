package handlers

import (
	"encoding/json"
	"net/http"
	"taskManager/database"
	"taskManager/util"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // optional, default to "user"
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Default role = user if empty
	if req.Role == "" {
		req.Role = "user"
	}

	// Insert user into DB
	_, err = database.DB.Exec(
		"INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)",
		req.Username, req.Email, string(hashedPassword), req.Role,
	)
	if err != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	util.SendData(w, map[string]string{"message": "User registered successfully"}, http.StatusCreated)
}
