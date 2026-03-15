package queue

import (
	"fmt"
	"sync"
)

type TaskType string

const (
	TaskBookingConfirmation TaskType = "booking_confirmation"
	TaskEventUpdate         TaskType = "event_update"
)

type Task struct {
	Type    TaskType
	Payload interface{}
}

var (
	taskQueue = make(chan Task, 100)
	wg        sync.WaitGroup
)

func AddTask(task Task) {
	taskQueue <- task
}

func StartWorker() {
	for task := range taskQueue {
		switch task.Type {
		case TaskBookingConfirmation:
			payload := task.Payload.(BookingConfirmationPayload)
			fmt.Printf("[Email] Booking confirmation sent to user %d for event %d\n", payload.UserID, payload.EventID)
		case TaskEventUpdate:
			payload := task.Payload.(EventUpdatePayload)
			for _, userID := range payload.UserIDs {
				fmt.Printf("[Notification] Event %d updated. Notified user %d\n", payload.EventID, userID)
			}
		}
	}
}

type BookingConfirmationPayload struct {
	UserID  int
	EventID int
}

type EventUpdatePayload struct {
	EventID int
	UserIDs []int
}
