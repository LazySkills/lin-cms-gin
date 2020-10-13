/** Created By 嗝嗝<354007048@qq.com>. Date 2020/10/13 */
package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"lin-cms-gin/internal/middleware/permission"
	"lin-cms-gin/internal/models"
	"lin-cms-gin/internal/validator/cms"
	"lin-cms-gin/pkg/e"
	"lin-cms-gin/pkg/lin"
	"lin-cms-gin/pkg/setting"
	"lin-cms-gin/pkg/tools"
	"net/http"
)

type LinLogPermission struct {
	lin.Permission
}

func (p *LinLogPermission) AuthMapping()  {
	p.Mapping("GetLogs","GET","查询所有日志","管理员",1)
	p.Mapping("GetUserLogs","GET","搜索日志","管理员",1)
	p.Mapping("GetUsers","GET","查询日志记录的用户","管理员",1)
}

// @Summary 查询所有日志
// @Tags 日志
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{....}"
// @Router /cms/log [get]
func GetLogs(c *gin.Context)  {
	var (
		appG = lin.Gin{C: c}
		req = &cms.LogFindValidator{}
		maps = make(map[string]interface{})
		data = make(map[string]interface{})
	)

	if !permission.GroupRequired(c.Request.Method,"GetLogs") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	limit := setting.AppSetting.PageSize
	if arg := appG.C.Query("count"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}

	data["items"] = models.GetLinLog(tools.GetPage(c), limit, maps)
	data["total"] = models.GetLinLogTotal(maps)
	data["count"] = limit
	data["page"] = tools.GetPage(c)
	appG.Response(http.StatusOK, data)
}


// @Summary 搜索日志
// @Tags 日志
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{....}"
// @Router /cms/log/search [get]
func GetUserLogs(c *gin.Context)  {
	var (
		appG = lin.Gin{C: c}
		req = &cms.LogFindValidator{}
		maps = make(map[string]interface{})
		data = make(map[string]interface{})
	)

	if !permission.GroupRequired(c.Request.Method,"GetUserLogs") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	limit := setting.AppSetting.PageSize
	if arg := appG.C.Query("count"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}

	maps["message"] = appG.C.DefaultQuery("keyword","")

	data["items"] = models.GetLinLog(tools.GetPage(c), limit, maps)
	data["total"] = models.GetLinLogTotal(maps)
	data["count"] = limit
	data["page"] = tools.GetPage(c)
	appG.Response(http.StatusOK, data)
}

// @Summary 查询日志记录的用户
// @Tags 日志
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{....}"
// @Router /cms/log/users [get]
func GetUsers(c *gin.Context)  {
	var (
		appG = lin.Gin{C: c}
		req = &cms.LogFindValidator{}
		data = make(map[string]interface{})
	)

	if !permission.GroupRequired(c.Request.Method,"GetUsers") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	limit := setting.AppSetting.PageSize
	if arg := appG.C.Query("count"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}

	data["items"] = models.GetLinLogUsers(tools.GetPage(c), limit)
	data["total"] = models.GetLinLogUsersTotal()
	data["count"] = limit
	data["page"] = tools.GetPage(c)
	appG.Response(http.StatusOK, data)
}
