package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"

	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/model/entity"
	"focus-single/internal/service/internal/dao"
)

type sContent struct{}

var insContent = sContent{}

// Content 内容管理服务
func Content() *sContent {
	return &insContent
}

// GetList 查询内容列表
func (s *sContent) GetList(ctx context.Context, in model.ContentGetListInput) (out *model.ContentGetListOutput, err error) {
	var (
		m = dao.Content.Ctx(ctx)
	)
	out = &model.ContentGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 默认查询topic
	if in.Type != "" {
		m = m.Where(dao.Content.Columns().Type, in.Type)
	} else {
		m = m.Where(dao.Content.Columns().Type, consts.ContentTypeTopic)
	}
	// 栏目检索
	if in.CategoryId > 0 {
		idArray, err := Category().GetSubIdList(ctx, in.CategoryId)
		if err != nil {
			return out, err
		}
		m = m.Where(dao.Content.Columns().CategoryId, idArray)
	}
	// 管理员可以查看所有文章
	if in.UserId > 0 && !User().IsAdminShow(ctx, in.UserId) {
		m = m.Where(dao.Content.Columns().UserId, in.UserId)
	}
	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	switch in.Sort {
	case consts.ContentSortHot:
		listModel = listModel.OrderDesc(dao.Content.Columns().ViewCount)

	case consts.ContentSortActive:
		listModel = listModel.OrderDesc(dao.Content.Columns().UpdatedAt)

	default:
		listModel = listModel.OrderDesc(dao.Content.Columns().Id)
	}
	// 执行查询
	var list []*entity.Content
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Content
	if err := listModel.ScanList(&out.List, "Content"); err != nil {
		return out, err
	}
	// Category
	err = dao.Category.Ctx(ctx).
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.Columns().Id, gutil.ListItemValuesUnique(out.List, "Content", "CategoryId")).
		ScanList(&out.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return out, err
	}
	// User
	err = dao.User.Ctx(ctx).
		Fields(model.ContentListUserItem{}).
		Where(dao.User.Columns().Id, gutil.ListItemValuesUnique(out.List, "Content", "UserId")).
		ScanList(&out.List, "User", "Content", "id:UserId")
	if err != nil {
		return out, err
	}
	return
}

// Search 搜索内容列表
func (s *sContent) Search(ctx context.Context, in model.ContentSearchInput) (out *model.ContentSearchOutput, err error) {
	var (
		m           = dao.Content.Ctx(ctx)
		likePattern = `%` + in.Key + `%`
	)
	out = &model.ContentSearchOutput{
		Page: in.Page,
		Size: in.Size,
	}
	m = m.WhereLike(dao.Content.Columns().Content, likePattern).WhereOrLike(dao.Content.Columns().Title, likePattern)
	// 内容类型
	if in.Type != "" {
		m = m.Where(dao.Content.Columns().Type, in.Type)
	}
	// 栏目检索
	if in.CategoryId > 0 {
		idArray, err := Category().GetSubIdList(ctx, in.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(dao.Content.Columns().CategoryId, idArray)
	}
	listModel := m.Page(in.Page, in.Size)
	switch in.Sort {
	case consts.ContentSortHot:
		listModel = listModel.OrderDesc(dao.Content.Columns().ViewCount)

	case consts.ContentSortActive:
		listModel = listModel.OrderDesc(dao.Content.Columns().UpdatedAt)

	// case model.ContentSortScore:
	//	listModel = listModel.OrderDesc("score")

	default:
		listModel = listModel.OrderDesc(dao.Content.Columns().Id)
	}
	all, err := listModel.All()
	if err != nil {
		return nil, err
	}
	// 没有数据
	if all.IsEmpty() {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return nil, err
	}
	// 搜索统计
	statsModel := m.Fields(dao.Content.Columns().Type, "count(*) total").Group(dao.Content.Columns().Type)
	statsAll, err := statsModel.All()
	if err != nil {
		return nil, err
	}
	out.Stats = make(map[string]int)
	for _, v := range statsAll {
		out.Stats[v["type"].String()] = v["total"].Int()
	}
	// Content
	if err := all.ScanList(&out.List, "Content"); err != nil {
		return nil, err
	}
	// Category
	err = dao.Category.Ctx(ctx).
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.Columns().Id, gutil.ListItemValuesUnique(out.List, "Content", "CategoryId")).
		ScanList(&out.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}
	// User
	err = dao.User.Ctx(ctx).
		Fields(model.ContentListUserItem{}).
		Where(dao.User.Columns().Id, gutil.ListItemValuesUnique(out.List, "Content", "UserId")).
		ScanList(&out.List, "User", "Content", "id:UserId")
	if err != nil {
		return nil, err
	}

	return out, nil
}

