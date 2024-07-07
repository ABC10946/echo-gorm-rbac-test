package model

type Group struct {
	ModelBase        // ID, CreatedAt, UpdatedAt, DeletedAt
	Name      string `json:"name"`
	User      []User `gorm:"many2many:user_groups;"`
}
