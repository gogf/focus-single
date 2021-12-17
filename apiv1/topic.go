package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TopicIndexReq struct {
	g.Meta `path:"/topic" method:"get" tags:"话题" summary:"展示Topic列表页面"`
	ContentGetListCommonReq
}
type TopicIndexRes struct {
	ContentGetListCommonRes
}

type TopicDetailReq struct {
	g.Meta `path:"/topic/{Id}" method:"get" tags:"话题" summary:"展示Topic详情页面" `
	Id     uint `in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}
type TopicDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
