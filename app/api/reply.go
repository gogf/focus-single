package api

import (
	"focus/app/internal/model"
	"github.com/gogf/gf/frame/g"
)

type ReplyDoCreateReq struct {
	g.Meta `method:"post" summary:"执行回复接口" tags:"回复"`
	model.ReplyCreateInput
	ParentId   uint   `v:"required#请输入账号"`    // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string `v:"required#评论内容类型错误"` // 评论类型: topic, ask, article, reply
	TargetId   uint   `v:"required#评论目标ID错误"` // 对应内容ID
	Content    string `v:"required#评论内容不能为空"` // 回复内容
}
type ReplyDoCreateRes struct{}

// 执行删除内容
type ReplyDoDeleteReq struct {
	g.Meta `method:"post" summary:"删除回复接口" tags:"回复"`
	Id     uint `v:"min:1#请选择需要删除的内容"` // 删除时ID不能为空
}
type ReplyDoDeleteRes struct{}

// 查询回复列表请求
type ReplyGetListReq struct {
	g.Meta `method:"post" summary:"查询回复列表" tags:"回复"`
	model.ReplyGetListInput
}
type ReplyGetListRes struct {
	Content string // HTML内容
}
