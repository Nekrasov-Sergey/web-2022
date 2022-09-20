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

func (r *Repository) GetProductByID(id uint) (*ds.Product, error) {
	product := &ds.Product{}

	err := r.db.First(product, id).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *Repository) NewRandRecord() error {
	newProduct := ds.Product{
		Code:  uint(rand.Intn(900000) + 100000), // код от 100000 до 999999
		Price: uint(rand.Intn(9000) + 1000),     // цена от 1000 до 9999
	}
	err := r.db.Create(&newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateProduct(product ds.Product) error {
	return r.db.Create(product).Error
}
