package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页"`
	ContentGetListCommonReq
}
type IndexRes struct {
	ContentGetListCommonRes
}
