package service

import (
    "encoding/json"
    "net/http"
    "BOOKING-BACKEND/models"
    "BOOKING-BACKEND/repository"
    "BOOKING-BACKEND/queue"
)

func HandleBookingCreate(w http.ResponseWriter, r *http.Request) {
    user := Authenticate(r)
    if user == nil || user.Role != models.RoleCustomer {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    var req struct {
        EventID int `json:"event_id"`
    }
    json.NewDecoder(r.Body).Decode(&req)
    event := repository.GetEventByID(req.EventID)
    if event == nil {
        http.Error(w, "Event not found", http.StatusNotFound)
        return
    }
    booking := models.Booking{
        EventID: event.ID,
        UserID:  user.ID,
    }
    booking = repository.CreateBooking(booking)

    queue.AddTask(queue.Task{
        Type: queue.TaskBookingConfirmation,
        Payload: queue.BookingConfirmationPayload{
            UserID:  user.ID,
            EventID: event.ID,
        },
    })

    json.NewEncoder(w).Encode(booking)
}

func HandleBookingList(w http.ResponseWriter, r *http.Request) {
    user := Authenticate(r)
    if user == nil {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    bookings := repository.GetBookingsByUser(user.ID)
    json.NewEncoder(w).Encode(bookings)
}