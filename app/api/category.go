package api

import (
	"focus/app/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// 查询栏目树形列表
type CategoryGetTreeReq struct {
	g.Meta      `method:"get" summary:"获取分类列表" dc:"获取分类列表，构造成树形结构返回" tags:"分类"`
	ContentType string `json:"contentType" in:"query" dc:"栏目类型：topic/question/article。当传递空时表示获取所有类型的栏目"`
}
type CategoryGetTreeRes struct {
	List []*model.CategoryTreeItem `json:"list" dc:"栏目列表"`
}
