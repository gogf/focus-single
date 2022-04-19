package user

import (
	"context"

	"focus-single/api/v1/content"
	v1 "focus-single/api/v1/user"
	"focus-single/internal/consts"
	"focus-single/internal/model"
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

func (c *controller) getContentList(ctx context.Context, userId uint, req content.GetListCommonReq) error {
	if out, err := user.GetList(ctx, model.UserGetContentListInput{
		ContentGetListInput: model.ContentGetListInput{
			Type:       req.Type,
			CategoryId: req.CategoryId,
			Page:       req.Page,
			Size:       req.Size,
			Sort:       req.Sort,
			UserId:     userId,
		},
	}); err != nil {
		return err
	} else {
		title := view.GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		})
		view.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        out,
			Title:       title,
		})
		return nil
	}
}

func (c *controller) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	if err = user.Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(middleware.LoginUrl)
	return
}
