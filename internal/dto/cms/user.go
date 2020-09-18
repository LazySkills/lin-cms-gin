package cms


// 添加用户验证器
type AddUserForm struct {
	Username string `form:"username" valid:"Required; MaxSize(50)"`
	Password string `form:"password" valid:"Required; MaxSize(50)"`
}

// 更新用户验证器
type UpdateUserForm struct {
	Username         string `form:"username" valid:"Required;MaxSize(24)"`
	Nickname          string `form:"nickname" valid:"Required;MaxSize(24)"`
	Email       string `form:"email" valid:"Required;MaxSize(100)"`
	Avatar string `form:"avatar" valid:"Required;MaxSize(500)"`
}



// 更新用户密码验证器
type UpdateUserPasswordForm struct {
	OldPassword string `form:"old_password" valid:"Required; MaxSize(50)"`
	NewPassword string `form:"new_password" valid:"Required; MaxSize(50)"`
}