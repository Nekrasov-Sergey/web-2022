package repository

import (
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"main/internal/app/ds"
	"time"
)

func (r *Repository) AddOrder(order ds.Order) error {
	var names []string
	log.Println(order.Stores)
	for _, val := range order.Stores {
		name, err := r.GetStoreName(val)
		if err != nil {
			return err
		}
		names = append(names, name)
	}
	log.Println(names)

	order.Stores = names
	//date := time.Now().Add(time.Hour * 3)
	var err error
	order.Date = time.Now() //, err = time.Parse("2006-01-02 15:04:05", date.Format("2006-01-02 15:04:05"))
	order.Status = "Оформлен"
	log.Println(order)

	err = r.db.Create(&order).Error
	if err != nil {
		return err
	}
	err = r.DeleteByUser(order.UserUUID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetOrders() ([]ds.Order, error) {
	var orders []ds.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *Repository) ChangeStatus(uuid uuid.UUID, status string) (int, error) {
	var order ds.Order
	err := r.db.First(&order, "uuid = ?", uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Model(&order).Update("Status", status).Error
	//if errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		return 500, err
	}
	return 0, nil
}
