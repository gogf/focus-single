package api

import (
	"focus/app/internal/model"
	"github.com/gogf/gf/frame/g"
)

// 用户注册
type UserRegisterReq struct {
	g.Meta `method:"post" summary:"执行用户注册" tags:"用户"`
	model.UserRegisterInput
	Passport string `v:"required#请输入账号"`  // 账号
	Password string `v:"required#请输入密码"`  // 密码
	Nickname string `v:"required#请输入昵称"`  // 昵称
	Captcha  string `v:"required#请输入验证码"` // 验证码
}
type UserRegisterRes struct{}

// 修改用户密码
type UserUpdatePasswordReq struct {
	g.Meta `method:"post" summary:"修改个人密码" tags:"用户"`
	model.UserPasswordInput
	OldPassword string `v:"required#请输入原始密码"` // 原密码
	NewPassword string `v:"required#请输入新密码"`  // 新密码
}
type UserUpdatePasswordRes struct{}

// 修改用户
type UserUpdateProfileReq struct {
	g.Meta `method:"post" summary:"修改个人资料" tags:"用户"`
	model.UserUpdateProfileInput
	Nickname string `v:"required#请输入昵称信息"` // 昵称
}
type UserUpdateProfileRes struct{}

// 禁用用户
type UserDisableReq struct {
	g.Meta `method:"post" summary:"禁用指定用户(测试)" tags:"用户"`
	Id     *uint `v:"required#请选择需要禁用的用户"` // 删除时ID不能为空
}
type UserDisableRes struct{}

// 用户登录
type UserLoginReq struct {
	g.Meta `method:"post" summary:"执行用户登录" tags:"用户"`
	model.UserLoginInput
	Passport string `v:"required#请输入账号"`  // 账号
	Password string `v:"required#请输入密码"`  // 密码(明文)
	Captcha  string `v:"required#请输入验证码"` // 验证码
}
type UserLoginRes struct {
	Referer string
}

// 查询用户列表请求
type UserGetContentListReq struct {
	g.Meta `method:"get" summary:"展示查询用户内容列表页面" tags:"用户"`
	model.UserGetContentListInput
	Id uint `json:"id" in:"path" description:"内容id"`
}
type UserGetContentListRes struct {
	g.Meta `mime:"text/html"`
}

// 查询用户列表查询请求
type UserGetMessageListReq struct {
	g.Meta `method:"get" summary:"展示查询用户消息列表页面" tags:"用户"`
	model.UserGetMessageListInput
}
type UserGetMessageListRes struct {
	g.Meta `mime:"text/html"`
}

// 用户个人资料页面
type UserProfileReq struct {
	g.Meta `method:"get" summary:"展示用户个人资料页面" tags:"用户"`
}
type UserProfileRes struct {
	g.Meta `mime:"text/html"`
}

// 用户头像管理页面
type UserAvatarReq struct {
	g.Meta `method:"get" summary:"展示个人头像管理页面" tags:"用户"`
}
type UserAvatarRes struct {
	g.Meta `mime:"text/html"`
}

// 用户密码管理页面
type UserPasswordReq struct {
	g.Meta `method:"get" summary:"展示个人密码修改页面" tags:"用户"`
}
type UserPasswordRes struct {
	g.Meta `mime:"text/html"`
}

// 用户注销
type UserLogoutReq struct {
	g.Meta `method:"post" summary:"执行用户注销接口" tags:"用户"`
}
type UserLogoutRes struct{}
