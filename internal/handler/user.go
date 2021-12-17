package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// 用户管理
	User = handlerUser{}
)

type handlerUser struct{}

func (a *handlerUser) Index(ctx context.Context, req *apiv1.UserIndexReq) (res *apiv1.UserIndexRes, err error) {
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *handlerUser) Article(ctx context.Context, req *apiv1.UserArticleReq) (res *apiv1.UserArticleRes, err error) {
	req.Type = consts.ContentTypeArticle
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *handlerUser) Topic(ctx context.Context, req *apiv1.UserTopicReq) (res *apiv1.UserTopicRes, err error) {
	req.Type = consts.ContentTypeTopic
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *handlerUser) Ask(ctx context.Context, req *apiv1.UserAskReq) (res *apiv1.UserAskRes, err error) {
	req.Type = consts.ContentTypeAsk
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *handlerUser) getContentList(ctx context.Context, userId uint, req apiv1.ContentGetListCommonReq) error {
	if out, err := service.User.GetList(ctx, model.UserGetContentListInput{
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
		title := service.View.GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		})
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        out,
			Title:       title,
		})
		return nil
	}
}
