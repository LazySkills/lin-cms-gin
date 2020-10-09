/** Created By wene<354007048@qq.com> . Date at 2020/7/3 */
package permission

import (
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/middleware/jwt"
	"lin-cms-gin/internal/models"
	"lin-cms-gin/pkg/e"
	"lin-cms-gin/pkg/lin"
	"net/http"
)

var UserGroup models.LinUserGroup

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			appG = lin.Gin{C: c}
			code int = e.SUCCESS
		)

		 if jwt.Claims.UniqueId > 0 {
			 if jwt.Claims.UniqueId != 1 {
				 UserGroup = models.GetLinUserGroupByUserID(jwt.Claims.UniqueId)
			 }
		}else{
			code = e.ERROR_USER_TO_JSON
		}

		if code != e.SUCCESS {
			appG.ResponseError(http.StatusForbidden,code,nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func GroupRequired(method string, action string) bool {
	if jwt.Claims.UniqueId == 1 {
		return true
	}
	permission := lin.GetPermissionMappingByName(method,action)
	if permission.ID <= 0 {
		return false
	}
	if !models.ExistUserPermissionByPermissionId(permission.ID, UserGroup.GroupId)  {
		return false
	}
	return true
}