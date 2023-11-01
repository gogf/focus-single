package register

import (
	"context"

	"focus-single/api/register/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) RegisterIndex(ctx context.Context, req *v1.RegisterIndexReq) (res *v1.RegisterIndexRes, err error) {
	service.View().Render(ctx, model.View{})
	return
}
