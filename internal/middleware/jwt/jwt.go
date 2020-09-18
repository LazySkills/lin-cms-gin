/** Created By wene<354007048@qq.com> . Date at 2020/6/2 */
package jwt

import (
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/pkg/e"
	"lin-cms-gin/internal/pkg/lin"
	"lin-cms-gin/internal/pkg/tools"
	"net/http"
	"strings"
	"time"
)

var Claims *tools.Claims

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			appG = lin.Gin{C: c}
			code int = e.SUCCESS
			err error
		)

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			code = e.AUTH_EMPTY
		} else {
			authHeaderMap := strings.SplitN(authHeader, " ", 2)
			if !(len(authHeaderMap) == 2 && authHeaderMap[0] == "Bearer") {
				appG.ResponseError(http.StatusOK,e.ERROR_AUTH_TYPE,nil)
				c.Abort()
				return
			}

			Claims, err = tools.ParseToken(authHeaderMap[1],false)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > Claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

		}

		if code != e.SUCCESS {
			appG.ResponseError(http.StatusUnauthorized,code,nil)
			c.Abort()
			return
		}

		c.Next()
	}
}