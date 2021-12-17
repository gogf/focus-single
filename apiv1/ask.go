package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AskIndexReq struct {
	g.Meta `path:"/ask" method:"get" tags:"问答" summary:"展示Ask列表页面"`
	ContentGetListCommonReq
}
type AskIndexRes struct {
	ContentGetListCommonRes
}

type AskDetailReq struct {
	g.Meta `path:"/ask/{Id}" method:"get" tags:"问答" summary:"展示Ask详情页面" `
	Id     uint `in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}
type AskDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
