package ds

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Store struct {
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string
	Discount uint64
	Price    uint64
	Quantity uint64
	Promo    pq.StringArray `gorm:"type:text[]"`
	Image    string
}

type StoreDocs struct {
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"-"`
	Name     string    `example:"Пятёрочка"`
	Discount uint64    `example:"400"`
	Price    uint64    `example:"200"`
	Quantity uint64    `example:"3"`
	Promo    []string  `example:"djzML,MdUI7,byP1f"`
	Image    string    `example:"https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/Five_gioiio.png"`
}

type QuantityStores struct {
	Quantity uint64 `example:"10"`
}

type PriceStore struct {
	Price uint64 `example:"300"`
}
