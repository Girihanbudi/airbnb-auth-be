package translation

import "time"

type MsgTranslation struct {
	Code       string `json:"code" gorm:"primaryKey"`
	LocaleCode string `json:"locale_code" gorm:"primaryKey"`
	Template   string `json:"template"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
