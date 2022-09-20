package ds

type Product struct {
	ID    uint `gorm:"primarykey"`
	Code  uint
	Price uint
}
