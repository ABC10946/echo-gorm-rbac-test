package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `json:"username"`
	Password   string `json:"password"`
}
