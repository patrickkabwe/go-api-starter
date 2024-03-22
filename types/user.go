package types

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName string    `json:"firstName" gorm:"not null"`
	LastName  string    `json:"lastName" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique; not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt int64     `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt int64     `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
