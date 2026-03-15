package repository

import (
    "BOOKING-BACKEND/models"
)

var bookings = []models.Booking{}
var bookingIDCounter = 1

func CreateBooking(booking models.Booking) models.Booking {
    booking.ID = bookingIDCounter
    bookingIDCounter++
    bookings = append(bookings, booking)
    return booking
}

func GetBookingsByUser(userID int) []models.Booking {
    result := []models.Booking{}
    for _, b := range bookings {
        if b.UserID == userID {
            result = append(result, b)
        }
    }
    return result
}

func GetBookingsByEvent(eventID int) []models.Booking {
    result := []models.Booking{}
    for _, b := range bookings {
        if b.EventID == eventID {
            result = append(result, b)
        }
    }
    return result
}