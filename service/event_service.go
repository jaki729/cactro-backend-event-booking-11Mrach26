package service

import (
    "encoding/json"
    "net/http"
    "BOOKING-BACKEND/models"
    "BOOKING-BACKEND/repository"
    "BOOKING-BACKEND/queue"
)

func HandleEventList(w http.ResponseWriter, r *http.Request) {
    events := repository.GetEvents()
    json.NewEncoder(w).Encode(events)
}

func HandleEventCreate(w http.ResponseWriter, r *http.Request) {
    user := Authenticate(r)
    if user == nil || user.Role != models.RoleOrganizer {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    var req struct {
        Name        string `json:"name"`
        Description string `json:"description"`
    }
    json.NewDecoder(r.Body).Decode(&req)
    event := models.Event{
        Name:        req.Name,
        Description: req.Description,
        OrganizerID: user.ID,
    }
    event = repository.CreateEvent(event)
    json.NewEncoder(w).Encode(event)
}

func HandleEventUpdate(w http.ResponseWriter, r *http.Request) {
    user := Authenticate(r)
    if user == nil || user.Role != models.RoleOrganizer {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    var req struct {
        ID          int    `json:"id"`
        Name        string `json:"name"`
        Description string `json:"description"`
    }
    json.NewDecoder(r.Body).Decode(&req)
    event := repository.GetEventByID(req.ID)
    if event == nil || event.OrganizerID != user.ID {
        http.Error(w, "Not found or forbidden", http.StatusNotFound)
        return
    }
    event.Name = req.Name
    event.Description = req.Description
    repository.UpdateEvent(*event)

    bookings := repository.GetBookingsByEvent(event.ID)
    userIDs := []int{}
    for _, b := range bookings {
        userIDs = append(userIDs, b.UserID)
    }
    queue.AddTask(queue.Task{
        Type: queue.TaskEventUpdate,
        Payload: queue.EventUpdatePayload{
            EventID: event.ID,
            UserIDs: userIDs,
        },
    })

    json.NewEncoder(w).Encode(event)
}

func HandleEventDelete(w http.ResponseWriter, r *http.Request) {
    user := Authenticate(r)
    if user == nil || user.Role != models.RoleOrganizer {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    var req struct {
        ID int `json:"id"`
    }
    json.NewDecoder(r.Body).Decode(&req)
    event := repository.GetEventByID(req.ID)
    if event == nil || event.OrganizerID != user.ID {
        http.Error(w, "Not found or forbidden", http.StatusNotFound)
        return
    }
    repository.DeleteEvent(event.ID)
    w.WriteHeader(http.StatusNoContent)
}