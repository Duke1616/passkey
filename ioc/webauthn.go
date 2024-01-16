package ioc

import (
	"passkey-demo/internal/service"
	"passkey-demo/pkg/confer"
	"strings"
)

func InitWebauthn() service.Service {
	rpId := confer.C().Webauthn.RPID
	rpDisplayName := confer.C().Webauthn.RPDisplayName
	rpOrigins := strings.Split(confer.C().Webauthn.RPOrigins, ",")
	return service.NewService(rpDisplayName, rpId, rpOrigins)
}
