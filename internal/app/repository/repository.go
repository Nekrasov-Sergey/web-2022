package repository

import (
	"context"
	"github.com/google/uuid"
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
	err := r.db.Find(&promos).Error
	return promos, err
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

func (r *Repository) NewRandRecords() error {
	rand.Seed(time.Now().UnixNano())

	price := rand.Intn(990) + 10

	storeList := []string{"Пятёрочка", "Магнит", "Лента", "ВиТ", "ДОДО", "Яндекс Плюс", "Lamoda", "OZON", "Wildberries"}
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
	return err
}

func (r *Repository) ChangePrice(uuid uuid.UUID, price string) error {
	var promo model.Promos
	promo.UUID = uuid
	err := r.db.Model(&promo).Update("Price", price).Error
	return err
}

func (r *Repository) DeletePromo(uuid string) error {
	var promo model.Promos
	err := r.db.Delete(&promo, "uuid", uuid).Error
	return err
}
