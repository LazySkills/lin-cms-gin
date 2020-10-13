/** Created By wene<354007048@qq.com> . Date at 2020/7/5 */
package admin

import (
	gin "github.com/gin-gonic/gin"
	"lin-cms-gin/internal/validator/cms"
	"lin-cms-gin/internal/middleware/jwt"
	"lin-cms-gin/internal/middleware/permission"
	"lin-cms-gin/internal/models"
	"lin-cms-gin/pkg/e"
	"lin-cms-gin/pkg/lin"
	unit "lin-cms-gin/pkg/tools"
	"net/http"
	"strings"
	"time"
)

type LinUserPermission struct {
	lin.Permission
}

func (p *LinUserPermission) AuthMapping()  {
	p.Mapping("UserLogin","POST","登录","用户",0)
	p.Mapping("UserGetToken","GET","刷新令牌","用户",0)
	p.Mapping("UserPermissions","GET","查询自己拥有的权限","用户",1)
	p.Mapping("UserRegister","POST","用户注册","用户",1)
	p.Mapping("UserUpdate","PUT","用户更新信息","用户",1)
	p.Mapping("UserUpdatePassword","PUT","修改密码","用户",1)
	p.Mapping("UserGetInreqation","GET","查询自己信息","用户",1)
}



// @Summary 用户注册
// @Tags 用户
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} json "{"code":200,"data":{"access_token":"...","refresh_token":"..."},"msg":"ok"}"
// @Router /cms/user/register [post]
func UserRegister(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		req = &cms.AddUserValidator{}
	)

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	if ! models.ExistLinUserByUsername(req.Username) {
		_,err := models.AddLinAuth(req.Username, req.Password)
		if err != nil {
			appG.ResponseError(http.StatusBadRequest, e.ERROR_USER_PASSWORD_PBKDF2_FAIL,nil)
			return
		}
	}

	appG.Response(http.StatusOK, nil)
}



// @Summary 用户更新信息
// @Tags 用户
// @Produce  json
// @Param ID query string true "用户名"
// @Param username query string true "用户名"
// @Param nickname query string true "昵称"
// @Param email query string true "邮箱"
// @Param avatar query string true "头像"
// @Success 200 {string} json "{"code":200,"data":{"access_token":"...","refresh_token":"..."},"msg":"ok"}"
// @Router /cms/user [put]
func UserUpdate(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		req = &cms.UpdateUserValidator{}
	)

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	if !models.UpdateUserInfo(permission.UserGroup.ID,req.Username,req.Nickname,req.Email,req.Avatar) {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_UPDATE_USER_FAIL,nil)
		return
	}

	appG.ResponseSuccess(http.StatusOK, e.SUCCESS, nil)
}



// @Summary 修改密码
// @Tags 用户
// @Produce  json
// @Param old_password query string true "旧密码"
// @Param new_password query string true "新密码"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /cms/user/password [put]
func UserUpdatePassword(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		req = &cms.UpdateUserPasswordValidator{}
	)

	if !permission.GroupRequired(c.Request.Method,"UserUpdatePassword") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	if !models.UpdateIdentityPasswordByUserID(jwt.Claims.UniqueId,req.OldPassword,req.NewPassword) {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_UPDATE_USER_FAIL,nil)
		return
	}

	appG.ResponseSuccess(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 登录
// @Tags 用户
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} json "{"access_token":"...","refresh_token":"..."}"
// @Router /cms/user/login [post]
func UserLogin(c *gin.Context) {
	var (
		appG  = lin.Gin{C: c}
		req  = &cms.AddUserValidator{}
		errCode = e.SUCCESS
		AuthM models.LinUserIdentity
	)

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	data := make(map[string]interface{})
	AuthM = models.ExistIdentityByName(req.Username)
	if AuthM.ID > 0 {
		if !models.CheckIdentityPassword(AuthM, req.Password) {
			errCode = e.ERROR_USER_PASSWORD_FAIL
		}

		token, err := unit.GenerateToken(AuthM.ID,false)
		if err != nil {
			errCode = e.ERROR_AUTH_TOKEN
		} else {
			data["access_token"] = token
			errCode = e.SUCCESS
		}

		refreshToken, err := unit.GenerateToken(AuthM.ID,true)
		if err != nil {
			errCode = e.ERROR_AUTH_TOKEN
		} else {
			data["refresh_token"] = refreshToken
			errCode = e.SUCCESS
		}
	} else {
		errCode = e.ERROR_AUTH
	}

	if errCode == 200 {
		appG.Response(http.StatusOK, data)
	}else {
		appG.ResponseError(http.StatusBadRequest, errCode, nil)
	}
}

// @Summary 刷新令牌
// @Tags 用户
// @Produce  json
// @Param Authorization header string true "刷新token"
// @Success 200 {string} json "{"access_token":"..."}"
// @Router /cms/user/refresh [get]
func UserGetToken(c *gin.Context) {
	var code int
	var data = make(map[string]interface{})
	var appG = lin.Gin{C: c}

	code = e.SUCCESS
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		code = e.INVALID_PARAMS
	} else {
		authHeaderMap := strings.SplitN(authHeader, " ", 2)
		if !(len(authHeaderMap) == 2 && authHeaderMap[0] == "Bearer") {
			appG.ResponseError(http.StatusBadRequest, e.ERROR_AUTH, nil)
			c.Abort()
			return
		}
		claims, err := unit.ParseToken(authHeaderMap[1], true)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if  time.Now().Unix()>claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}

		token, err := unit.GenerateToken(claims.UniqueId,false)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["access_token"] = token
		}
	}

	if code != e.SUCCESS {
		appG.ResponseError(http.StatusBadRequest, code, nil)
		c.Abort()
		return
	}

	appG.Response(http.StatusOK,data)
}


// @Summary 查询自己拥有的权限
// @Tags 用户
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{....}"
// @Router /cms/user/permissions [get]
func UserGetPermissions(c *gin.Context)  {
	var (
		appG     = lin.Gin{C: c}
		httpCode = http.StatusOK
		User     models.LinUser
		data     = make(map[string]interface{})
	)

	if !permission.GroupRequired(c.Request.Method,"UserGetPermissions") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	User = models.GetLinUserByID(jwt.Claims.UniqueId)
	data["id"] = User.ID
	data["nickname"] = User.Nickname
	data["email"] = User.Email
	data["avatar"] = User.Avatar
	data["admin"] = false
	if User.ID == 1 {
		data["admin"] = true
	}

	if permission.UserGroup.GroupId >0 {
		uP := models.GetLinGroupPermissionByGroupId(permission.UserGroup.GroupId)
		data["permissions"] = lin.FormatLinGroupPermission(uP)
	}else {
		data["permissions"] = lin.FormatLinPermission(models.GetAllLinPermission())
	}
	appG.Response(httpCode,data)

}

// @Summary 查询用户信息
// @Tags 用户
// @Produce  json
// @Param Authorization header string true "授权token"
// @Success 200 {string} json "{....}"
// @Router /cms/user/information [get]
func UserGetInFormation(c *gin.Context)  {
	var (
		appG     = lin.Gin{C: c}
		httpCode = http.StatusOK
		user     models.LinUser
	)

	if !permission.GroupRequired(c.Request.Method,"UserGetInFormation") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	if jwt.Claims.UniqueId >0  {
		user = models.GetLinUserByID(jwt.Claims.UniqueId)
	}else {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_NOT_EXIST_USER,nil)
		c.Abort()
		return
	}

	appG.Response(httpCode, user)
}