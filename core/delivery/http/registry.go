package http

import (
	"go-socket/config"
	"go-socket/core/delivery/http/handler"
)

func BuildRegistry(config *config.Config) map[string]routingConfig {
	return map[string]routingConfig{
		"POST:/api/v1/auth/login": {
			handler: handler.NewLoginHandler(),
		},
	}
}
