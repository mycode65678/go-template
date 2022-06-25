package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-hash/a"
	"go-hash/web/admin"
	"go-hash/web/middleware"
)

func AddHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Header("Access-Control-Allow-Origin","http://192.168.1.7:9527")
		allow := viper.GetString("base.allow")
		c.Header("Access-Control-Allow-Origin", allow)
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
			c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Content-Type", "application/json")
			c.AbortWithStatus(200)
			return
		}
		//c.Header("Access-Control-Allow-Origin","http://localhost:8080")
		//c.Header("Access-Control-Allow-Origin","http://192.168.1.7:9528")

		/*session := sessions.Default(c)
		res := session.Get("user")
		if res == nil {
			c.JSON(403,"Permission verification failed")
			c.Abort()
		} else {
			c.Next()
		}*/
		c.Next()
	}
}

func StartHttp() {
	a.Config()
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	r.Use(middleware.Session(viper.GetString("base.session_key")))
	api := r.Group("/api")
	// 管理员路由表
	admin.Router(api.Group("/admin"))
	// 代理路由表
	//proxy.Router(api.Group("/proxy"))
	// 会员路由表
	//user.Router(api.Group("/user"))

	//r.POST("/server-to-client", socket.ServerToClient)
	fmt.Println("port", viper.GetString("web.port"))
	r.Run("0.0.0.0:" + viper.GetString("web.port"))
}
