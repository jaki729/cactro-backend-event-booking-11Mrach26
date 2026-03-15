package repository

import (
    "BOOKING-BACKEND/models"
)

var users = []models.User{}
var userIDCounter = 1

func CreateUser(user models.User) models.User {
    user.ID = userIDCounter
    userIDCounter++
    users = append(users, user)
    return user
}

func GetUserByUsername(username string) *models.User {
    for _, u := range users {
        if u.Username == username {
            return &u
        }
    }
    return nil
}

func GetUserByID(id int) *models.User {
    for _, u := range users {
        if u.ID == id {
            return &u
        }
    }
    return nil
}