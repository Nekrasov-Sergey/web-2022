package ds

type Promos struct {
	ID    uint `gorm:"primarykey"`
	Code  uint
	Store string
	Promo string
	Price uint
}
