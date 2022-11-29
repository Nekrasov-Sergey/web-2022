package ds

import (
	"github.com/google/uuid"
	"main/internal/app/role"
)

type User struct {
	UUID uuid.UUID `db:"uuid" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name string    `db:"name"`
	Role role.Role `db:"role" sql:"type:string;"`
	Pass string    `db:"pass"`
}
