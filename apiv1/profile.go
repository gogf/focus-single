package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProfileIndexReq struct {
	g.Meta `path:"/profile" method:"get" summary:"展示个人资料页面" tags:"个人"`
}
type ProfileIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ProfileUpdateReq struct {
	g.Meta   `path:"/profile" method:"post" summary:"修改个人资料" tags:"个人"`
	Id       uint   `json:"id"     dc:"用户ID"`
	Avatar   string `json:"avatar" dc:"头像地址"`
	Gender   int    `json:"gender" dc:"性别 0: 未设置 1: 男 2: 女"`
	Nickname string `json:"nickname" v:"required#请输入昵称信息" dc:"昵称"`
}
type ProfileUpdateRes struct{}

type ProfileAvatarReq struct {
	g.Meta `path:"/profile/avatar" method:"get" summary:"展示头像管理页面" tags:"个人"`
}
type ProfileAvatarRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ProfileUpdateAvatarReq struct {
	g.Meta `path:"/profile/avatar" method:"post" summary:"修改个人头像" tags:"个人"`
	Id     uint   `json:"id"     dc:"用户ID"`
	Avatar string `json:"avatar" dc:"头像地址"`
}
type ProfileUpdateAvatarRes struct{}

type ProfilePasswordReq struct {
	g.Meta `path:"/profile/password" method:"get" summary:"展示密码修改页面" tags:"个人"`
}
type ProfilePasswordRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ProfileUpdatePasswordReq struct {
	g.Meta      `path:"/profile/password" method:"post" summary:"修改个人密码" tags:"个人"`
	OldPassword string `json:"oldPassword" v:"required#请输入原始密码" dc:"原密码"`
	NewPassword string `json:"newPassword" v:"required#请输入新密码"   dc:"新密码"`
}
type ProfileUpdatePasswordRes struct{}

type ProfileMessageReq struct {
	g.Meta `path:"/profile/message" method:"get" summary:"展示查询用户消息列表页面" tags:"个人"`
	CommonPaginationReq
	TargetType string `json:"targetType" dc:"数据类型"`
	TargetId   uint   `json:"targetId"   dc:"数据ID"`
}
type ProfileMessageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
