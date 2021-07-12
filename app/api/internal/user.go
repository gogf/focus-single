package internal

import "focus/app/model"

// 用户注册
type UserRegisterReq struct {
	model.UserRegisterInput
	Passport string `v:"required#请输入账号"`  // 账号
	Password string `v:"required#请输入密码"`  // 密码
	Nickname string `v:"required#请输入昵称"`  // 昵称
	Captcha  string `v:"required#请输入验证码"` // 验证码
}

// 修改用户密码
type UserPasswordReq struct {
	model.UserPasswordInput
	OldPassword string `v:"required#请输入原始密码"` // 原密码
	NewPassword string `v:"required#请输入新密码"`  // 新密码
}

// 修改用户
type UserUpdateProfileReq struct {
	model.UserUpdateProfileInput
	Nickname string `v:"required#请输入昵称信息"` // 昵称
}

// 禁用用户
type UserDisableReq struct {
	Id *uint `v:"required#请选择需要禁用的用户"` // 删除时ID不能为空
}

// Api用户登录
type UserLoginReq struct {
	model.UserLoginInput
	Passport string `v:"required#请输入账号"`  // 账号
	Password string `v:"required#请输入密码"`  // 密码(明文)
	Captcha  string `v:"required#请输入验证码"` // 验证码
}

// 修改用户
type UserServiceUpdateProfileReq struct {
	Nickname string // 昵称
	Gender   int    // 性别 0: 未设置 1: 男 2: 女
}

// 查询用户列表请求
type UserGetListReq struct {
	model.UserGetListInput
	Id uint
}

// 查询用户列表查询请求
type UserGetMessageListReq struct {
	model.UserGetMessageListInput
}
