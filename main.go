package main

import (
	"log"
	"net/http"

	"BOOKING-BACKEND/handlers"
	"BOOKING-BACKEND/middleware"
	"BOOKING-BACKEND/queue"
)

func main() {
	// Start background worker
	go queue.StartWorker()

	mux := http.NewServeMux()

	// Auth routes
	mux.HandleFunc("/login", middleware.LoggerMiddleware(handlers.LoginHandler))
	
	// Event routes
	mux.HandleFunc("/events", middleware.LoggerMiddleware(middleware.AuthMiddleware(handlers.EventListHandler)))
	mux.HandleFunc("/events/create", middleware.LoggerMiddleware(middleware.AuthMiddleware(handlers.EventCreateHandler)))
	mux.HandleFunc("/events/update", middleware.LoggerMiddleware(middleware.AuthMiddleware(handlers.EventUpdateHandler)))
	mux.HandleFunc("/events/delete", middleware.LoggerMiddleware(middleware.AuthMiddleware(handlers.EventDeleteHandler)))

	// Booking routes
	mux.HandleFunc("/bookings/create", middleware.LoggerMiddleware(middleware.AuthMiddleware(handlers.BookingCreateHandler)))
	mux.HandleFunc("/bookings/my", middleware.LoggerMiddleware(middleware.AuthMiddleware(handlers.BookingListHandler)))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
