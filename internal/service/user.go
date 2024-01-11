package service

import (
	"context"
	"passkey-demo/internal/domain"
	"passkey-demo/internal/repository"
	"passkey-demo/internal/repository/dao"
)

//go:generate mockgen -source=./user.go -package=svcmocks -destination=./mocks/user.mock.go UserService
type UserService interface {
	FindOrCreateByWebauthn(ctx context.Context, username string) (domain.User, error)
	Update(ctx context.Context, u domain.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) FindOrCreateByWebauthn(ctx context.Context, username string) (domain.User, error) {
	u, err := svc.repo.FindByUsername(ctx, username)

	if err == nil {
		return u, err
	}

	err = svc.repo.Create(ctx, domain.User{
		Username: username,
	})
	if err != nil && err != dao.ErrDuplicateUsername {
		return domain.User{}, err
	}
	return svc.repo.FindByUsername(ctx, username)
}

func (svc *userService) Update(ctx context.Context, u domain.User) error {
	return svc.repo.Update(ctx, u)
}
