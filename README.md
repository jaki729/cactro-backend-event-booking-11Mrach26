# Booking Backend

A simple event booking backend system in Go supporting organizers and customers, with role-based access and background tasks.

## Features

- **User Roles:** Organizer and Customer
- **Organizer:** Create, update, delete events
- **Customer:** Browse events, book tickets
- **Background Tasks:**
  - Booking confirmation (simulated email via console log)
  - Event update notification (simulated notification via console log)
- **Basic Authentication:** All protected endpoints require Basic Auth

## Project Structure

```
booking-backend/
├── main.go
├── go.mod
├── README.md
├── curls.md
├── handlers/
│   ├── auth_handler.go
│   ├── booking_handler.go
│   └── event_handler.go
├── middleware/
│   └── auth.go
├── models/
│   ├── booking.go
│   ├── event.go
│   └── user.go
├── queue/
│   └── worker.go
├── repository/
│   ├── booking_repo.go
│   ├── event_repo.go
│   └── user_repo.go
├── service/
│   ├── auth_service.go
│   ├── booking_service.go
│   └── event_service.go
```

## Setup

1. **Initialize Go module (if needed):**
   ```sh
   go mod init BOOKING-BACKEND
   go mod tidy
   ```

2. **Run the server:**
   ```sh
   go run main.go
   ```
   The server starts at `http://localhost:8080`.

## API Usage

See [curls.md](curls.md) for example curl commands for all endpoints.

## Logging

- All API requests and important operations are logged to the console.
- Booking confirmation and event update notifications are logged as background tasks.

## Notes

- All data is stored in memory; restart will reset users/events/bookings.
- Use Basic Auth for all protected endpoints.

## Background Job Queue and Notifications

- **Queue Implementation:** The system uses a buffered channel (`taskQueue`) with a capacity of 100 tasks to handle asynchronous jobs, preventing blocking on high load.
- **Worker Goroutine:** A dedicated goroutine runs the [`StartWorker`](queue/worker.go) function, processing tasks in a loop without blocking the main server thread.
- **Supported Tasks:**
  - **Booking Confirmation:** Simulates email sending by logging a confirmation message to the console when a booking is created.
  - **Event Update Notification:** Simulates notifications by logging updates to all affected users when an event is modified or deleted.
- **Task Addition:** Tasks are added via [`AddTask`](queue/worker.go), ensuring non-blocking enqueueing.
- **Concurrency:** The queue is thread-safe using Go's channel semantics; no additional locks are needed.
- **Limitations:** Notifications are console-based simulations; in production, integrate with email services (e.g., SMTP) or push notification systems.