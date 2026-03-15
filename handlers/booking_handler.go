package handlers

import (
    "net/http"
    "BOOKING-BACKEND/service"
)

func BookingCreateHandler(w http.ResponseWriter, r *http.Request) {
    service.HandleBookingCreate(w, r)
}

func BookingListHandler(w http.ResponseWriter, r *http.Request) {
    service.HandleBookingList(w, r)
}