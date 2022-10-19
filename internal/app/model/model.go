package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Promos struct {
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Store    string
	Discount uint64
	Price    uint64
	Quantity uint64
	Promo    pq.StringArray `gorm:"type:text[]"`
}

type PromosDocs struct {
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Store    string
	Discount uint64
	Price    uint64
	Quantity uint64
	Promo    []string `gorm:"type:text[]"`
}

type Amount struct {
	Amount int
}
