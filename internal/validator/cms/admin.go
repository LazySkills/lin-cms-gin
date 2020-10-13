/** Created By 嗝嗝<354007048@qq.com>. Date 2020/7/30 */
package cms

// 更新用户密码验证器
type ChangeUserPasswordValidator struct {
	ConfirmPassword string `form:"confirm_password" valid:"Required; MaxSize(50)"`
	NewPassword string `form:"new_password" valid:"Required; MaxSize(50)"`
}

// 创建权限组验证
type NewGroupValidator struct {
	Name string `form:"name" json:"name" binding:"required"`
	Info string `form:"info" json:"info" binding:"required"`
	PermissionIds []int `form:"permission_ids" json:"permission_ids" binding:"required"`
}

// 创建权限组验证
type UpdateGroupValidator struct {
	Name string `form:"name" json:"name" binding:"required"`
	Info string `form:"info" json:"info" binding:"required"`
}

// 创建权限组验证
type DispatchPermissionValidator struct {
	GroupId int `form:"group_id" json:"group_id" binding:"required,numeric"`
	PermissionId int `form:"permission_id" json:"permission_id" binding:"required,numeric"`
}


// 创建权限组验证
type DispatchPermissionValidators struct {
	GroupId int `form:"group_id" json:"group_id" binding:"required,numeric"`
	PermissionIds []int `form:"permission_ids" json:"permission_ids" binding:"required"`
}