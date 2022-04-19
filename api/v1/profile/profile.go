package profile

import (
	"focus-single/api/v1/common"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/profile" method:"get" summary:"展示个人资料页面" tags:"个人"`
}
type IndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UpdateReq struct {
	g.Meta   `path:"/profile" method:"post" summary:"修改个人资料" tags:"个人"`
	Id       uint   `json:"id"     dc:"用户ID"`
	Avatar   string `json:"avatar" dc:"头像地址"`
	Gender   int    `json:"gender" dc:"性别 0: 未设置 1: 男 2: 女"`
	Nickname string `json:"nickname" v:"required#请输入昵称信息" dc:"昵称"`
}
type UpdateRes struct{}

type AvatarReq struct {
	g.Meta `path:"/profile/avatar" method:"get" summary:"展示头像管理页面" tags:"个人"`
}
type AvatarRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UpdateAvatarReq struct {
	g.Meta `path:"/profile/avatar" method:"post" summary:"修改个人头像" tags:"个人"`
	Id     uint   `json:"id"     dc:"用户ID"`
	Avatar string `json:"avatar" dc:"头像地址"`
}
type UpdateAvatarRes struct{}

type PasswordReq struct {
	g.Meta `path:"/profile/password" method:"get" summary:"展示密码修改页面" tags:"个人"`
}
type PasswordRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UpdatePasswordReq struct {
	g.Meta      `path:"/profile/password" method:"post" summary:"修改个人密码" tags:"个人"`
	OldPassword string `json:"oldPassword" v:"required#请输入原始密码" dc:"原密码"`
	NewPassword string `json:"newPassword" v:"required#请输入新密码"   dc:"新密码"`
}
type UpdatePasswordRes struct{}

type MessageReq struct {
	g.Meta     `path:"/profile/message" method:"get" summary:"展示查询用户消息列表页面" tags:"个人"`
	TargetType string `json:"targetType" dc:"数据类型"`
	TargetId   uint   `json:"targetId"   dc:"数据ID"`
	common.PaginationReq
}
type MessageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
