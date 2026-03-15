package repository

import (
    "BOOKING-BACKEND/models"
)

var events = []models.Event{}
var eventIDCounter = 1

func CreateEvent(event models.Event) models.Event {
    event.ID = eventIDCounter
    eventIDCounter++
    events = append(events, event)
    return event
}

func GetEvents() []models.Event {
    return events
}

func GetEventByID(id int) *models.Event {
    for _, e := range events {
        if e.ID == id {
            return &e
        }
    }
    return nil
}

func UpdateEvent(event models.Event) bool {
    for i, e := range events {
        if e.ID == event.ID {
            events[i] = event
            return true
        }
    }
    return false
}

func DeleteEvent(id int) bool {
    for i, e := range events {
        if e.ID == id {
            events = append(events[:i], events[i+1:]...)
            return true
        }
    }
    return false
}