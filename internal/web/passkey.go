package web

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
	"passkey-demo/internal/service"
	"passkey-demo/internal/service/passkey"
)

//var _ handler = (*UserHandler)(nil)

type WebauthnHandler struct {
	svc     passkey.Service
	userSvc service.UserService
}

func NewWebauthnHandler(svc passkey.Service, userSvc service.UserService) *WebauthnHandler {
	return &WebauthnHandler{
		svc:     svc,
		userSvc: userSvc,
	}
}

func (w *WebauthnHandler) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/api/v1")
	g.GET("/login/begin/:username", w.loginBegin)
	g.POST("/login/finish/:username", w.loginFinish)
	g.GET("/registration/start/:username", w.registrationBegin)
	g.POST("/registration/finish/:username", w.registrationFinish)
	g.GET("/set_session", w.setSession)
	g.GET("/get_session", w.getSession)
}

func (w *WebauthnHandler) setSession(c *gin.Context) {
	//初始化 session 对象
	session := sessions.Default(c) //设置过期时间
	// 过期时间6h
	session.Options(sessions.Options{MaxAge: 3600 * 6})
	//设置 Session
	session.Set("username", "lxx")
	session.Save()
	c.JSON(200, gin.H{"msg": "设置session成功----userrname:lxx"})
}

func (w *WebauthnHandler) getSession(c *gin.Context) {
	session := sessions.Default(c)
	// 通过 session.Get 读取 session 值
	username := session.Get("username")
	fmt.Println("get", username)
	c.JSON(200, gin.H{"msg": "获取session成功"})
}

func (w *WebauthnHandler) loginBegin(ctx *gin.Context) {
	username := ctx.Param("username")
	// TODO 目前是注册即创建用户
	webauthnUser, err := w.userSvc.FindOrCreateByWebauthn(ctx, username)
	if err != nil {
		return
	}

	options, sessionData, err := w.svc.BeginLogin(ctx, &webauthnUser)
	if err != nil {
		return
	}

	// 保存session
	setSess(ctx, "login", sessionData)

	ctx.JSON(http.StatusOK, options)
}

func (w *WebauthnHandler) loginFinish(ctx *gin.Context) {
	username := ctx.Param("username")

	webauthnUser, err := w.userSvc.FindOrCreateByWebauthn(ctx, username)
	if err != nil {
		return
	}

	sessionData, err := getSess(ctx, "login")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	credential, err := w.svc.FinishLogin(&webauthnUser, sessionData, ctx.Request)
	if err != nil {
		return
	}

	fmt.Println(credential)

	ctx.JSON(200, "Login Success")
}

func (w *WebauthnHandler) registrationBegin(ctx *gin.Context) {
	//var req LoginWebauthnReq
	//if err := ctx.Bind(&req); err != nil {
	//	return
	//}

	username := ctx.Param("username")
	// TODO 目前是注册即创建用户
	webauthnUser, err := w.userSvc.FindOrCreateByWebauthn(ctx, username)
	if err != nil {
		return
	}

	// BeginRegistration
	options, sessionData, err := w.svc.BeginRegistration(ctx, &webauthnUser)

	// 保存session
	setSess(ctx, "register", sessionData)

	ctx.JSON(http.StatusOK, options)
}

func (w *WebauthnHandler) registrationFinish(ctx *gin.Context) {
	//var req LoginWebauthnReq
	//if err := ctx.Bind(&req); err != nil {
	//	return
	//}

	username := ctx.Param("username")
	webAuthnUser, err := w.userSvc.FindOrCreateByWebauthn(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	sessionData, err := getSess(ctx, "register")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	credential, err := w.svc.FinishRegistration(&webAuthnUser, sessionData, ctx.Request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	webAuthnUser.Credentials = append(webAuthnUser.Credentials, *credential)

	err = w.userSvc.Update(ctx, webAuthnUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, "Registration Success")
}

func setSess(ctx *gin.Context, key string, sessionData *webauthn.SessionData) {
	sess := sessions.Default(ctx)
	data, _ := json.Marshal(sessionData)

	sess.Set(key, data)
	//sess.Options(sessions.Options{
	//	MaxAge: 30,
	//})

	sess.Save()
}

func getSess(ctx *gin.Context, key string) (webauthn.SessionData, error) {
	session := sessions.Default(ctx)
	data := session.Get(key)
	sessionData := webauthn.SessionData{}
	if err := json.Unmarshal(data.([]byte), &sessionData); err != nil {
		return sessionData, err
	}

	return sessionData, nil
}
