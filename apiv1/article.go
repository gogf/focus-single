package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleIndexReq struct {
	g.Meta `path:"/article" method:"get" tags:"文章" summary:"展示Article列表页面"`
	ContentGetListCommonReq
}
type ArticleIndexRes struct {
	ContentGetListCommonRes
}

type ArticleDetailReq struct {
	g.Meta `path:"/article/{Id}" method:"get" tags:"文章" summary:"展示Article详情页面" `
	Id     uint `in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}
type ArticleDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
