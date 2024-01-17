package ioc

import (
	"passkey-demo/config"
	"passkey-demo/internal/service"
	"strings"
)

func InitWebauthn() service.Service {
	rpId := config.C().Webauthn.RPID
	rpDisplayName := config.C().Webauthn.RPDisplayName
	rpOrigins := strings.Split(config.C().Webauthn.RPOrigins, ",")
	return service.NewService(rpDisplayName, rpId, rpOrigins)
}
