/** Created By 嗝嗝<354007048@qq.com>. Date 2020/9/17 */
package lin

import (
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/pkg/e"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	var appG = Gin{C: c}
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			appG.ResponseError(http.StatusBadRequest,e.ERROR, errorToString(r))
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}