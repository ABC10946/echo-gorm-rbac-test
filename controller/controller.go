package controller

import (
	"http-db-exp/model"
	"sync"

	"gorm.io/gorm"
)

type Controller struct {
	DB    *gorm.DB
	Items []model.Item
	Mutex sync.Mutex
}
