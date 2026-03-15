package models

type UserRole string

const (
    RoleOrganizer UserRole = "organizer"
    RoleCustomer  UserRole = "customer"
)

type User struct {
    ID       int
    Username string
    Password string
    Role     UserRole
}