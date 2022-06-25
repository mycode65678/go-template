package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

//返回session
func Session() *sessions.CookieStore {
	sessionKey := viper.GetString("base.session_key")

	//写入session中
	return sessions.NewCookieStore([]byte(sessionKey))
}

func GetSession(c *gin.Context, name string) *sessions.Session {
	//store.Options =
	s, _ := store.Get(c.Request, name)
	return s
}
