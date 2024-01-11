package domain

import (
	"encoding/binary"
	"github.com/go-webauthn/webauthn/webauthn"
)

type User struct {
	Id int64
	//Email       string
	//Password    string
	Username string
	Nickname string
	// YYYY-MM-DD
	//Birthday time.Time
	//AboutMe  string

	//Phone string

	// UTC 0 的时区
	//Ctime time.Time

	// passkey 登录方式
	Credentials []webauthn.Credential
}

// TodayIsBirthday 判定今天是不是我的生日
//func (u *User) TodayIsBirthday() bool {
//	now := time.Now()
//	return now.Month() == u.Birthday.Month() && now.Day() == u.Birthday.Day()
//}

func (u *User) WebAuthnID() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, uint64(u.Id))
	return buf
}

func (u *User) WebAuthnName() string {
	return u.Username
}

func (u *User) WebAuthnDisplayName() string {
	return u.Username
}

func (u *User) WebAuthnIcon() string {
	return ""
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}
