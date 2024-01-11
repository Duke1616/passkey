package ioc

import (
	"passkey-demo/internal/service"
)

func InitWebauthn() service.Service {
	rpId := "localhost"
	rpDisplayName := "WebAuthn Example Application"
	rpOrigins := []string{"http://localhost:8100"}
	return service.NewService(rpDisplayName, rpId, rpOrigins)
}
