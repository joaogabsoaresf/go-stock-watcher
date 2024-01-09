package schemas

import (
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
