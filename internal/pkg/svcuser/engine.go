package svcuser

import (
	"airbnb-auth-be/internal/pkg/log"
)

func (svc *Client) Stop() error {
	log.Event(Instance, "close service connection...")
	return svc.RpcConn.Close()
}
