package api

import "github.com/gogf/gf/frame/g"

type RegisterIndexReq struct {
	g.Meta `method:"get" summary:"展示注册页面" tags:"注册"`
}
type RegisterIndexRes struct {
	g.Meta `mime:"text/html"`
}
