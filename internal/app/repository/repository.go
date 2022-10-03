package repository

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/internal/app/ds"
	"main/internal/app/dsn"
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

func (r *Repository) GetPromoByID(id uint) (*ds.Promos, error) {
	promo := &ds.Promos{}

	err := r.db.First(promo, id).Error
	if err != nil {
		return nil, err
	}

	return promo, nil
}

func (r *Repository) NewRandRecords() error {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	price := rand.Intn(990) + 10
	storeList := []string{"Пятёрочка", "Магнит", "Вит", "ДОДО", "Яндекс Плюс", "Lamoda"}
	storeRandom := rand.Intn(len(storeList))
	store := storeList[storeRandom]
	newPromo := ds.Promos{
		Code:  uint(code),                  // код от 100000 до 999999
		Price: strconv.Itoa(price) + "р",   // цена от 10 до 999
		Promo: strconv.Itoa(price*2) + "р", //промо даёт в 2 раза больше цены промо
		Store: store,
	}
	err := r.db.Create(&newPromo).Error
	if err != nil {
		return err
	}
	return nil
}
