package internal

import (
	"focus/app/model"
)

type ReplyDoCreateReq struct {
	model.ReplyCreateInput
	ParentId   uint   `v:"required#请输入账号"`    // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string `v:"required#评论内容类型错误"` // 评论类型: topic, ask, article, reply
	TargetId   uint   `v:"required#评论目标ID错误"` // 对应内容ID
	Content    string `v:"required#评论内容不能为空"` // 回复内容
}

// 执行删除内容
type ReplyDoDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的内容"` // 删除时ID不能为空
}

// 查询回复列表请求
type ReplyGetListReq struct {
	model.ReplyGetListInput
}
