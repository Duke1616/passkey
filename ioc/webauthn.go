package ioc

import (
	"github.com/Duke1616/passkey/config"
	"github.com/Duke1616/passkey/internal/service"
	"strings"
)

func InitWebauthn() service.Service {
	rpId := config.C().Webauthn.RPID
	rpDisplayName := config.C().Webauthn.RPDisplayName
	rpOrigins := strings.Split(config.C().Webauthn.RPOrigins, ",")
	return service.NewService(rpDisplayName, rpId, rpOrigins)
}
