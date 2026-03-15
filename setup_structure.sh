#!/bin/bash

# Set base directory
BASE_DIR="/home/jara/Documents/Personal-Docs/booking-backend"
cd "$BASE_DIR"

# Create directories
mkdir -p models handlers middleware queue repository service

# Create files
touch main.go
touch models/user.go models/event.go models/booking.go
touch handlers/auth_handler.go handlers/event_handler.go handlers/booking_handler.go
touch middleware/auth.go
touch queue/worker.go
touch repository/user_repo.go repository/event_repo.go repository/booking_repo.go
touch service/auth_service.go service/event_service.go service/booking_service.go

echo "Project structure created successfully."
