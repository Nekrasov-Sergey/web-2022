package ds

type Promos struct {
	UUID     uint `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"uuid"`
	Code     uint
	Store    string
	Promo    string
	Price    string
	Quantity uint
}
