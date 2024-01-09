package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name   string
	Email  string
	Phone  string
	Client bool
	Age    int64
}

type LeadRespose struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Client    bool      `json:"client"`
	Age       int64     `json:"salary"`
}
