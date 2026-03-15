package service

import (
	"BOOKING-BACKEND/models"
	"BOOKING-BACKEND/repository"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	user := repository.GetUserByUsername(req.Username)
	if user == nil || user.Password != req.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string          `json:"username"`
		Password string          `json:"password"`
		Role     models.UserRole `json:"role"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}
	user = repository.CreateUser(user)
	log.Printf("[Register] User: %s Role: %s", user.Username, user.Role)
	json.NewEncoder(w).Encode(user)
}

func Authenticate(r *http.Request) *models.User {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Basic ") {
		return nil
	}
	payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
	if err != nil {
		return nil
	}
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return nil
	}
	username, password := pair[0], pair[1]
	user := repository.GetUserByUsername(username)
	if user == nil || user.Password != password {
		return nil
	}
	return user
}
