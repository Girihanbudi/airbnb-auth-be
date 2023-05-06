package svcuser

import (
	"airbnb-auth-be/internal/pkg/log"
	"airbnb-auth-be/internal/pkg/svcuser/config"
	"airbnb-auth-be/internal/pkg/svcuser/transport/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const Instance = "User Client"

type Options struct {
	config.Config
}

type Client struct {
	Options
	RpcConn *grpc.ClientConn
	User    rpc.UserServiceClient
	Locale  rpc.LocaleServiceClient
	Country rpc.CountryServiceClient
}

func NewClient(options Options) *Client {
	conn, err := grpc.Dial(options.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(Instance, "connection failed", err)
	}

	return &Client{
		Options: options,
		RpcConn: conn,
		User:    rpc.NewUserServiceClient(conn),
		Locale:  rpc.NewLocaleServiceClient(conn),
		Country: rpc.NewCountryServiceClient(conn),
	}
}
