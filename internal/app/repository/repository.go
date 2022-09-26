package repository

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/internal/app/ds"
	"main/internal/app/dsn"
	"math/rand"
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
	newPromo := ds.Promos{
		Code:  uint(rand.Intn(900000) + 100000), // код от 100000 до 999999
		Price: uint(rand.Intn(990) + 10),        // цена от 10 до 999
	}
	err := r.db.Create(&newPromo).Error
	if err != nil {
		return err
	}
	return nil
}
