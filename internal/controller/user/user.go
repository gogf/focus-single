package user

import (
	"context"

	contentV1 "focus-single/api/v1/content"
	v1 "focus-single/api/v1/user"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service/bizctx"
	"focus-single/internal/service/content"
	"focus-single/internal/service/middleware"
	"focus-single/internal/service/user"
	"focus-single/internal/service/view"
	"github.com/gogf/gf/v2/frame/g"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	err = c.getContentList(ctx, req.UserId, req.GetListCommonReq)
	return
}

func (c *controller) Article(ctx context.Context, req *v1.ArticleReq) (res *v1.ArticleRes, err error) {
	req.Type = consts.ContentTypeArticle
	err = c.getContentList(ctx, req.UserId, req.GetListCommonReq)
	return
}

func (c *controller) Topic(ctx context.Context, req *v1.TopicReq) (res *v1.TopicRes, err error) {
	req.Type = consts.ContentTypeTopic
	err = c.getContentList(ctx, req.UserId, req.GetListCommonReq)
	return
}

func (c *controller) Ask(ctx context.Context, req *v1.AskReq) (res *v1.AskRes, err error) {
	req.Type = consts.ContentTypeAsk
	err = c.getContentList(ctx, req.UserId, req.GetListCommonReq)
	return
}

func (c *controller) getContentList(ctx context.Context, userId uint, req contentV1.GetListCommonReq) (err error) {
	type getContentListInfo struct {
		Content *model.ContentGetListOutput `json:"content"` // 查询用户
		User    *model.UserGetProfileOutput `json:"user"`    // 查询用户
		Stats   map[string]int              // 发布内容数量
	}
	var (
		data    = getContentListInfo{}
		ctxUser = bizctx.Get(ctx).User
	)
	// 用户内容信息
	data.Content, err = content.GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
		UserId:     userId,
	})
	if err != nil {
		return err
	}
	// 用户资料信息
	data.User, err = user.GetProfileById(ctx, ctxUser.Id)
	if err != nil {
		return err
	}
	// 用户统计信息
	data.Stats, err = user.GetUserStats(ctx, ctxUser.Id)
	if err != nil {
		return err
	}

	title := view.GetTitle(ctx, &model.ViewGetTitleInput{
		ContentType: req.Type,
		CategoryId:  req.CategoryId,
	})
	view.Render(ctx, model.View{
		ContentType: req.Type,
		Data:        data,
		Title:       title,
	})
	return nil
}

func (c *controller) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	if err = user.Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(middleware.LoginUrl)
	return
}
