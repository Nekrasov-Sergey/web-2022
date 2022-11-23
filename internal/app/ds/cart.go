package ds

import "github.com/google/uuid"

type Cart struct {
	Store    uuid.UUID `example:"976c088c-f218-422b-aff6-f9e1cf792860"`
	Quantity uint64    `example:"3"`
}

func (Cart) TableName() string {
	return "cart"
}
