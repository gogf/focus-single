package service

import (
	"context"

	"focus-single/internal/consts"
	"focus-single/internal/model/entity"
	"focus-single/internal/service/internal/dao"
	"focus-single/internal/service/internal/dto"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 交互管理服务
var Interact = serviceInteract{}

type serviceInteract struct{}

const (
	contextMapKeyForMyInteractList = "ContextMapKeyForMyInteractList"
)

// 赞
func (s *serviceInteract) Zan(ctx context.Context, targetType string, targetId uint) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		customCtx := Context.Get(ctx)
		if customCtx == nil || customCtx.User == nil {
			return nil
		}
		r, err := dao.Interact.Ctx(ctx).Data(&entity.Interact{
			UserId:     customCtx.User.Id,
			TargetId:   targetId,
			TargetType: targetType,
			Type:       consts.InteractTypeZan,
		}).FieldsEx(dao.Interact.Columns.Id).InsertIgnore()
		if err != nil {
			return err
		}
		if n, _ := r.RowsAffected(); n == 0 {
			return gerror.New("您已经赞过啦")
		}
		return s.updateCount(ctx, consts.InteractTypeZan, targetType, targetId, 1)
	})
}

// 取消赞
func (s *serviceInteract) CancelZan(ctx context.Context, targetType string, targetId uint) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		customCtx := Context.Get(ctx)
		if customCtx == nil || customCtx.User == nil {
			return nil
		}
		r, err := dao.Interact.Ctx(ctx).Where(dto.Interact{
			Type:       consts.InteractTypeZan,
			UserId:     Context.Get(ctx).User.Id,
			TargetId:   targetId,
			TargetType: targetType,
		}).Delete()
		if err != nil {
			return err
		}
		if n, _ := r.RowsAffected(); n == 0 {
			return nil
		}
		return s.updateCount(ctx, consts.InteractTypeZan, targetType, targetId, -1)
	})
}

// 我是否有对指定内容赞
func (s *serviceInteract) DidIZan(ctx context.Context, targetType string, targetId uint) (bool, error) {
	list, err := s.getMyList(ctx)
	if err != nil {
		return false, err
	}
	for _, v := range list {
		if v.TargetId == targetId && v.TargetType == targetType && v.Type == consts.InteractTypeZan {
			return true, nil
		}
	}
	return false, nil
}

// 踩
func (s *serviceInteract) Cai(ctx context.Context, targetType string, targetId uint) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		customCtx := Context.Get(ctx)
		if customCtx == nil || customCtx.User == nil {
			return nil
		}
		r, err := dao.Interact.Ctx(ctx).Data(&entity.Interact{
			UserId:     customCtx.User.Id,
			TargetId:   targetId,
			TargetType: targetType,
			Type:       consts.InteractTypeCai,
		}).FieldsEx(dao.Interact.Columns.Id).InsertIgnore()
		if err != nil {
			return err
		}
		if n, _ := r.RowsAffected(); n == 0 {
			return gerror.New("您已经踩过啦")
		}
		return s.updateCount(ctx, consts.InteractTypeCai, targetType, targetId, 1)
	})
}

// 取消踩
func (s *serviceInteract) CancelCai(ctx context.Context, targetType string, targetId uint) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		customCtx := Context.Get(ctx)
		if customCtx == nil || customCtx.User == nil {
			return nil
		}
		r, err := dao.Interact.Ctx(ctx).Where(g.Map{
			dao.Interact.Columns.UserId:     Context.Get(ctx).User.Id,
			dao.Interact.Columns.TargetId:   targetId,
			dao.Interact.Columns.TargetType: targetType,
			dao.Interact.Columns.Type:       consts.InteractTypeCai,
		}).Delete()
		if err != nil {
			return err
		}
		if n, _ := r.RowsAffected(); n == 0 {
			return nil
		}
		return s.updateCount(ctx, consts.InteractTypeCai, targetType, targetId, -1)
	})
}

// 我是否有对指定内容踩
func (s *serviceInteract) DidICai(ctx context.Context, targetType string, targetId uint) (bool, error) {
	list, err := s.getMyList(ctx)
	if err != nil {
		return false, err
	}
	for _, v := range list {
		if v.TargetId == targetId && v.TargetType == targetType && v.Type == consts.InteractTypeCai {
			return true, nil
		}
	}
	return false, nil
}

// 获得我的互动数据列表，内部带请求上下文缓存
func (s *serviceInteract) getMyList(ctx context.Context) ([]*entity.Interact, error) {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil, nil
	}
	if v, ok := customCtx.Data[contextMapKeyForMyInteractList]; ok {
		return v.([]*entity.Interact), nil
	}
	var list []*entity.Interact
	err := dao.Interact.Ctx(ctx).Where(dao.Interact.Columns.UserId, customCtx.User.Id).Scan(&list)
	if err != nil {
		return nil, err
	}
	customCtx.Data[contextMapKeyForMyInteractList] = list
	return list, err
}

// 根据业务类型更新指定模块的赞/踩数量
func (s *serviceInteract) updateCount(ctx context.Context, interactType int, targetType string, targetId uint, count int) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		defer func() {
			// 清空上下文对应的互动数据缓存
			if customCtx := Context.Get(ctx); customCtx != nil {
				delete(customCtx.Data, contextMapKeyForMyInteractList)
			}
		}()

		var err error
		switch targetType {
		// 内容赞踩
		case consts.InteractTargetTypeContent:
			switch interactType {
			case consts.InteractTypeZan:
				_, err = dao.Content.Ctx(ctx).
					Where(dao.Content.Columns.Id, targetId).
					WhereGTE(dao.Content.Columns.ZanCount, 0).
					Increment(dao.Content.Columns.ZanCount, count)
				if err != nil {
					return err
				}

			case consts.InteractTypeCai:
				_, err = dao.Content.Ctx(ctx).
					Where(dao.Content.Columns.Id, targetId).
					WhereGTE(dao.Content.Columns.CaiCount, 0).
					Increment(dao.Content.Columns.CaiCount, count)
				if err != nil {
					return err
				}
			}
		// 评论赞踩
		case consts.InteractTargetTypeReply:
			switch interactType {
			case consts.InteractTypeZan:
				_, err = dao.Reply.Ctx(ctx).
					Where(dao.Content.Columns.Id, targetId).
					WhereGTE(dao.Content.Columns.ZanCount, 0).
					Increment(dao.Content.Columns.ZanCount, count)
				if err != nil {
					return err
				}

			case consts.InteractTypeCai:
				_, err = dao.Reply.Ctx(ctx).
					Where(dao.Content.Columns.Id, targetId).
					WhereGTE(dao.Content.Columns.CaiCount, 0).
					Increment(dao.Content.Columns.CaiCount, count)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
