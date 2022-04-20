package category

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"

	"focus-single/internal/dao"
	"focus-single/internal/model/entity"
)

// 栏目树形列表
type TreeItem struct {
	Id       uint        `json:"id"`              // 分类ID，自增主键
	ParentId uint        `json:"parent_id"`       // 父级分类ID，用于层级管理
	Name     string      `json:"name"`            // 分类名称
	Thumb    string      `json:"thumb"`           // 封面图
	Brief    string      `json:"brief"`           // 简述
	Content  string      `json:"content"`         // 详细介绍
	Indent   string      `json:"indent"`          // 缩进字符串，包含：&nbsp;, " │", " ├", " └"
	Items    []*TreeItem `json:"items,omitempty"` // 子级数据项
}

const (
	mapCacheKey       = "category_map_cache"
	mapCacheDuration  = time.Hour
	treeCacheKey      = "category_tree_cache"
	treeCacheDuration = time.Hour
)

// GetTree 查询列表
func GetTree(ctx context.Context, contentType string) ([]*TreeItem, error) {
	// 缓存控制
	var (
		cacheKey  = treeCacheKey + contentType
		cacheFunc = func(ctx context.Context) (interface{}, error) {
			entities, err := GetList(ctx)
			if err != nil {
				return nil, err
			}
			tree, err := formTree(0, contentType, entities)
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
	var result []*TreeItem
	err = v.Scan(&result)
	return result, err
}

// GetSubIdList 获取指定栏目ID及其下面所有子ID，构成数组返回。
// 注意，返回的ID列表中包含查询的栏目ID.
func GetSubIdList(ctx context.Context, id uint) ([]uint, error) {
	m, err := GetMap(ctx)
	if err != nil {
		return nil, err
	}
	categoryEntity := m[id]
	if categoryEntity == nil {
		return nil, gerror.Newf(`%d栏目不存在`, id)
	}
	tree, err := GetTree(ctx, categoryEntity.ContentType)
	if err != nil {
		return nil, err
	}
	return append([]uint{id}, getSubIdListByTree(id, tree)...), nil
}

// 递归获取指定栏目ID下的所有子级
func getSubIdListByTree(id uint, trees []*TreeItem) []uint {
	idArray := make([]uint, 0)
	for _, item := range trees {
		if item.ParentId == id {
			idArray = append(idArray, item.Id)
			if len(item.Items) > 0 {
				idArray = append(idArray, getSubIdListByTree(item.Id, item.Items)...)
			}
		} else if len(item.Items) > 0 {
			idArray = append(idArray, getSubIdListByTree(id, item.Items)...)
		}
	}
	return idArray
}

// 构造树形栏目列表。
func formTree(parentId uint, contentType string, entities []*entity.Category) ([]*TreeItem, error) {
	tree := make([]*TreeItem, 0)
	for _, categoryEntity := range entities {
		if contentType != "" && categoryEntity.ContentType != contentType {
			continue
		}
		if categoryEntity.ParentId == parentId {
			subTree, err := formTree(categoryEntity.Id, contentType, entities)
			if err != nil {
				return nil, err
			}
			item := &TreeItem{
				Items: subTree,
			}
			if err = gconv.Struct(categoryEntity, item); err != nil {
				return nil, err
			}
			tree = append(tree, item)
		}
	}
	return tree, nil
}

// GetList 获得所有的栏目列表。
func GetList(ctx context.Context) (list []*entity.Category, err error) {
	var cls = dao.Category.Columns()
	err = dao.Category.Ctx(ctx).OrderAsc(cls.Sort).OrderAsc(cls.Id).Scan(&list)
	return
}

// GetItem 查询单个栏目信息
func GetItem(ctx context.Context, id uint) (*entity.Category, error) {
	m, err := GetMap(ctx)
	if err != nil {
		return nil, err
	}
	return m[id], nil
}

// GetMap 获得所有的栏目列表，构成Map返回，键名为栏目ID
func GetMap(ctx context.Context) (map[uint]*entity.Category, error) {
	var (
		cacheKey  = mapCacheKey
		cacheFunc = func(ctx context.Context) (interface{}, error) {
			entities, err := GetList(ctx)
			if err != nil {
				return nil, err
			}
			m := make(map[uint]*entity.Category)
			for _, categoryEntity := range entities {
				item := categoryEntity
				m[categoryEntity.Id] = item
			}
			return m, nil
		}
	)
	v, err := gcache.GetOrSetFunc(ctx, cacheKey, cacheFunc, mapCacheDuration)
	if err != nil {
		return nil, err
	}
	var result map[uint]*entity.Category
	err = v.Scan(&result)
	return result, err
}
