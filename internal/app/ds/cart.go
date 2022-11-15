package ds

import "github.com/google/uuid"

type Cart struct {
	Store    uuid.UUID
	Quantity uint64
}

func (Cart) TableName() string {
	return "cart"
}
