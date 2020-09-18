/** Created By wene<354007048@qq.com> . Date at 2020/6/2 */
package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	AUTH_FAIL:             			"用户权限错误",
	AUTH_EMPTY:             		"用户权限不存在",
	ERROR_EXIST_BOOK:               "已存在该标签名称",
	ERROR_NOT_EXIST_BOOK:           "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH_TYPE:               "Token类型错误",
	ERROR_AUTH:                     "Token错误",
	ERROR_USER_TO_JSON:             "用户信息转Json错误",
	ERROR_EXIST_USER: 				"用户已存在",
	ERROR_USER_PASSWORD_PBKDF2_FAIL: 	"用户密码PBKDF2加密错误",
	ERROR_USER_PASSWORD_FAIL: 	"用户密码错误",
	ERROR_UPDATE_USER_FAIL: 	"更新用户失败",
	ERROR_NOT_EXIST_USER: 	"用户不存在",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL: "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL: "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
