package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
}
