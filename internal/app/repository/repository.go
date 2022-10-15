package repository

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/internal/app/dsn"
	"main/internal/app/model"
	"math/rand"
	"strconv"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func New(ctx context.Context) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetPromos() ([]model.Promos, error) {
	var promos []model.Promos
	result := r.db.Find(&promos).Error
	return promos, result

}

func (r *Repository) AddPromo(promo model.Promos) (model.Promos, error) {
	err := r.db.Create(&promo).Error
	if err != nil {
		return model.Promos{}, err
	}
	return promo, nil
}

func RandPromo(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (r *Repository) NewRandRecords() (model.Promos, error) {
	rand.Seed(time.Now().UnixNano())

	price := rand.Intn(990) + 10

	storeList := []string{"Пятёрочка", "Магнит", "ВиТ", "ДОДО", "Яндекс Плюс", "Lamoda", "OZON", "Wildberries"}
	storeRandom := rand.Intn(len(storeList))
	store := storeList[storeRandom]

	discountList := []string{"-" + strconv.Itoa(price*2) + "р", "-10%", "-20%", "-30%"}
	discountRandom := rand.Intn(len(discountList))
	discount := discountList[discountRandom]

	quantity := rand.Intn(5)

	var promoList []string
	for i := 0; i < quantity; i++ {
		promoList = append(promoList, RandPromo(5))
	}

	newPromo := model.Promos{
		Store:    store,                     // магазин
		Discount: discount,                  // скидка
		Price:    strconv.Itoa(price) + "р", // цена от 10 до 999
		Quantity: uint64(quantity),          // оставшееся кол-во от 0 до 4
		Promo:    promoList,                 // слайс промокодов
	}
	err := r.db.Create(&newPromo).Error
	if err != nil {
		return model.Promos{}, err
	}
	return newPromo, nil
}
