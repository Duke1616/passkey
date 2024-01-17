//go:build wireinject

package main

import (
	"github.com/Duke1616/passkey/cmd/app"
	"github.com/Duke1616/passkey/internal/repository"
	"github.com/Duke1616/passkey/internal/repository/cache"
	"github.com/Duke1616/passkey/internal/repository/dao"
	"github.com/Duke1616/passkey/internal/service"
	"github.com/Duke1616/passkey/internal/web"
	"github.com/Duke1616/passkey/ioc"
	"github.com/google/wire"
)

func InitWebServer() *app.App {
	wire.Build(
		// 第三方依赖
		ioc.InitDB, ioc.InitLoggerSlog,
		ioc.InitRedis,

		// DAO 部分
		dao.NewUserDAO,

		// cache 部分
		cache.NewUserCache,

		// repository 部分
		repository.NewCachedUserRepository,

		// Service 部分
		ioc.InitWebauthn,
		service.NewUserService,

		// handler 部分
		web.NewUserHandler,
		web.NewWebauthnHandler,
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,

		wire.Struct(new(app.App), "*"),
	)
	return new(app.App)
}
