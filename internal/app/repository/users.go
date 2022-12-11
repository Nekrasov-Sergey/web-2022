package repository

import "main/internal/app/ds"

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
