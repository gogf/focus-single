package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/cnt"
	"focus/app/internal/model"
	"focus/app/internal/service"
)

var (
	// 文章管理
	Article = articleAct{}
)

type articleAct struct{}

func (a *articleAct) Index(ctx context.Context, req *api.ContentGetListReq) (res *api.ContentGetListRes, err error) {
	req.Type = cnt.ContentTypeArticle
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
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: req.Type,
				CategoryId:  req.ContentListCommonReq.CategoryId,
			}),
		})
	}
	return
}

func (a *articleAct) Detail(ctx context.Context, req *api.ContentDetailReq) (res *api.ContentDetailRes, err error) {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		if getDetailRes == nil {
			service.View.Render404(ctx)
			return nil, nil
		}
		if err = service.Content.AddViewCount(ctx, req.Id, 1); err != nil {
			return res, err
		}
		service.View.Render(ctx, model.View{
			ContentType: cnt.ContentTypeArticle,
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
	}
	return
}
