package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 用户注册
type UserRegisterReq struct {
	g.Meta   `method:"post" summary:"执行用户注册" tags:"用户"`
	Passport string `json:"passport" v:"required#请输入账号" dc:"账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
	Nickname string `json:"nickname" v:"required#请输入昵称" dc:"昵称"`
	Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}
type UserRegisterRes struct{}

// 修改用户密码
type UserUpdatePasswordReq struct {
	g.Meta      `method:"post" summary:"修改个人密码" tags:"用户"`
	OldPassword string `json:"oldPassword" v:"required#请输入原始密码" dc:"原密码"`
	NewPassword string `json:"newPassword" v:"required#请输入新密码"   dc:"新密码"`
}
type UserUpdatePasswordRes struct{}

// 修改用户
type UserUpdateProfileReq struct {
	g.Meta   `method:"post" summary:"修改个人资料" tags:"用户"`
	Id       uint   `json:"id"     dc:"用户ID"`
	Avatar   string `json:"avatar" dc:"头像地址"`
	Gender   int    `json:"gender" dc:"性别 0: 未设置 1: 男 2: 女"`
	Nickname string `json:"nickname" v:"required#请输入昵称信息" dc:"昵称"`
}
type UserUpdateProfileRes struct{}

// 禁用用户
type UserDisableReq struct {
	g.Meta `method:"post" summary:"禁用指定用户(测试)" tags:"用户"`
	Id     uint `json:"id" v:"min:1#请选择需要禁用的用户" dc:"删除时ID不能为空"`
}
type UserDisableRes struct{}

// 用户登录
type UserLoginReq struct {
	g.Meta   `method:"post" summary:"执行用户登录" tags:"用户"`
	Passport string `json:"passport" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
	Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}
type UserLoginRes struct {
	Referer string `json:"referer" dc:"引导客户端跳转地址"`
}

// 查询用户列表请求
type UserGetContentListReq struct {
	g.Meta `method:"get" summary:"展示查询用户内容列表页面" tags:"用户"`
	ContentGetListReq
	UserId uint `json:"userId" in:"query" dc:"要查询的用户ID"`
}
type UserGetContentListRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 查询用户列表查询请求
type UserGetMessageListReq struct {
	g.Meta `method:"get" summary:"展示查询用户消息列表页面" tags:"用户"`
	CommonListReq
	TargetType string `json:"targetType" dc:"数据类型"`
	TargetId   uint   `json:"targetId"   dc:"数据ID"`
}
type UserGetMessageListRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 用户个人资料页面
type UserProfileReq struct {
	g.Meta `method:"get" summary:"展示用户个人资料页面" tags:"用户"`
}
type UserProfileRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 用户头像管理页面
type UserAvatarReq struct {
	g.Meta `method:"get" summary:"展示个人头像管理页面" tags:"用户"`
}
type UserAvatarRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 用户密码管理页面
type UserPasswordReq struct {
	g.Meta `method:"get" summary:"展示个人密码修改页面" tags:"用户"`
}
type UserPasswordRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 用户注销
type UserLogoutReq struct {
	g.Meta `method:"post" summary:"执行用户注销接口" tags:"用户"`
}
type UserLogoutRes struct{}
