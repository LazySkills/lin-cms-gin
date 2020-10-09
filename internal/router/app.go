/** Created By 嗝嗝<354007048@qq.com>. Date 2020/9/8 */
package router

import (
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/app/admin"
	v1 "lin-cms-gin/internal/app/v1"
	"lin-cms-gin/pkg/setting"
	"lin-cms-gin/pkg/tools"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)


	// v1 相关操作
	v1.IncludePermissionModule()
	v1.IncludeRouter(r)

	// admin 相关操作
	admin.IncludePermissionModule()
	admin.IncludeRouter(r)

	tools.InitSwagger(r)
	return r
}
