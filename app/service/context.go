package service

import (
	"context"
	"focus/app/cnt"
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 上下文管理服务
var Context = contextShared{}

type contextShared struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextShared) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(cnt.ContextKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextShared) Get(ctx context.Context) *model.Context {
	value := ctx.Value(cnt.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *contextShared) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *contextShared) SetData(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}
