package topic

import (
	"context"

	"focus-single/api/topic/v1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) TopicIndex(ctx context.Context, req *v1.TopicIndexReq) (res *v1.TopicIndexRes, err error) {
	req.Type = consts.ContentTypeTopic
	out, err := service.Content().GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	})
	if err != nil {
		return nil, err
	}
	title := service.View().GetTitle(ctx, &model.ViewGetTitleInput{
		ContentType: req.Type,
		CategoryId:  req.CategoryId,
	})
	service.View().Render(ctx, model.View{
		ContentType: req.Type,
		Data:        out,
		Title:       title,
	})
	return
}
