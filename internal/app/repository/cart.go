package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"main/internal/app/ds"
)

func (r *Repository) GetCart(userUUID uuid.UUID) ([]ds.Cart, error) {
	var cart []ds.Cart
	err := r.db.Order("store_uuid").Find(&cart, "user_uuid = ?", userUUID).Error
	return cart, err
}

func (r *Repository) GetCart1(storeUUID uuid.UUID, userUUID uuid.UUID) (ds.Cart, error) {
	var cart ds.Cart
	err := r.db.First(&cart, "store_uuid = ? and user_uuid = ?", storeUUID, userUUID).Error
	return cart, err
}

func (r *Repository) DeleteCart(storeUUID uuid.UUID, userUUID uuid.UUID) (int, error) {
	var cart ds.Cart
	err := r.db.First(&cart, "store_uuid = ? and user_uuid = ?", storeUUID, userUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}

	err = r.db.Delete(&cart, "store_uuid = ? and user_uuid = ?", storeUUID, userUUID).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}

func (r *Repository) IncreaseQuantity(storeUUID uuid.UUID, userUUID uuid.UUID) (uint64, error) {
	var store ds.Store
	err := r.db.First(&store, storeUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}

	var cart ds.Cart
	err = r.db.First(&cart, "store_uuid = ? and user_uuid = ?", storeUUID, userUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart.StoreUUID = storeUUID
			cart.UserUUID = userUUID
			cart.Quantity = 0
			err = r.db.Create(cart).Error
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	if cart.Quantity < store.Quantity {
		err = r.db.Model(&cart).Where("store_uuid = ? and user_uuid = ?", storeUUID, userUUID).Update("Quantity", cart.Quantity+1).Error
		if err != nil {
			return 0, err
		}
	}

	return cart.Quantity, nil
}

func (r *Repository) DecreaseQuantity(storeUUID uuid.UUID, userUUID uuid.UUID) (uint64, int, error) {
	var cart ds.Cart
	err := r.db.First(&cart, "store_uuid = ? and user_uuid = ?", storeUUID, userUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, 404, err
		}
		return 0, 500, err
	}
	if cart.Quantity > 0 {
		err = r.db.Model(&cart).Where("store_uuid = ? and user_uuid = ?", storeUUID, userUUID).Update("Quantity", cart.Quantity-1).Error
		if err != nil {
			return 0, 500, err
		}
		if cart.Quantity == 0 {
			err = r.db.Delete(&cart, "store_uuid = ? and user_uuid = ?", storeUUID, userUUID).Error
			if err != nil {
				return 0, 500, err
			}
			return 0, 0, nil
		}
	}
	return cart.Quantity, 0, nil
}

func (r *Repository) DeleteByUser(userUUID uuid.UUID) error {
	var cart ds.Cart
	err := r.db.Where("user_uuid = ?", userUUID).Delete(&cart).Error
	if err != nil {
		return err
	}
	return nil
}
