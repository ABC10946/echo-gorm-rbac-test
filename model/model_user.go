package model

type User struct {
	ModelBase         // ID, CreatedAt, UpdatedAt, DeletedAt
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Group     []Group `gorm:"many2many:user_groups;"`
}
