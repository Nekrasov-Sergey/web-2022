package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/internal/app/dsn"
	"main/internal/app/model"
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

func (r *Repository) GetPromos() ([]model.Promos, error) {
	var promo []model.Promos
	err := r.db.Find(&promo).Error
	return promo, err
}

func (r *Repository) GetPromoPrice(uuid uuid.UUID, promo *model.Promos) (uint64, error) {
	err := r.db.First(&promo, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	return 0, nil
}

func (r *Repository) AddPromo(promo model.Promos) error {
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

func (r *Repository) NewRandRecords() error {
	rand.Seed(time.Now().UnixNano())

	price := rand.Intn(990) + 10

	discount := price * 2

	storeList := []string{"Пятёрочка", "Магнит", "Лента", "ВиТ", "ДОДО", "Яндекс Плюс", "Lamoda", "OZON", "Wildberries"}
	storeRandom := rand.Intn(len(storeList))
	store := storeList[storeRandom]
	quantity := rand.Intn(5) + 1

	var promoList []string
	for i := 0; i < quantity; i++ {
		promoList = append(promoList, RandPromo(5))
	}

	newPromo := model.Promos{
		Store:    store,            // магазин
		Discount: uint64(discount), // скидка
		Price:    uint64(price),    // цена от 10 до 999
		Quantity: uint64(quantity), // оставшееся кол-во от 0 до 4
		Promo:    promoList,        // слайс промокодов
		Image:    image[store],
	}
	err := r.db.Create(&newPromo).Error
	return err
}

func (r *Repository) ChangePrice(uuid uuid.UUID, price uint64) (int, error) {
	var promo model.Promos
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

func (r *Repository) DeletePromo(uuid uuid.UUID) (int, error) {
	var promo model.Promos
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
