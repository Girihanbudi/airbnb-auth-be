package user

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	UserId       string    `json:"user_id" gorm:"primaryKey"`
	Provider     string    `json:"provider" gorm:"primaryKey"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
	TokenType    string    `json:"token_type"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (m *Account) BeforeCreate(tx *gorm.DB) (err error) {

	if m.UserId == "" {
		err = errors.New("user id cannot empty")
		return
	}

	if m.Provider == "" {
		err = errors.New("provider cannot empty")
		return
	}

	return
}
