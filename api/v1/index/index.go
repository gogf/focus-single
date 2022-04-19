package index

import (
	"focus-single/api/v1/content"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页"`
	content.GetListCommonReq
}
type Res struct {
	content.GetListCommonRes
}
