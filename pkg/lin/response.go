/** Created By wene<354007048@qq.com> . Date at 2020/6/21 */
package lin

import (
	"github.com/gin-gonic/gin"
	"lin-cms-gin/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Request string 	 `json:"request"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, data interface{}) {
	g.C.JSON(httpCode, data)
	return
}

// Response setting gin.JSON
func (g *Gin) ResponseSuccess(httpCode int, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

func (g *Gin) ResponseError(httpCode int, errCode int, data interface{}) {
	g.C.JSON(httpCode, ResponseError{
		Code:    errCode,
		Msg:     e.GetMsg(errCode),
		Request: g.C.Request.Method+" "+g.C.Request.RequestURI,
		Data:    data,
	})
	return
}
