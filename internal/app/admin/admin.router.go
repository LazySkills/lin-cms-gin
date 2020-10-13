/** Created By wene<354007048@qq.com> . Date at 2020/7/5 */
package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/middleware/jwt"
	"lin-cms-gin/internal/middleware/permission"
	"lin-cms-gin/pkg/lin"
	"lin-cms-gin/pkg/setting"
	"net/http"
)

// 权限注册蓝图
func IncludePermissionModule()  {
	// 注册项目权限
	lin.Include(&LinUserPermission{})
	lin.Include(&LinAdminPermission{})
	lin.Include(&LinFilePermission{})
}

// 路由注册蓝图
func IncludeRouter(r *gin.Engine)  {
	// 设置静态文件浏览
	r.StaticFS("/assets", http.Dir(fmt.Sprintf("tmp/%s", setting.FileSetting.FileSavePath)))

	apiV1 := r.Group("/cms")

	//用户-登录
	apiV1.POST("/user/login", UserLogin)
	//用户-刷新授权
	apiV1.GET("/user/refresh", UserGetToken)

	apiV1.Use(jwt.JWT(), permission.Permission())
	{
		// 用户-用户注册
		apiV1.POST("/user/register", UserRegister)
		// 用户-用户更新信息
		apiV1.PUT("/user", UserUpdate)
		// 用户-修改密码
		apiV1.PUT("/user/change_password", UserUpdatePassword)
		// 用户-获取权限
		apiV1.GET("/user/permissions", UserGetPermissions)
		// 用户-查询自己信息
		apiV1.GET("/user/information", UserGetInFormation)

		//文件管理-上传文件
		r.MaxMultipartMemory = setting.FileSetting.FileMultipartMaxSize
		apiV1.POST("/file", UploadImage)


		// 管理员-查询所有可分配的权限
		apiV1.GET("/admin/permission", GetAllPermissions)
		// 管理员-查询所有用户
		apiV1.GET("/admin/users", GetAdminUsers)
		// 管理员-修改用户密码
		apiV1.PUT("/admin/user/:id/password", ChangeUserPassword)
		// 管理员-修改用户密码
		apiV1.DELETE("/admin/user/:id", DeleteUser)
		// 管理员-管理员更新用户信息
		apiV1.PUT("/admin/user/:id", UpdateUser)
		// 管理员-查询所有权限组及其权限
		apiV1.GET("/admin/group", GetAdminGroups)

		apiV1.GET("/admin/group/:id", func(c *gin.Context) {
			if str := c.Param("id"); str == "all" {
				GetAllGroup(c) // 管理员-查询所有权限组
			}else{
				GetGroup(c)  // 管理员-查询一个权限组及其权限
			}
		})

		// 管理员-新建权限组
		apiV1.POST("/admin/group", CreateGroup)
		// 管理员-更新一个权限组
		apiV1.PUT("/admin/group/:id", UpdateGroup)
		// 管理员-删除一个权限组
		apiV1.DELETE("/admin/group/:id", DeleteGroup)
		// 管理员-分配单个权限
		apiV1.POST("/admin/permission/dispatch", DispatchPermission)
		// 管理员-分配多个权限
		apiV1.POST("/admin/permission/dispatch/batch", DispatchPermissions)
		// 管理员-删除多个权限
		apiV1.POST("/admin/permission/remove", RemovePermissions)


		// 日志-查询所有日志
		apiV1.GET("/log", GetLogs)
		// 日志-搜索日志
		apiV1.GET("/log/search", GetUserLogs)
		// 日志-查询日志记录的用户
		apiV1.GET("/log/users", GetUsers)
	}
}
