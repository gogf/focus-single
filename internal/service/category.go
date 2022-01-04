package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"

	"focus-single/internal/model"
	"focus-single/internal/model/entity"
	"focus-single/internal/service/internal/dao"
)

type sCategory struct{}

var insCategory = sCategory{}

// Category 栏目管理服务
func Category() *sCategory {
	return &insCategory
}

const (
	mapCacheKey       = "category_map_cache"
	mapCacheDuration  = time.Hour
	treeCacheKey      = "category_tree_cache"
	treeCacheDuration = time.Hour
)

// GetTree 查询列表
func (s *sCategory) GetTree(ctx context.Context, contentType string) ([]*model.CategoryTreeItem, error) {
	// 缓存控制
	var (
		cacheKey  = treeCacheKey + contentType
		cacheFunc = func(ctx context.Context) (interface{}, error) {
			entities, err := s.GetList(ctx)
			if err != nil {
				return nil, err
			}
			tree, err := s.formTree(0, contentType, entities)
			if err != nil {
				return nil, err
			}
			return tree, nil
		}
	)
	v, err := gcache.GetOrSetFunc(ctx, cacheKey, cacheFunc, treeCacheDuration)
	if err != nil {
		return nil, err
	}
	var (
		result []*model.CategoryTreeItem
	)
	err = v.Scan(&result)
	return result, err
}

// GetSubIdList 获取指定栏目ID及其下面所有子ID，构成数组返回。
// 注意，返回的ID列表中包含查询的栏目ID.
func (s *sCategory) GetSubIdList(ctx context.Context, id uint) ([]uint, error) {
	m, err := s.GetMap(ctx)
	if err != nil {
		return nil, err
	}
	entity := m[id]
	if entity == nil {
		return nil, gerror.Newf(`%d栏目不存在`, id)
	}
	tree, err := s.GetTree(ctx, entity.ContentType)
	if err != nil {
		return nil, err
	}
	return append([]uint{id}, s.getSubIdListByTree(id, tree)...), nil
}

// 递归获取指定栏目ID下的所有子级
func (s *sCategory) getSubIdListByTree(id uint, trees []*model.CategoryTreeItem) []uint {
	idArray := make([]uint, 0)
	for _, item := range trees {
		if item.ParentId == id {
			idArray = append(idArray, item.Id)
			if len(item.Items) > 0 {
				idArray = append(idArray, s.getSubIdListByTree(item.Id, item.Items)...)
			}
		} else if len(item.Items) > 0 {
			idArray = append(idArray, s.getSubIdListByTree(id, item.Items)...)
		}
	}
	return idArray
}

// 构造树形栏目列表。
func (s *sCategory) formTree(parentId uint, contentType string, entities []*entity.Category) ([]*model.CategoryTreeItem, error) {
	tree := make([]*model.CategoryTreeItem, 0)
	for _, entity := range entities {
		if contentType != "" && entity.ContentType != contentType {
			continue
		}
		if entity.ParentId == parentId {
			subTree, err := s.formTree(entity.Id, contentType, entities)
			if err != nil {
				return nil, err
			}
			item := &model.CategoryTreeItem{
				Items: subTree,
			}
			if err = gconv.Struct(entity, item); err != nil {
				return nil, err
			}
			tree = append(tree, item)
		}
	}
	return tree, nil
}

// GetList 获得所有的栏目列表。
func (s *sCategory) GetList(ctx context.Context) (list []*entity.Category, err error) {
	err = dao.Category.Ctx(ctx).
		OrderAsc(dao.Category.Columns().Sort).
		OrderAsc(dao.Category.Columns().Id).
		Scan(&list)
	return
}

// GetItem 查询单个栏目信息
func (s *sCategory) GetItem(ctx context.Context, id uint) (*entity.Category, error) {
	m, err := s.GetMap(ctx)
	if err != nil {
		return nil, err
	}
	return m[id], nil
}

// GetMap 获得所有的栏目列表，构成Map返回，键名为栏目ID
func (s *sCategory) GetMap(ctx context.Context) (map[uint]*entity.Category, error) {
	var (
		cacheKey  = mapCacheKey
		cacheFunc = func(ctx context.Context) (interface{}, error) {
			entities, err := s.GetList(ctx)
			if err != nil {
				return nil, err
			}
			m := make(map[uint]*entity.Category)
			for _, entity := range entities {
				item := entity
				m[entity.Id] = item
			}
			return m, nil
		}
	)
	v, err := gcache.GetOrSetFunc(ctx, cacheKey, cacheFunc, mapCacheDuration)
	if err != nil {
		return nil, err
	}
	var (
		result map[uint]*entity.Category
	)
	err = v.Scan(&result)
	return result, err
}
