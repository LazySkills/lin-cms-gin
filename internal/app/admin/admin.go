/** Created By wene<354007048@qq.com> . Date at 2020/7/5 */
package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"lin-cms-gin/internal/dto/cms"
	"lin-cms-gin/internal/middleware/permission"
	"lin-cms-gin/internal/models"
	"lin-cms-gin/pkg/e"
	"lin-cms-gin/pkg/lin"
	"lin-cms-gin/pkg/setting"
	"lin-cms-gin/pkg/tools"
	"net/http"
)

type LinAdminPermission struct {
	lin.Permission
}

func (p *LinAdminPermission) AuthMapping()  {
	p.Mapping("GetAllPermissions","GET","查询所有可分配的权限","管理员",1)
	p.Mapping("GetAdminUsers","GET","查询所有用户","管理员",1)
	p.Mapping("ChangeUserPassword","PUT","修改用户密码","管理员",1)
	p.Mapping("DeleteUser","DELETE","删除用户","管理员",1)
	p.Mapping("UpdateUser","PUT","管理员更新用户信息","管理员",1)
	p.Mapping("GetAdminGroups","GET","查询所有权限组及其权限","管理员",1)
	p.Mapping("GetAllGroup","GET","查询所有权限组","管理员",1)
	p.Mapping("GetGroup","GET","查询一个权限组及其权限","管理员",1)
	p.Mapping("CreateGroup","POST","新建权限组","管理员",1)
	p.Mapping("UpdateGroup","PUT","更新一个权限组","管理员",1)
	p.Mapping("DeleteGroup","DELETE","删除一个权限组","管理员",1)
	p.Mapping("DispatchPermission","POST","分配单个权限","管理员",1)
	p.Mapping("DispatchPermissions","POST","分配多个权限","管理员",1)
	p.Mapping("RemovePermissions","POST","删除多个权限","管理员",1)
	p.Mapping("UpdateGroup","PUT","更新一个权限组","管理员",1)
}



// @Summary 查询所有可分配的权限
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{....}"
// @Router /cms/admin/permission [get]
func GetAllPermissions(c *gin.Context)  {
	var (
		appG = lin.Gin{C: c}
		httpCode = http.StatusOK
		linPermission = make(map[string][]lin.LinPermission)
	)

	if !permission.GroupRequired(c.Request.Method,"GetAllPermissions") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	fmt.Printf("permission：%v \n",permission.UserGroup)
	if permission.UserGroup.GroupId > 0 {
		linPermission = lin.FormatLinGroupPermission(models.GetLinGroupPermissionByGroupId(permission.UserGroup.GroupId))
	}else {
		linPermission =  lin.FormatLinPermission(models.GetAllLinPermission())
	}

	appG.Response(httpCode, linPermission)
}

// @Summary 查询所有用户
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{....}"
// @Router /cms/admin/users [get]
func GetAdminUsers(c *gin.Context)  {
	var (
		appG = lin.Gin{C: c}
		httpCode = http.StatusOK
		maps = make(map[string]interface{})
		data = make(map[string]interface{})
	)

	if !permission.GroupRequired(c.Request.Method,"UserGetPermissions") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	limit := setting.AppSetting.PageSize
	if arg := appG.C.Query("count"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}

	maps["user_id <>"] = 1   // 排除管理员

	if arg := appG.C.Query("group_id"); arg != "" {
		var linUsersIds []int
		linUserIds := models.GetLinUserGroupUserIds(com.StrTo(arg).MustInt(),maps)
		for _,v := range linUserIds {
			linUsersIds = append(linUsersIds,v.UserId)
		}
		maps["user_id in"] = linUsersIds
	}

	linUsers := models.GetLinUser(tools.GetPage(c), limit, maps)
	data["total"] = models.GetLinUserTotal(maps)
	data["items"] = linUsers
	data["count"] = limit
	data["page"] = tools.GetPage(c)

	appG.Response(httpCode, data)
}


// @Summary 修改用户密码
// @Tags 用户
// @Produce  json
// @Param Authorization header string true "授权token"
// @Param new_password query string true "新密码"
// @Param confirm_password query string true "确认密码"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /cms/admin/user/:id/password [put]
func ChangeUserPassword(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		req = &cms.ChangeUserPasswordForm{}
	)

	if !permission.GroupRequired(c.Request.Method,"ChangeUserPassword") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	id := com.StrTo(c.Param("id")).MustInt()

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	if !models.UpdateIdentityCredentialByUserID(id,req.NewPassword) {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_UPDATE_USER_FAIL,nil)
		return
	}

	appG.ResponseSuccess(http.StatusOK, e.SUCCESS,nil)
}


// @Summary 删除用户
// @Tags 用户
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /cms/admin/user/:id [delete]
func DeleteUser(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		httpCode = http.StatusOK
		req = &cms.ChangeUserPasswordForm{}
	)

	if !permission.GroupRequired(c.Request.Method,"DeleteUser") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	id := com.StrTo(c.Param("id")).MustInt()

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	if !models.DeleteLinUser(id) {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_UPDATE_USER_FAIL,nil)
		return
	}

	appG.ResponseSuccess(httpCode, e.SUCCESS,nil)
}


