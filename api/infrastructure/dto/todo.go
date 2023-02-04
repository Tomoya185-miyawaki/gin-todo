package dto

import (
	"time"
)

type Todo struct {
	ID        uint `json:"id" binding:"required"`
	Title     string
	CreateAt  time.Time
	UpdatedAt time.Time
}
