package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserIndexReq struct {
	g.Meta `path:"/user/{UserId}" method:"get" summary:"展示查询用户内容列表页面" tags:"用户"`
	ContentGetListCommonReq
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}
type UserIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserArticleReq struct {
	g.Meta `path:"/user/article" method:"get" summary:"展示查询用户Article列表页面" tags:"用户"`
	ContentGetListCommonReq
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}
type UserArticleRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserTopicReq struct {
	g.Meta `path:"/user/topic" method:"get" summary:"展示查询用户Topic列表页面" tags:"用户"`
	ContentGetListCommonReq
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}
type UserTopicRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserAskReq struct {
	g.Meta `path:"/user/ask" method:"get" summary:"展示查询用户Ask列表页面" tags:"用户"`
	ContentGetListCommonReq
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}
type UserAskRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserLogoutReq struct {
	g.Meta `path:"/user/logout" method:"get" summary:"执行用户注销接口" tags:"个人"`
}

type UserLogoutRes struct{
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}