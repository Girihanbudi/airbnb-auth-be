package seeder

import (
	"airbnb-auth-be/internal/pkg/env"
)

var batchSize = 100

var envOps = env.Options{
	Path:     "../../../env",
	FileName: "config",
	Ext:      "yaml",
}
