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