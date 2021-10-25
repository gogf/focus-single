package api

import "github.com/gogf/gf/v2/frame/g"

type LoginIndexReq struct {
	g.Meta `method:"get" summary:"展示登录页面" tags:"登录"`
}
type LoginIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
