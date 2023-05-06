package translation

import (
	"time"
)

type ErrTranslation struct {
	Code       string `json:"code" gorm:"primaryKey"`
	LocaleCode string `json:"locale_code" gorm:"primaryKey"`
	Message    string `json:"message" gorm:"not null"`
	HttpCode   int    `json:"http_code" gorm:"not null"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
