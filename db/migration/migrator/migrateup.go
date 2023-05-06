package migration

import (
	accountmodule "airbnb-auth-be/internal/app/account"
	translationmodule "airbnb-auth-be/internal/app/translation"
	orm "airbnb-auth-be/internal/pkg/gorm"
	"airbnb-auth-be/internal/pkg/log"

	"gorm.io/gorm"
)

func MigrateUp(db gorm.DB) {
	models := []interface{}{
		&translationmodule.ErrTranslation{},
		&translationmodule.MsgTranslation{},
		&accountmodule.Account{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal(orm.Instance, "failed to run migration", err)
	}
}
