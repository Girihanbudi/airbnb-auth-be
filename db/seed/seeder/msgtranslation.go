package seeder

import (
	translationmodule "airbnb-auth-be/internal/app/translation"

	"gorm.io/gorm"
)

func SeedMsgTranslation(db gorm.DB) error {

	data := []translationmodule.MsgTranslation{
		/*
			Default
		*/
		// En translation
		MakeMsgTranslation("otp", "en", "Your Airbnb verification code is %s"),
		// Id translation
		MakeMsgTranslation("otp", "id", "Kode verifikasi Airbnb Anda adalah %s"),
	}

	var msgTranslationRecords []translationmodule.ErrTranslation
	if err := db.Find(&msgTranslationRecords).Error; err != nil {
		return err
	}

	if len(msgTranslationRecords) > 0 {
		if err := db.Delete(&msgTranslationRecords).Error; err != nil {
			return err
		}
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func MakeMsgTranslation(code, localeCode string, template string) translationmodule.MsgTranslation {
	return translationmodule.MsgTranslation{
		Code:       code,
		LocaleCode: localeCode,
		Template:   template,
	}
}
