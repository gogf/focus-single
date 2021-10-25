package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/cnt"
	"focus/app/internal/model"
	"focus/app/internal/service"
)

var (
	// 主题管理
	Topic = topicAct{}
)

type topicAct struct{}

func (a *topicAct) Index(ctx context.Context, req *api.ContentGetListReq) (res *api.ContentGetListRes, err error) {
	req.Type = cnt.ContentTypeTopic
	if getListRes, err := service.Content.GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
		UserId:     service.Session.GetUser(ctx).Id,
	}); err != nil {
		return nil, err
	} else {
		title := service.View.GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.ContentListCommonReq.CategoryId,
		})
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       title,
		})
		return nil, nil
	}
}

func (a *topicAct) Detail(ctx context.Context, req *api.ContentDetailReq) (res *api.ContentDetailRes, err error) {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		if getDetailRes == nil {
			service.View.Render404(ctx)
			return nil, nil
		}
		err = service.Content.AddViewCount(ctx, req.Id, 1)
		service.View.Render(ctx, model.View{
			ContentType: cnt.ContentTypeTopic,
			Data:        getDetailRes,
			Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			}),
			BreadCrumb: service.View.GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			}),
		})
		return nil, nil
	}
}
