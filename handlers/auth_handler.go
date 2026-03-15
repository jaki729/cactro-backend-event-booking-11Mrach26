package handlers

import (
    "net/http"
    "BOOKING-BACKEND/service"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    service.HandleLogin(w, r)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    service.HandleRegister(w, r)
}