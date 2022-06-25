package middleware

import (
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var uid string
		//token := c.GetHeader("X-Token")
		//if token != "" {
		//	uid, _ = cache.GetUserByToken(token)
		//} else {
		//	session := sessions.Default(c)
		//	uid, _ = session.Get("user_id").(string)
		//}
		//if uid != "" {
		//	user, err := model.GetUser(uid)
		//	if err == nil {
		//		c.Set("user", &user)
		//	}
		//}
		c.Next()
	}
}

// CurrentAdminUser 获取登录管理员用户
func CurrentAdminUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var uid string
		//session := sessions.Default(c)
		//uid, _ = session.Get("admin_user_id").(string)
		//if uid != "" {
		//	adminUser, err := model.GetAdminUser(uid)
		//	if err == nil {
		//		c.Set("admin_user", &adminUser)
		//	}
		//}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// app登录校验
		//if user, _ := c.Get("user"); user != nil {
		//	if _, ok := user.(*model.User); ok {
		//		c.Next()
		//		return
		//	}
		//}
		//// admin登录校验
		//if adminUser, _ := c.Get("admin_user"); adminUser != nil {
		//	if _, ok := adminUser.(*model.AdminUser); ok {
		//		c.Next()
		//		return
		//	}
		//}
		//c.JSON(http.StatusOK, serializer.CheckLogin())
		c.Abort()
	}
}
