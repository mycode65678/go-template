package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go-hash/common"
	"time"
)

//检测登录
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//写入session中
		var store = common.Session()
		session, _ := store.Get(c.Request, "admin")

		if session.Values["id"] == nil {
			c.AbortWithStatus(403)
		}

		if _, ok := session.Values["expired"]; !ok {
			c.AbortWithStatus(403)
		}

		// 判断是否过期了
		expired := session.Values["expired"].(int64)
		fmt.Println("expired", time.Now().Unix(), expired)
		if time.Now().Unix() > expired {
			c.AbortWithStatus(403)
		}
		// 更新时间
		mm, _ := time.ParseDuration("300m")
		mm1 := time.Now().Local().Add(mm)
		// 更新时间
		session.Values["expired"] = mm1.Unix()
		sessions.Save(c.Request, c.Writer)

		fmt.Println("type", session.Values["type"])
		c.Set("admin_id", session.Values["id"].(int64))
		if session.Values["type"].(int) == 1 {
			c.Next()
			return
		}

		//fmt.Println("gin",c.Request.URL.String())
		c.Next()
	}
}

func Router(r *gin.RouterGroup) {
	//r.POST("/login", Login)
	// TODO 上线进行删除
	//r.POST("/login", TestLogin)

	//r.Any("/add-time22", Game22)
	//r.Any("/add-time20", Game20)
	//r.Any("/add-time23", Game23)

	// 权限验证
	r.Use(CheckLogin())
	{

	}
}
