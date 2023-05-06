package seeder

import (
	"airbnb-auth-be/internal/pkg/env"
	"testing"

	"airbnb-auth-be/internal/pkg/gorm"
)

func TestSeedAll(t *testing.T) {
	env.InitEnv(envOps)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(gorm.Options{Config: config})
	db := *engine.DB

	t.Log("seeding error translation...")
	if err := SeedErrTranslation(db); err != nil {
		t.Error("failed to seed error translation", err)
	}

	t.Log("finish seeding")
}
