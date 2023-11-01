package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"focus-single/api/content/v1"
)

type IndexReq struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页"`
	v1.ContentGetListCommonReq
}
type IndexRes struct {
	v1.ContentGetListCommonRes
}
