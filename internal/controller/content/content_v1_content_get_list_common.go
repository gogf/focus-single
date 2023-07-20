package content

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"focus-single/api/content/v1"
)

func (c *ControllerV1) ContentGetListCommon(ctx context.Context, req *v1.ContentGetListCommonReq) (res *v1.ContentGetListCommonRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
