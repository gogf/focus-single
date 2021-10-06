package service

import (
	"context"
	"focus/app/internal/cnt"
	"focus/app/internal/dao"
	"focus/app/internal/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

// 内容管理服务
var Content = contentService{}

type contentService struct{}

// 查询内容列表
func (s *contentService) GetList(ctx context.Context, input model.ContentGetListInput) (output *model.ContentGetListOutput, err error) {
	var (
		m = dao.Content.Ctx(ctx)
	)
	output = &model.ContentGetListOutput{
		Page: input.Page,
		Size: input.Size,
	}
	// 默认查询topic
	if input.Type != "" {
		m = m.Where(dao.Content.Columns.Type, input.Type)
	} else {
		m = m.Where(dao.Content.Columns.Type, cnt.ContentTypeTopic)
	}
	// 栏目检索
	if input.CategoryId > 0 {
		idArray, err := Category.GetSubIdList(ctx, input.CategoryId)
		if err != nil {
			return output, err
		}
		m = m.Where(dao.Content.Columns.CategoryId, idArray)
	}
	// 管理员可以查看所有文章
	if input.UserId > 0 && !User.IsAdminShow(ctx, input.UserId) {
		m = m.Where(dao.Content.Columns.UserId, input.UserId)
	}
	// 分配查询
	listModel := m.Page(input.Page, input.Size)
	// 排序方式
	switch input.Sort {
	case cnt.ContentSortHot:
		listModel = listModel.OrderDesc(dao.Content.Columns.ViewCount)

	case cnt.ContentSortActive:
		listModel = listModel.OrderDesc(dao.Content.Columns.UpdatedAt)

	default:
		listModel = listModel.OrderDesc(dao.Content.Columns.Id)
	}
	// 执行查询
	var list []*model.Content
	if err := listModel.Scan(&list); err != nil {
		return output, err
	}
	// 没有数据
	if len(list) == 0 {
		return output, nil
	}
	output.Total, err = m.Count()
	if err != nil {
		return output, err
	}
	// Content
	if err := listModel.ScanList(&output.List, "Content"); err != nil {
		return output, err
	}
	// Category
	err = dao.Category.Ctx(ctx).
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.Columns.Id, gutil.ListItemValuesUnique(output.List, "Content", "CategoryId")).
		ScanList(&output.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return output, err
	}
	// User
	err = dao.User.Ctx(ctx).
		Fields(model.ContentListUserItem{}).
		Where(dao.User.Columns.Id, gutil.ListItemValuesUnique(output.List, "Content", "UserId")).
		ScanList(&output.List, "User", "Content", "id:UserId")
	if err != nil {
		return output, err
	}
	return
}

// 搜索内容列表
func (s *contentService) Search(ctx context.Context, input model.ContentSearchInput) (output *model.ContentSearchOutput, err error) {
	var (
		m           = dao.Content.Ctx(ctx)
		likePattern = `%` + input.Key + `%`
	)
	output = &model.ContentSearchOutput{
		Page: input.Page,
		Size: input.Size,
	}
	m = m.WhereLike(dao.Content.Columns.Content, likePattern).WhereOrLike(dao.Content.Columns.Title, likePattern)
	// 内容类型
	if input.Type != "" {
		m = m.Where(dao.Content.Columns.Type, input.Type)
	}
	// 栏目检索
	if input.CategoryId > 0 {
		idArray, err := Category.GetSubIdList(ctx, input.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(dao.Content.Columns.CategoryId, idArray)
	}
	listModel := m.Page(input.Page, input.Size)
	switch input.Sort {
	case cnt.ContentSortHot:
		listModel = listModel.OrderDesc(dao.Content.Columns.ViewCount)

	case cnt.ContentSortActive:
		listModel = listModel.OrderDesc(dao.Content.Columns.UpdatedAt)

	//case model.ContentSortScore:
	//	listModel = listModel.OrderDesc("score")

	default:
		listModel = listModel.OrderDesc(dao.Content.Columns.Id)
	}
	all, err := listModel.All()
	if err != nil {
		return nil, err
	}
	// 没有数据
	if all.IsEmpty() {
		return output, nil
	}
	output.Total, err = m.Count()
	if err != nil {
		return nil, err
	}
	// 搜索统计
	statsModel := m.Fields(dao.Content.Columns.Type, "count(*) total").Group(dao.Content.Columns.Type)
	statsAll, err := statsModel.All()
	if err != nil {
		return nil, err
	}
	output.Stats = make(map[string]int)
	for _, v := range statsAll {
		output.Stats[v["type"].String()] = v["total"].Int()
	}
	// Content
	if err := all.ScanList(&output.List, "Content"); err != nil {
		return nil, err
	}
	// Category
	err = dao.Category.Ctx(ctx).
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.Columns.Id, gutil.ListItemValuesUnique(output.List, "Content", "CategoryId")).
		ScanList(&output.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}
	// User
	err = dao.User.Ctx(ctx).
		Fields(model.ContentListUserItem{}).
		Where(dao.User.Columns.Id, gutil.ListItemValuesUnique(output.List, "Content", "UserId")).
		ScanList(&output.List, "User", "Content", "id:UserId")
	if err != nil {
		return nil, err
	}

	return output, nil
}

