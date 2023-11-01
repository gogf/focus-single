package user

import (
	"context"

	v1 "focus-single/api/content/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) getContentList(ctx context.Context, userId uint, req v1.ContentGetListCommonReq) (err error) {
	type getContentListInfo struct {
		Content *model.ContentGetListOutput `json:"content"` // 查询用户
		User    *model.UserGetProfileOutput `json:"user"`    // 查询用户
		Stats   map[string]int              // 发布内容数量
	}
	var (
		data    = getContentListInfo{}
		ctxUser = service.BizCtx().Get(ctx).User
	)
	// 用户内容信息
	data.Content, err = service.Content().GetList(ctx, model.ContentGetListInput{
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
	data.User, err = service.User().GetProfileById(ctx, ctxUser.Id)
	if err != nil {
		return err
	}
	// 用户统计信息
	data.Stats, err = service.User().GetUserStats(ctx, ctxUser.Id)
	if err != nil {
		return err
	}

	title := service.View().GetTitle(ctx, &model.ViewGetTitleInput{
		ContentType: req.Type,
		CategoryId:  req.CategoryId,
	})
	service.View().Render(ctx, model.View{
		ContentType: req.Type,
		Data:        data,
		Title:       title,
	})
	return nil
}
