package repository

import (
	"context"
	"github.com/Duke1616/passkey/internal/domain"
	"github.com/Duke1616/passkey/internal/repository/cache"
	"github.com/Duke1616/passkey/internal/repository/dao"
)

//go:generate mockgen -source=./user.go -package=repomocks -destination=./mocks/user.mock.go UserRepository
type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	Update(ctx context.Context, art domain.User) error
	FindByUsername(ctx context.Context, username string) (domain.User, error)
}

type CachedUserRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

func NewCachedUserRepository(userDao dao.UserDAO, userCache cache.UserCache) UserRepository {
	return &CachedUserRepository{
		dao:   userDao,
		cache: userCache,
	}
}

func (repo *CachedUserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, repo.toEntity(u))
}

func (repo *CachedUserRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	u, err := repo.dao.FindByUsername(ctx, username)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *CachedUserRepository) Update(ctx context.Context, art domain.User) error {
	return repo.dao.UpdateById(ctx, repo.toEntity(art))
}

func (repo *CachedUserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:          u.Id,
		Username:    u.Username,
		Credentials: u.Credentials,
	}
}

func (repo *CachedUserRepository) toEntity(u domain.User) dao.User {
	return dao.User{
		Id:          u.Id,
		Username:    u.Username,
		Credentials: u.Credentials,
	}
}
