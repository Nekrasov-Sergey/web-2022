package repository

import (
	"github.com/google/uuid"
	"main/internal/app/ds"
)

func (r *Repository) Register(user *ds.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.User, error) {
	user := &ds.User{}

	err := r.db.First(&user, "name = ?", login).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserByUUID(uuid uuid.UUID) (string, error) {
	user := &ds.User{}
	err := r.db.First(&user, "uuid = ?", uuid).Error
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
