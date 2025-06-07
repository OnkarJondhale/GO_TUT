package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id        uuid.UUID `json:"id,omitempty" db:"id"`
	UserId    uuid.UUID `json:"userId,omitempty" db:"userId"`
	PostId    uuid.UUID `json:"postId,omitempty" db:"postId"`
	Content   string    `json:"content,omitempty" db:"content"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"createdAt"`
}
