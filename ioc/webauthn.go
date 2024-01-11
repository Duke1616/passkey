package ioc

import (
	"passkey-demo/internal/service/passkey"
)

func InitWebauthn() passkey.Service {
	rpId := "localhost"
	rpDisplayName := "WebAuthn Example Application"
	rpOrigins := []string{"http://localhost:8100"}
	return passkey.NewService(rpDisplayName, rpId, rpOrigins)
}