// 查询详情
func (s *contentService) GetDetail(ctx context.Context, id uint) (output *model.ContentGetDetailOutput, err error) {
	output = &model.ContentGetDetailOutput{}
	if err := dao.Content.Ctx(ctx).WherePri(id).Scan(&output.Content); err != nil {
		return nil, err
	}
	// 没有数据
	if output.Content == nil {
		return nil, nil
	}
	err = dao.User.Ctx(ctx).WherePri(output.Content.UserId).Scan(&output.User)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// 创建内容
func (s *contentService) Create(ctx context.Context, input model.ContentCreateInput) (output model.ContentCreateOutput, err error) {
	if input.UserId == 0 {
		input.UserId = Context.Get(ctx).User.Id
	}
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(input); err != nil {
		return output, err
	}
	lastInsertId, err := dao.Content.Ctx(ctx).Data(input).InsertAndGetId()
	if err != nil {
		return output, err
	}
	return model.ContentCreateOutput{ContentId: uint(lastInsertId)}, err
}

// 修改
func (s *contentService) Update(ctx context.Context, input model.ContentUpdateInput) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(input); err != nil {
			return err
		}
		_, err := dao.Content.
			Ctx(ctx).
			Data(input).
			FieldsEx(dao.Content.Columns.Id).
			Where(dao.Content.Columns.Id, input.Id).
			Where(dao.Content.Columns.UserId, Context.Get(ctx).User.Id).
			Update()
		return err
	})
}

// 删除
func (s *contentService) Delete(ctx context.Context, id uint) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		user := Context.Get(ctx).User
		// 管理员直接删除文章和评论
		if user.IsAdmin {
			_, err := dao.Content.Ctx(ctx).Where(dao.Content.Columns.Id, id).Delete()
			if err == nil {
				_, err = dao.Reply.Ctx(ctx).Where(dao.Reply.Columns.TargetId, id).Delete()
			}
			return err
		}
		// 删除内容
		_, err := dao.Content.Ctx(ctx).Where(g.Map{
			dao.Content.Columns.Id:     id,
			dao.Content.Columns.UserId: Context.Get(ctx).User.Id,
		}).Delete()
		// 删除评论
		if err == nil {
			err = Reply.DeleteByUserContentId(ctx, user.Id, id)
		}
		return err
	})
}

// 浏览次数增加
func (s *contentService) AddViewCount(ctx context.Context, id uint, count int) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Content.Ctx(ctx).WherePri(id).Increment(dao.Content.Columns.ViewCount, count)
		if err != nil {
			return err
		}
		return nil
	})
}

// 回复次数增加
func (s *contentService) AddReplyCount(ctx context.Context, id uint, count int) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Content.Ctx(ctx).WherePri(id).Increment(dao.Content.Columns.ReplyCount, count)
		if err != nil {
			return err
		}
		return nil
	})
}

// 采纳回复
func (s *contentService) AdoptReply(ctx context.Context, id uint, replyId uint) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Content.Ctx(ctx).
			Data(dao.Content.Columns.AdoptedReplyId, replyId).
			Where(dao.Content.Columns.UserId, Context.Get(ctx).User.Id).
			WherePri(id).
			Update()
		if err != nil {
			return err
		}
		return nil
	})
}

// 取消采纳回复
func (s *contentService) UnacceptedReply(ctx context.Context, id uint) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Content.Ctx(ctx).
			Data(dao.Content.Columns.AdoptedReplyId, 0).
			WherePri(id).
			Update()
		if err != nil {
			return err
		}
		return nil
	})
}
