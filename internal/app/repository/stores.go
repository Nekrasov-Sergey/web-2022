package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"main/internal/app/ds"
	"math/rand"
	"strings"
	"time"
)

func (r *Repository) GetStores() ([]ds.Store, error) {
	var stores []ds.Store
	err := r.db.Order("uuid").Find(&stores).Error
	return stores, err
}

func (r *Repository) GetStore(UUID uuid.UUID) (ds.Store, error) {
	var store ds.Store
	err := r.db.First(&store, UUID).Error
	return store, err
}

func (r *Repository) GetStoreName(uuid uuid.UUID) (string, error) {
	var store ds.Store
	err := r.db.Select("name").First(&store, "uuid = ?", uuid).Error
	return store.Name, err
}

func (r *Repository) GetPromoStore(quantity uint64, storeUUID uuid.UUID, userUUID uuid.UUID) (int, string, error) {
	var store ds.Store
	err := r.db.First(&store, storeUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, "", err
		}
		return 500, "", err
	}

	var cart ds.Cart
	err = r.db.First(&cart, storeUUID, userUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, "", err
		}
		return 500, "", err
	}
	err = r.db.Delete(&cart, storeUUID, userUUID).Error
	if err != nil {
		return 500, "", err
	}

	err = r.AddOrder(userUUID, storeUUID, quantity)
	if err != nil {
		return 500, "", err
	}

	PromoSlice := store.Promo[0:quantity]
	PromoString := strings.Join(PromoSlice, ", ")
	if store.Quantity > 0 {
		err = r.db.Model(&store).Update("Promo", store.Promo[quantity:]).Error
		if err != nil {
			return 500, "", err
		}
		err = r.db.Model(&store).Update("Quantity", store.Quantity-quantity).Error
		if err != nil {
			return 500, "", err
		}
	}

	if store.Quantity == 0 {
		err = r.db.Delete(&store, storeUUID).Error
		if err != nil {
			return 500, "", err
		}
	}

	return 0, PromoString, nil
}

func (r *Repository) CreateStore(store ds.Store) error {
	err := r.db.Create(&store).Error
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

	price := (rand.Intn(100) + 1) * 10

	discount := price * 2

	nameList := []string{"Пятёрочка", "Магнит", "Лента", "ВиТ", "ДОДО", "Яндекс Плюс", "Lamoda", "OZON", "Wildberries"}
	nameRandom := rand.Intn(len(nameList))
	name := nameList[nameRandom]
	quantity := rand.Intn(20) + 1

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

func (r *Repository) ChangeStore(uuid uuid.UUID, store ds.Store) (int, error) {
	store.UUID = uuid

	err := r.db.Model(&store).Updates(ds.Store{Name: store.Name, Discount: store.Discount, Price: store.Price, Quantity: store.Quantity, Promo: store.Promo, Image: store.Image}).Error
	if err != nil {
		return 500, err
	}

	return 0, nil
}

func (r *Repository) DeleteStore(uuid uuid.UUID) (int, error) {
	var store ds.Store
	err := r.db.First(&store, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Delete(&store, uuid).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}
