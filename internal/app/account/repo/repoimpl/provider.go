package repoimpl

import (
	"airbnb-auth-be/internal/pkg/gorm"
)

type Options struct {
	Gorm *gorm.Engine
}

type Repo struct {
	Options
}

func NewAccountRepo(options Options) *Repo {
	return &Repo{options}
}
