package repository

import (
	"context"
	"github.com/Duke1616/passkey/internal/domain"
	"github.com/Duke1616/passkey/internal/repository/dao"
	daomocks "github.com/Duke1616/passkey/internal/repository/dao/mocks"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCachedUserRepository_FindByUsername(t *testing.T) {
	nowMs := time.Now().UnixMilli()
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) dao.UserDAO

		ctx      context.Context
		username string

		wantUser domain.User
		wantErr  error
	}{
		{
			name: "查找成功",
			mock: func(ctrl *gomock.Controller) dao.UserDAO {
				username := "passkey"
				d := daomocks.NewMockUserDAO(ctrl)
				d.EXPECT().FindByUsername(gomock.Any(), username).
					Return(dao.User{
						Username:    username,
						Credentials: []webauthn.Credential{},
						Ctime:       nowMs,
						Utime:       102,
					}, nil)
				return d
			},
			username: "passkey",
			ctx:      context.Background(),
			wantUser: domain.User{
				Username:    "passkey",
				Credentials: []webauthn.Credential{},
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := tc.mock(ctrl)
			svc := NewCachedUserRepository(uc, nil)
			user, err := svc.FindByUsername(tc.ctx, tc.username)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, user)
		})
	}
}
