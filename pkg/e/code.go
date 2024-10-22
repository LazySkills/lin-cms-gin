/** Created By wene<354007048@qq.com> . Date at 2020/6/2 */
package e

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	// 用户权限错误
	AUTH_FAIL = 1001
	// 用户权限错误
	AUTH_EMPTY = 1002

	ERROR_EXIST_BOOK = 10001
	ERROR_NOT_EXIST_BOOK = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN = 20003
	ERROR_AUTH = 20004
	ERROR_AUTH_TYPE = 20005



	// 保存图片失败
	ERROR_UPLOAD_SAVE_IMAGE_FAIL = 30001
	// 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL = 30002
	// 校验图片错误，图片格式或大小有问题
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003

	// 用户转json失败
	ERROR_USER_TO_JSON = 40001

	// 用户已存在
	ERROR_EXIST_USER = 40002
	// 用户密码PBKDF2加密错误
	ERROR_USER_PASSWORD_PBKDF2_FAIL = 40003
	// 更新用户失败
	ERROR_UPDATE_USER_FAIL = 40004
	// 用户密码错误
	ERROR_USER_PASSWORD_FAIL = 40005
	// 用户不存在
	ERROR_NOT_EXIST_USER = 40006


	// 分组不存在
	ERROR_NOT_EXIST_GROUP = 50000
	// 不能删除root分组
	ERROR_ROOT_GROUP_DELETE = 50001
	ERROR_GUEST_GROUP_DELETE = 50002
	// 删除分组失败
	DELETE_FAIL_GROUP = 50003

)