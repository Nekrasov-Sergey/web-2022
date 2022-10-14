package ds

import "github.com/google/uuid"

type Promos struct {
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Store    string
	Discount string
	Price    string
	Quantity uint64
	Promo    []byte `gorm:"type:json"`
}
