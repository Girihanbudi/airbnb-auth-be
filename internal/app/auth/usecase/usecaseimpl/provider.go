package usecaseimpl

import (
	accountrepo "airbnb-auth-be/internal/app/account/repo"
	kafkaproducer "airbnb-auth-be/internal/pkg/kafka/producer"
	"airbnb-auth-be/internal/pkg/oauth/facebook"
	"airbnb-auth-be/internal/pkg/oauth/google"
	"airbnb-auth-be/internal/pkg/svcuser"
)

type Options struct {
	GoogleOauth   google.Oauth
	FacebookOauth facebook.Oauth
	AccountRepo   accountrepo.IAccount
	EventProducer *kafkaproducer.Producer
	SvcUser       *svcuser.Client
}

type Usecase struct {
	Options
}

// Auth Usecase provide a module for authentication
func NewAuthUsecase(options Options) *Usecase {
	return &Usecase{options}
}
