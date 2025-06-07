package model

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	Id        uuid.UUID `json:"id,omitempty" db:"id"`
	Title     string    `json:"title,omitempty" db:"title"`
	Content   string    `json:"content,omitempty" db:"content"`
	UserId    uuid.UUID `json:"userId,omitempty" db:"userId"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"createdAt"`
}
