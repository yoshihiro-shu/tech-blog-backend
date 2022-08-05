package controller

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
