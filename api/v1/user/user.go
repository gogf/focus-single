package user

import (
	"focus-single/api/v1/content"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/user/{UserId}" method:"get" summary:"展示查询用户内容列表页面" tags:"用户"`
	content.GetListCommonReq
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}
type IndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ArticleReq struct {
	g.Meta `path:"/user/article" method:"get" summary:"展示查询用户Article列表页面" tags:"用户"`
	content.GetListCommonReq
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}
type ArticleRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type TopicReq struct {
	g.Meta `path:"/user/topic" method:"get" summary:"展示查询用户Topic列表页面" tags:"用户"`
	content.GetListCommonReq
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}
type TopicRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type AskReq struct {
	g.Meta `path:"/user/ask" method:"get" summary:"展示查询用户Ask列表页面" tags:"用户"`
	content.GetListCommonReq
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}
type AskRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"get" summary:"执行用户注销接口" tags:"个人"`
}

type LogoutRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
