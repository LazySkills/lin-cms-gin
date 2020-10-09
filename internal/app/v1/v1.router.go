/** Created By wene<354007048@qq.com> . Date at 2020/7/5 */
package v1

import (
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/middleware/jwt"
	"lin-cms-gin/internal/middleware/permission"
	"lin-cms-gin/pkg/lin"
)

// 权限注册蓝图
func IncludePermissionModule()  {
	// 注册项目权限
	lin.Include(&BookPermission{})
}

// 路由注册蓝图
func IncludeRouter(r *gin.Engine)  {
	apiV1 := r.Group("/v1")
	apiV1.Use(jwt.JWT(), permission.Permission())
	{
		//获取图书列表
		apiV1.GET("/book", GetBooks)
		//获取图书列表
		apiV1.GET("/book/:id", GetBooksInfo)
		//新建图书
		apiV1.POST("/book", AddBook)
		//更新指定图书
		apiV1.PUT("/book/:id", EditBook)
		//删除指定图书
		apiV1.DELETE("/book/:id", DeleteBook)
	}
}

