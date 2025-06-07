package model

import (
    "time"

    "github.com/google/uuid"
)

type Users struct {
    Id        uuid.UUID `json:"id,omitempty" db:"id"` 
    FirstName string    `json:"fname,omitempty" db:"fname"`
    LastName  string    `json:"lname,omitempty" db:"lname"`
    Email     string    `json:"email,omitempty" db:"email"`
    Password  string    `json:"password,omitempty" db:"password"`
    CreatedAt time.Time `json:"createdAt,omitempty" db:"createdAt"`
}