// GetDetail 查询详情
func (s *sContent) GetDetail(ctx context.Context, id uint) (out *model.ContentGetDetailOutput, err error) {
	out = &model.ContentGetDetailOutput{}
	if err := dao.Content.Ctx(ctx).WherePri(id).Scan(&out.Content); err != nil {
		return nil, err
	}
	// 没有数据
	if out.Content == nil {
		return nil, nil
	}
	err = dao.User.Ctx(ctx).WherePri(out.Content.UserId).Scan(&out.User)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Create 创建内容
func (s *sContent) Create(ctx context.Context, in model.ContentCreateInput) (out model.ContentCreateOutput, err error) {
	if in.UserId == 0 {
		in.UserId = Context().Get(ctx).User.Id
	}
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.Content.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ContentCreateOutput{ContentId: uint(lastInsertID)}, err
}

// Update 修改
func (s *sContent) Update(ctx context.Context, in model.ContentUpdateInput) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.Content.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Content.Columns().Id).
			Where(dao.Content.Columns().Id, in.Id).
			Where(dao.Content.Columns().UserId, Context().Get(ctx).User.Id).
			Update()
		return err
	})
}

// Delete 删除
func (s *sContent) Delete(ctx context.Context, id uint) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		user := Context().Get(ctx).User
		// 管理员直接删除文章和评论
		if user.IsAdmin {
			_, err := dao.Content.Ctx(ctx).Where(dao.Content.Columns().Id, id).Delete()
			if err == nil {
				_, err = dao.Reply.Ctx(ctx).Where(dao.Reply.Columns().TargetId, id).Delete()
			}
			return err
		}
		// 删除内容
		_, err := dao.Content.Ctx(ctx).Where(g.Map{
			dao.Content.Columns().Id:     id,
			dao.Content.Columns().UserId: Context().Get(ctx).User.Id,
		}).Delete()
		// 删除评论
		if err == nil {
			err = Reply().DeleteByUserContentId(ctx, user.Id, id)
		}
		return err
	})
}

// AddViewCount 浏览次数增加
func (s *sContent) AddViewCount(ctx context.Context, id uint, count int) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Content.Ctx(ctx).WherePri(id).Increment(dao.Content.Columns().ViewCount, count)
		if err != nil {
			return err
		}
		return nil
	})
}

// AddReplyCount 回复次数增加
func (s *sContent) AddReplyCount(ctx context.Context, id uint, count int) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Content.Ctx(ctx).WherePri(id).Increment(dao.Content.Columns().ReplyCount, count)
		if err != nil {
			return err
		}
		return nil
	})
}

// AdoptReply 采纳回复
func (s *sContent) AdoptReply(ctx context.Context, id uint, replyID uint) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Content.Ctx(ctx).
			Data(dao.Content.Columns().AdoptedReplyId, replyID).
			Where(dao.Content.Columns().UserId, Context().Get(ctx).User.Id).
			WherePri(id).
			Update()
		if err != nil {
			return err
		}
		return nil
	})
}

// UnacceptedReply 取消采纳回复
func (s *sContent) UnacceptedReply(ctx context.Context, id uint) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Content.Ctx(ctx).
			Data(dao.Content.Columns().AdoptedReplyId, 0).
			WherePri(id).
			Update()
		if err != nil {
			return err
		}
		return nil
	})
}
