package handlers

import (
    "net/http"
    "BOOKING-BACKEND/service"
)

func EventListHandler(w http.ResponseWriter, r *http.Request) {
    service.HandleEventList(w, r)
}

func EventCreateHandler(w http.ResponseWriter, r *http.Request) {
    service.HandleEventCreate(w, r)
}

func EventUpdateHandler(w http.ResponseWriter, r *http.Request) {
    service.HandleEventUpdate(w, r)
}

func EventDeleteHandler(w http.ResponseWriter, r *http.Request) {
    service.HandleEventDelete(w, r)
}