package content

import (
	"context"

	"focus-single/api/content/v1"
	"focus-single/internal/service"
)

func (c *ControllerV1) ContentDelete(ctx context.Context, req *v1.ContentDeleteReq) (res *v1.ContentDeleteRes, err error) {
	err = service.Content().Delete(ctx, req.Id)
	return
}
