package bizctx

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"focus-single/internal/consts"
)

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	Id       uint   // 用户ID
	Passport string // 用户账号
	Nickname string // 用户名称
	Avatar   string // 用户头像
	IsAdmin  bool   // 是否是管理员
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func Init(r *ghttp.Request, customCtx *Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func Get(ctx context.Context) *Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func SetUser(ctx context.Context, ctxUser *ContextUser) {
	Get(ctx).User = ctxUser
}

// SetData 将上下文信息设置到上下文请求中，注意是完整覆盖
func SetData(ctx context.Context, data g.Map) {
	Get(ctx).Data = data
}
