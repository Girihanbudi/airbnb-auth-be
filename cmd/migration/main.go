package main

import (
	m "airbnb-auth-be/db/migration/migrator"
	"airbnb-auth-be/internal/pkg/env"
	"airbnb-auth-be/internal/pkg/gorm"
	"airbnb-auth-be/internal/pkg/log"
	"flag"
)

var Instance = "Migration"

func main() {
	cmd := flag.String("migration", "", "migration direction (up/down)")
	flag.Parse()
	if cmd == nil || *cmd == "" {
		log.Fatal(Instance, "migration failed with arguments not found", nil)
	}

	defaultEnvOps := env.NewDefaultOptions()
	env.InitEnv(defaultEnvOps)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(gorm.Options{Config: config})

	switch *cmd {
	case m.MigrationUp.String():
		log.Event(Instance, "migrating...")
		m.MigrateUp(*engine.DB)
		log.Event(Instance, "finish migrating")
	default:
		log.Fatal(Instance, "unknown command", nil)
	}
}
