package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/internal/app/ds"
	"main/internal/app/dsn"
	"math/rand"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func New() (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetStores() ([]ds.Store, error) {
	var stores []ds.Store
	err := r.db.Order("uuid").Find(&stores).Error
	return stores, err
}

func (r *Repository) GetPriceStore(uuid uuid.UUID, promo *ds.Store) (uint64, error) {
	err := r.db.First(&promo, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	return 0, nil
}

func (r *Repository) GetPromoStore(uuid uuid.UUID) (int, string, error) {
	var promo ds.Store
	err := r.db.First(&promo, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, "", err
		}
		return 500, "", err
	}

	Promo := promo.Promo[0]

	if promo.Quantity > 0 {
		err = r.db.Model(&promo).Update("Promo", promo.Promo[1:]).Error
		if err != nil {
			return 500, "", err
		}
		err = r.db.Model(&promo).Update("Quantity", promo.Quantity-1).Error
		if err != nil {
			return 500, "", err
		}
	}

	if promo.Quantity == 0 {
		err = r.db.Delete(&promo, uuid).Error
		if err != nil {
			return 500, "", err
		}
	}

	return 0, Promo, nil
}

func (r *Repository) CreateStore(promo ds.Store) error {
	err := r.db.Create(&promo).Error
	return err
}

func RandPromo(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var image = map[string]string{
	"Пятёрочка":   "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/Five_gioiio.png",
	"Магнит":      "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665905/Promos/Magnit_mtz50g.png",
	"Лента":       "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665905/Promos/Lenta_ocur5v.png",
	"ВиТ":         "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/ViT_f2xubg.png",
	"ДОДО":        "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665905/Promos/DODO_kkmdol.webp",
	"Яндекс Плюс": "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/YandexPlus_cch1ec.jpg",
	"Lamoda":      "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665905/Promos/Lamoda_xor4fl.jpg",
	"OZON":        "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/OZON_w2no08.png",
	"Wildberries": "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/WB_ri56ng.png",
}

func (r *Repository) CreateRandomStores() error {
	rand.Seed(time.Now().UnixNano())

	price := (rand.Intn(99) + 1) * 10

	discount := price * 2

	nameList := []string{"Пятёрочка", "Магнит", "Лента", "ВиТ", "ДОДО", "Яндекс Плюс", "Lamoda", "OZON", "Wildberries"}
	nameRandom := rand.Intn(len(nameList))
	name := nameList[nameRandom]
	quantity := rand.Intn(5) + 1

	var promoList []string
	for i := 0; i < quantity; i++ {
		promoList = append(promoList, RandPromo(5))
	}

	newStore := ds.Store{
		Name:     name,             // магазин
		Discount: uint64(discount), // скидка
		Price:    uint64(price),    // цена от 10 до 999
		Quantity: uint64(quantity), // оставшееся кол-во от 0 до 4
		Promo:    promoList,        // слайс промокодов
		Image:    image[name],
	}
	err := r.db.Create(&newStore).Error
	return err
}

func (r *Repository) ChangePriceStore(uuid uuid.UUID, price uint64) (int, error) {
	var promo ds.Store
	err := r.db.First(&promo, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Model(&promo).Update("Price", price).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}

func (r *Repository) DeleteStore(uuid uuid.UUID) (int, error) {
	var promo ds.Store
	err := r.db.First(&promo, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Delete(&promo, uuid).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}
