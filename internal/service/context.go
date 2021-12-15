package service

import (
	"context"

	"focus-single/internal/consts"
	"focus-single/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 上下文管理服务
var Context = serviceContext{}

type serviceContext struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *serviceContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *serviceContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *serviceContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *serviceContext) SetData(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}
