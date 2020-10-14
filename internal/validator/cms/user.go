package cms


// 添加用户验证器
type AddUserValidator struct {
	Username string `form:"username" binding:"required,max=50"`
	Password string `form:"password" binding:"required,max=50"`
}

// 更新用户验证器
type UpdateUserValidator struct {
	Username         string `form:"username" binding:"required,max=24"`
	Nickname          string `form:"nickname" binding:"required,max=24"`
	Email       string `form:"email" binding:"required,max=100"`
	Avatar string `form:"avatar" binding:"required;max=500"`
}



// 更新用户密码验证器
type UpdateUserPasswordValidator struct {
	OldPassword string `form:"old_password" binding:"required,max=50"`
	NewPassword string `form:"new_password" binding:"required,max=50"`
}
