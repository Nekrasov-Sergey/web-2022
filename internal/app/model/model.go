package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Promos struct {
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Store    string
	Discount string
	Price    string
	Quantity uint64
	Promo    pq.StringArray `gorm:"type:text[]"`
}

type PromosDocs struct {
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Store    string
	Discount string
	Price    string
	Quantity uint64
	Promo    []string `gorm:"type:text[]"`
}
