package model

type Item struct {
	ModelBase        // ID, CreatedAt, UpdatedAt, DeletedAt
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
}
