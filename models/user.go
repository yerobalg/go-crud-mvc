package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null;unique" json:"username"`
	Password string `gorm:"type:text;" json:"-"`
}

type UserParam struct {
	UserID int64 `uri:"id"`
	PaginationParam
}

type UserBodyParam struct {
	Username        string `json:"username"`
	CurrentPassword string `json:"password,omitempty"`
	NewPassword     string `json:"new_password,omitempty"`
}
