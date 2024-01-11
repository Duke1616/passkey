package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"passkey-demo/internal/domain"
	"passkey-demo/internal/repository"
	repomocks "passkey-demo/internal/repository/mocks"
	"testing"
)

func TestFindOrCreateByWebauthn(t *testing.T) {
	testCases := []struct {
		name string

		mock func(ctrl *gomock.Controller) repository.UserRepository

		// 预期输入
		ctx      context.Context
		username string

		wantUser domain.User
		wantErr  error
	}{
		{
			name: "查询成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().
					FindByUsername(gomock.Any(), "passkey").
					Return(domain.User{
						Username: "passkey",
					}, nil)
				return repo
			},

			username: "passkey",
			wantUser: domain.User{
				Username: "passkey",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserService(repo)
			user, err := svc.FindOrCreateByWebauthn(tc.ctx, tc.username)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, user)
		})
	}
}
