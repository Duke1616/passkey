//go:build wireinject

package main

import (
	"github.com/google/wire"
	"passkey-demo/internal/repository"
	"passkey-demo/internal/repository/dao"
	"passkey-demo/internal/service"
	"passkey-demo/internal/web"
	"passkey-demo/ioc"
)

func InitWebServer() *App {
	wire.Build(
		// 第三方依赖
		ioc.InitDB, ioc.InitLoggerSlog,
		// DAO 部分
		dao.NewUserDAO,
		// cache 部分

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

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
