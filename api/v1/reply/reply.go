package reply

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetListContentReq struct {
	g.Meta     `path:"/reply" method:"get" summary:"查询回复列表" tags:"回复"`
	Page       int    `json:"page"       dc:"分页码"`
	Size       int    `json:"size"       dc:"分页数量"`
	TargetType string `json:"targetType" v:"required#评论内容类型错误" dc:"评论类型: topic/ask/article/reply"`
	TargetId   uint   `json:"targetId"   v:"required#评论目标ID错误" dc:"对应内容ID"`
}
type GetListContentRes struct {
	Content string `json:"content" dc:"HTML内容"`
}

type CreateReq struct {
	g.Meta     `path:"/reply" method:"put" summary:"执行回复接口" tags:"回复"`
	Title      string `json:"title" dc:"回复标题"`
	ParentId   uint   `json:"parentId" dc:"回复对应的上一级回复ID(没有的话默认为0)"`
	TargetType string `json:"targetType" v:"required#评论内容类型错误" dc:"评论类型: topic/ask/article/reply"`
	TargetId   uint   `json:"targetId"   v:"required#评论目标ID错误" dc:"对应内容ID"`
	Content    string `json:"content"    v:"required#评论内容不能为空" dc:"回复内容"`
}
type CreateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/reply" method:"delete" summary:"删除回复接口" tags:"回复"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"删除时ID不能为空"`
}
type DeleteRes struct{}
