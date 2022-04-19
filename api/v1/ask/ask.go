package ask

import (
	"focus-single/api/v1/content"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/ask" method:"get" tags:"问答" summary:"展示Ask列表页面"`
	content.GetListCommonReq
}
type IndexRes struct {
	content.GetListCommonRes
}

type DetailReq struct {
	g.Meta `path:"/ask/{Id}" method:"get" tags:"问答" summary:"展示Ask详情页面" `
	Id     uint `in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}
type DetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
