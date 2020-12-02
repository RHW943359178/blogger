package utils

import (
	session "blogger/cookie_session"
	"github.com/gin-gonic/gin"
	"log"
)

/**
cookie_session 中间件
*/
func SessionMiddleware() gin.HandlerFunc {
	//	生成 session 管理器对象
	mgr, err := session.CreateSessionMgr(session.Redis, "81.69.255.188:6379")
	if err != nil {
		log.Fatalf("Create manager obj failed, err: %v\n", err)
	}
	//	初始化 session 中间件
	sm := session.SessionMiddleware(mgr, session.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   60 * 60 * 24 * 30, //	秒钟
		Secure:   false,
		HttpOnly: false,
	})
	return sm
}
