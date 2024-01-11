package service

import (
	"context"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
	"passkey-demo/internal/domain"
)

//go:generate mockgen -source=./passkey.go -package=svcmocks -destination=./mocks/passkey.mock.go Service
type Service interface {
	BeginRegistration(ctx context.Context, user *domain.User) (creation *protocol.CredentialCreation,
		session *webauthn.SessionData, err error)
	FinishRegistration(user *domain.User, session webauthn.SessionData, request *http.Request) (
		creation *webauthn.Credential, err error)
	BeginLogin(ctx context.Context, user *domain.User) (creation *protocol.CredentialAssertion,
		session *webauthn.SessionData, err error)
	FinishLogin(user *domain.User, session webauthn.SessionData, request *http.Request) (
		creation *webauthn.Credential, err error)
}

type service struct {
	webAuthn *webauthn.WebAuthn
}

func NewService(RPDisplayName, RPID string, RPOrigins []string) Service {
	wa, err := webauthn.New(&webauthn.Config{
		RPDisplayName: RPDisplayName,
		RPID:          RPID,
		RPOrigins:     RPOrigins,
	})

	if err != nil {

	}

	return &service{
		webAuthn: wa,
	}
}

func (s *service) BeginRegistration(ctx context.Context, user *domain.User) (creation *protocol.CredentialCreation,
	session *webauthn.SessionData, err error) {
	return s.webAuthn.BeginRegistration(user, webauthn.WithAuthenticatorSelection(protocol.AuthenticatorSelection{
		ResidentKey:      protocol.ResidentKeyRequirementRequired,
		UserVerification: protocol.VerificationPreferred,
	}))
}

func (s *service) FinishRegistration(user *domain.User, session webauthn.SessionData, request *http.Request) (
	creation *webauthn.Credential, err error) {
	return s.webAuthn.FinishRegistration(user, session, request)
}

func (s *service) BeginLogin(ctx context.Context, user *domain.User) (creation *protocol.CredentialAssertion,
	session *webauthn.SessionData, err error) {
	return s.webAuthn.BeginLogin(user, webauthn.WithUserVerification(protocol.VerificationPreferred))
}

func (s *service) FinishLogin(user *domain.User, session webauthn.SessionData, request *http.Request) (
	creation *webauthn.Credential, err error) {
	return s.webAuthn.FinishLogin(user, session, request)
}
