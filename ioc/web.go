package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"passkey-demo/internal/web"
	"strings"
	"time"
)

func InitWebServer(mdls []gin.HandlerFunc,
	userHdl *web.UserHandler, webauthnHdl *web.WebauthnHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	userHdl.RegisterRoutes(server)
	webauthnHdl.RegisterRoutes(server)

	server.StaticFile("/", "./views/index.html")
	server.StaticFile("/index.js", "./views/index.js")
	return server
}

func InitGinMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		corsHdl(),
		sessHdl(),
		func(ctx *gin.Context) {
			//println("这是我的 Middleware")
		},
	}
}

func corsHdl() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 你不加这个，前端是拿不到的
		ExposeHeaders: []string{"x-jwt-token", "x-refresh-token"},
		// 是否允许你带 cookie 之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// 你的开发环境
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	})
}

func sessHdl() gin.HandlerFunc {
	store := cookie.NewStore([]byte("111222333"))
	return sessions.Sessions("session", store)
}