// @Summary 管理员更新用户信息
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Param ID query string true "用户名"
// @Param username query string true "用户名"
// @Param nickname query string true "昵称"
// @Param email query string true "邮箱"
// @Param avatar query string true "头像"
// @Success 200 {string} json "{"code":200,"data":{"access_token":"...","refresh_token":"..."},"msg":"ok"}"
// @Router /cms/admin/user/:id [put]
func UpdateUser(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		req = &cms.UpdateUserForm{}
	)

	if !permission.GroupRequired(c.Request.Method,"UpdateUser") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	id := com.StrTo(c.Param("id")).MustInt()

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	if !models.UpdateUserInfo(id,req.Username,req.Nickname,req.Email,req.Avatar) {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_UPDATE_USER_FAIL,nil)
		return
	}

	appG.ResponseSuccess(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 查询所有权限组及其权限
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{"code":200,"data":{"access_token":"...","refresh_token":"..."},"msg":"ok"}"
// @Router /cms/admin/group [get]
func GetAdminGroups(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		maps = make(map[string]interface{})
		data = make(map[string]interface{})
	)

	if !permission.GroupRequired(c.Request.Method,"GetAdminGroups") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	limit := setting.AppSetting.PageSize
	if arg := appG.C.Query("count"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}

	data["items"] = models.GetLinGroup(tools.GetPage(c), limit, maps)
	data["total"] = models.GetLinGroupTotal(maps)
	data["count"] = limit
	data["page"] = tools.GetPage(c)

	appG.Response(http.StatusOK, data)
}


// @Summary 查询所有权限组
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "[{"id":2,"name":"guest","info":"游客组"}]"
// @Router /cms/admin/all [get]
func GetAllGroup(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		maps = make(map[string]interface{})
		data []models.LinGroupJson
	)

	if !permission.GroupRequired(c.Request.Method,"GetAllGroup") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	maps["level"] = 1
	limit := models.GetLinGroupTotal(maps)
	data = models.GetLinGroup(tools.GetPage(c), int(limit), maps)
	appG.Response(http.StatusOK, data)
}

// @Summary 查询一个权限组及其权限
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "[{"id":2,"name":"guest","info":"游客组"}]"
// @Router /cms/admin/group/:id [get]
func GetGroup(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
	)

	if !permission.GroupRequired(c.Request.Method,"GetGroup") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}
	id := com.StrTo(c.Param("id")).MustInt()

	linPermission := lin.FormatLinGroupPermission(models.GetLinGroupPermissionByGroupId(id))

	appG.Response(http.StatusOK, linPermission)
}

// @Summary 新建权限组
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Param name query string true "组名"
// @Param info query string true "备注"
// @Param permission_ids query array true "权限ID集合"
// @Success 200 {string} json "[{"id":2,"name":"guest","info":"游客组"}]"
// @Router /cms/admin/group [post]
func CreateGroup(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		req = &cms.NewGroupForm{}
	)

	if !permission.GroupRequired(c.Request.Method,"CreateGroup") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	linGourp := models.AddLinGroup(req.Name,req.Info,3)
	for _,k := range req.PermissionIds {
		models.AddLinGroupPermission(linGourp.ID, k)
	}

	appG.ResponseSuccess(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 更新一个权限组
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Param name query string true "组名"
// @Param info query string true "备注"
// @Success 200 {string} json "[{"id":2,"name":"guest","info":"游客组"}]"
// @Router /cms/admin/group/:id [put]
func UpdateGroup(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		req = &cms.UpdateGroupForm{}
	)

	if !permission.GroupRequired(c.Request.Method,"UpdateGroup") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	linGroup := models.GetLinGroupByID(com.StrTo(c.Param("id")).MustInt())
	if linGroup.ID > 0 {
		linGroup.UpdateLinGroup(req.Name,req.Info)
	}else {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_NOT_EXIST_GROUP,nil)
		return
	}

	appG.ResponseSuccess(http.StatusOK, e.SUCCESS, nil)
}


// @Summary 更新一个权限组
// @Tags 管理员
// @Produce  json
// @Param Authorization header string true "授权token"
// @Param name query string true "组名"
// @Param info query string true "备注"
// @Success 200 {string} json "[{"id":2,"name":"guest","info":"游客组"}]"
// @Router /cms/admin/group/:id [put]
func DeleteGroup(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
	)

	if !permission.GroupRequired(c.Request.Method,"DeleteGroup") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	linGroup := models.GetLinGroupByID(com.StrTo(c.Param("id")).MustInt())
	if linGroup.ID < 1 {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_NOT_EXIST_GROUP,nil)
		return
	}

	if linGroup.Level == setting.LinSetting.GroupLevelRoot {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_ROOT_GROUP_DELETE,nil)
		return
	}else if linGroup.Level == setting.LinSetting.GroupLevelGuest  {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_GUEST_GROUP_DELETE,nil)
		return
	}

	appG.ResponseSuccess(http.StatusOK, e.SUCCESS, nil)
}
