package apiv1

import "github.com/gogf/gf/v2/frame/g"

type InteractZanReq struct {
	g.Meta `path:"/interact/zan" method:"put" summary:"赞" tags:"交互"`
	Id     uint   `json:"id"   v:"min:1#请选择需要赞的内容"`
	Type   string `json:"type" v:"required#请提交需要赞的内容类型" dc:"content/reply"`
}
type InteractZanRes struct{}

type InteractCancelZanReq struct {
	g.Meta `path:"/interact/zan" method:"delete" summary:"取消赞" tags:"交互"`
	Id     uint   `json:"id"   v:"min:1#请选择需要取消赞的内容"`
	Type   string `json:"type" v:"required#请提交需要取消赞的内容类型" `
}
type InteractCancelZanRes struct{}

type InteractCaiReq struct {
	g.Meta `path:"/interact/cai" method:"put" summary:"踩" tags:"交互"`
	Id     uint   `json:"id"   v:"min:1#请选择需要踩的内容"`
	Type   string `json:"type" v:"required#请提交需要踩的内容类型" dc:"content/reply"`
}
type InteractCaiRes struct{}

type InteractCancelCaiReq struct {
	g.Meta `path:"/interact/cai" method:"delete" summary:"取消踩" tags:"交互"`
	Id     uint   `json:"id"   v:"min:1#请选择需要取消踩的内容"`
	Type   string `json:"type" v:"required#请提交需要取消踩的内容类型" dc:"content/reply"`
}
type InteractCancelCaiRes struct{}
