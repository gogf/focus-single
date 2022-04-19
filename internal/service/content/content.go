package content

import (
	"context"

	"focus-single/internal/model/do"
	"focus-single/internal/service/bizctx"
	"focus-single/internal/service/category"
	"focus-single/internal/service/user"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/util/gutil"

	"focus-single/internal/consts"
	"focus-single/internal/dao"
	"focus-single/internal/model"
	"focus-single/internal/model/entity"
)

// GetList 查询内容列表
func GetList(ctx context.Context, in model.ContentGetListInput) (out *model.ContentGetListOutput, err error) {
	var (
		orm = dao.Content.Ctx(ctx)
		cls = dao.Content.Columns()
	)
	out = &model.ContentGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 默认查询topic
	if in.Type != "" {
		orm = orm.Where(cls.Type, in.Type)
	} else {
		orm = orm.Where(cls.Type, consts.ContentTypeTopic)
	}
	// 栏目检索
	if in.CategoryId > 0 {
		idArray, err := category.GetSubIdList(ctx, in.CategoryId)
		if err != nil {
			return out, err
		}
		orm = orm.Where(cls.CategoryId, idArray)
	}
	// 管理员可以查看所有文章
	if in.UserId > 0 && !user.IsAdmin(ctx, in.UserId) {
		orm = orm.Where(cls.UserId, in.UserId)
	}
	// 分配查询
	listModel := orm.Page(in.Page, in.Size)
	// 排序方式
	switch in.Sort {
	case consts.ContentSortHot:
		listModel = listModel.OrderDesc(cls.ViewCount)

	case consts.ContentSortActive:
		listModel = listModel.OrderDesc(cls.UpdatedAt)

	default:
		listModel = listModel.OrderDesc(cls.Id)
	}
	// 执行查询
	var list []*entity.Content
	if err = listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = orm.Count()
	if err != nil {
		return out, err
	}
	// Content
	if err = listModel.ScanList(&out.List, "Content"); err != nil {
		return out, err
	}
	// Category
	var categoryIds = gutil.ListItemValuesUnique(out.List, "Content", "CategoryId")
	err = dao.Category.Ctx(ctx).
		Where(dao.Category.Columns().Id, categoryIds).
		ScanList(&out.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return out, err
	}
	// User
	var userIds = gutil.ListItemValuesUnique(out.List, "Content", "UserId")
	err = dao.User.Ctx(ctx).
		Where(dao.User.Columns().Id, userIds).
		ScanList(&out.List, "User", "Content", "id:UserId")
	if err != nil {
		return out, err
	}
	return
}

// Search 搜索内容列表
func Search(ctx context.Context, in model.ContentSearchInput) (out *model.ContentSearchOutput, err error) {
	var (
		m           = dao.Content.Ctx(ctx)
		cls         = dao.Content.Columns()
		likePattern = `%` + in.Key + `%`
	)
	out = &model.ContentSearchOutput{
		Page: in.Page,
		Size: in.Size,
	}
	m = m.WhereLike(cls.Content, likePattern).WhereOrLike(cls.Title, likePattern)
	// 内容类型
	if in.Type != "" {
		m = m.Where(cls.Type, in.Type)
	}
	// 栏目检索
	if in.CategoryId > 0 {
		idArray, err := category.GetSubIdList(ctx, in.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(cls.CategoryId, idArray)
	}
	listModel := m.Page(in.Page, in.Size)
	switch in.Sort {
	case consts.ContentSortHot:
		listModel = listModel.OrderDesc(cls.ViewCount)

	case consts.ContentSortActive:
		listModel = listModel.OrderDesc(cls.UpdatedAt)

	default:
		listModel = listModel.OrderDesc(cls.Id)
	}
	listAll, err := listModel.All()
	if err != nil {
		return nil, err
	}
	// 没有数据
	if listAll.IsEmpty() {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return nil, err
	}
	// 搜索统计
	statsModel := m.Fields(cls.Type, "count(*) total").Group(cls.Type)
	statsAll, err := statsModel.All()
	if err != nil {
		return nil, err
	}
	out.Stats = make(map[string]int)
	for _, v := range statsAll {
		out.Stats[v[cls.Type].String()] = v["total"].Int()
	}
	// Content
	if err = listAll.ScanList(&out.List, "Content"); err != nil {
		return nil, err
	}
	// Category
	var categoryIds = gutil.ListItemValuesUnique(out.List, "Content", "CategoryId")
	err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().Id, categoryIds).
		ScanList(&out.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}
	// User
	var userIds = gutil.ListItemValuesUnique(out.List, "Content", "UserId")
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userIds).
		ScanList(&out.List, "User", "Content", "id:UserId")
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDetail 查询详情
func GetDetail(ctx context.Context, id uint) (out *model.ContentGetDetailOutput, err error) {
	out = &model.ContentGetDetailOutput{}
	if err = dao.Content.Ctx(ctx).WherePri(id).Scan(&out.Content); err != nil {
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
func Create(ctx context.Context, in model.ContentCreateInput) (out model.ContentCreateOutput, err error) {
	if in.UserId == 0 {
		in.UserId = bizctx.Get(ctx).User.Id
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
func Update(ctx context.Context, in model.ContentUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	var (
		cls       = dao.Content.Columns()
		ctxUser   = bizctx.Get(ctx).User
		condition = do.Content{
			Id: in.Id,
		}
	)
	if !ctxUser.IsAdmin {
		condition.UserId = ctxUser.Id
	}
	_, err := dao.Content.Ctx(ctx).Data(in).FieldsEx(cls.Id).Where(condition).Update()
	return err
}

// Delete 删除
func Delete(ctx context.Context, contentId uint) (err error) {
	var (
		ctxUser   = bizctx.Get(ctx).User
		condition = do.Content{
			Id:     contentId,
			UserId: ctxUser.Id,
		}
	)
	// 管理员直接删除文章和评论
	if ctxUser.IsAdmin {
		condition.UserId = ctxUser.Id
	}
	// 删除内容
	_, err = dao.Content.Ctx(ctx).Where(condition).Delete()
	// 删除评论
	if err == nil {
		_, err = dao.Reply.Ctx(ctx).Where(do.Reply{
			TargetId: contentId,
		}).Delete()
	}
	return err
}

// AddViewCount 浏览次数增加
func AddViewCount(ctx context.Context, id uint, count int) error {
	_, err := dao.Content.Ctx(ctx).WherePri(id).Increment(dao.Content.Columns().ViewCount, count)
	if err != nil {
		return err
	}
	return nil
}

// AddReplyCount 回复次数增加
func AddReplyCount(ctx context.Context, id uint, count int) error {
	_, err := dao.Content.Ctx(ctx).WherePri(id).Increment(dao.Content.Columns().ReplyCount, count)
	if err != nil {
		return err
	}
	return nil
}

// UnacceptedReply 取消采纳回复
func UnacceptedReply(ctx context.Context, id uint) error {
	_, err := dao.Content.Ctx(ctx).
		Data(dao.Content.Columns().AdoptedReplyId, 0).
		WherePri(id).
		Update()
	if err != nil {
		return err
	}
	return nil
}
