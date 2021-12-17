package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SearchIndexReq struct {
	g.Meta `path:"/search" method:"get" summary:"展示内容搜索页面" tags:"搜索"`
	ContentGetListCommonReq
	Key        string `json:"key" v:"required#请输入搜索关键字" dc:"关键字"`
	Type       string `json:"type" dc:"内容模型"`
	CategoryId uint   `json:"cate" dc:"栏目ID"`
	Sort       int    `json:"sort" dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`
}
type SearchIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
