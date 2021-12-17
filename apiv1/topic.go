package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Topic列表
type TopicIndexReq struct {
	g.Meta `path:"/topic" method:"get" tags:"内容" summary:"展示Topic列表页面"`
	ContentGetListCommonReq
}
type TopicIndexRes struct {
	ContentGetListCommonRes
}

// Topic详情
type TopicDetailReq struct {
	g.Meta `path:"/topic/{Id}" method:"get" tags:"内容" summary:"展示Topic详情页面" `
	Id     uint `in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}
type TopicDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
