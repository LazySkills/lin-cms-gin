/** Created By 嗝嗝<354007048@qq.com>. Date 2020/10/13 */
package lin

import (
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/models"
	"runtime"
)

type LoggerStruct struct {
	Message string
	UserId int
	Username string
	StatusCode int
	Method string
	Path string
	Permission string
}

func Logger(c gin.Context,message string)  {
	//user := models.GetLinUserByID(jwt.Claims.UniqueId)

	models.AddLinLog(&LoggerStruct{
		Message: message,
		//UserId: jwt.Claims.UniqueId,
		//Username: user.Username,
		StatusCode: 0,
		Method: c.Request.Method,
		Path: c.Request.RequestURI,
		//Permission: GetPermissionMappingByName(c.Request.Method,""),
	})
}

func LoggerAction() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func LoggerCallAction() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
