package api

import "github.com/gogf/gf/frame/g"

// 赞
type InteractZanReq struct {
	g.Meta `method:"post" summary:"赞" tags:"交互"`
	Id     uint   `v:"min:1#请选择需要赞的内容"`
	Type   string `v:"required#请提交需要赞的内容类型"` // content, reply
}
type InteractZanRes struct{}

// 取消赞
type InteractCancelZanReq struct {
	g.Meta `method:"post" summary:"取消赞" tags:"交互"`
	Id     uint   `v:"min:1#请选择需要取消赞的内容"`
	Type   string `v:"required#请提交需要取消赞的内容类型"` // content, reply
}
type InteractCancelZanRes struct{}

// 踩
type InteractCaiReq struct {
	g.Meta `method:"post" summary:"踩" tags:"交互"`
	Id     uint   `v:"min:1#请选择需要踩的内容"`
	Type   string `v:"required#请提交需要踩的内容类型"` // content, reply
}
type InteractCaiRes struct{}

// 取消踩
type InteractCancelCaiReq struct {
	g.Meta `method:"post" summary:"取消踩" tags:"交互"`
	Id     uint   `v:"min:1#请选择需要取消踩的内容"`
	Type   string `v:"required#请提交需要取消踩的内容类型"` // content, reply
}
type InteractCancelCaiRes struct{}
