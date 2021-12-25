package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"focus-single/internal/model/entity"
	"focus-single/internal/service/internal/dao"
)

// 持久化Key-Value管理服务
var Setting = serviceSetting{}

type serviceSetting struct{}

// 设置KV。
func (s *serviceSetting) Set(ctx context.Context, key, value string) error {
	_, err := dao.Setting.Ctx(ctx).Data(entity.Setting{
		K: key,
		V: value,
	}).Save()
	return err
}

// 查询KV。
func (s *serviceSetting) Get(ctx context.Context, key string) (string, error) {
	v, err := s.GetVar(ctx, key)
	return v.String(), err
}

// 查询KV，返回泛型，便于转换。
func (s *serviceSetting) GetVar(ctx context.Context, key string) (*g.Var, error) {
	v, err := dao.Setting.Ctx(ctx).Fields(dao.Setting.Columns().V).Where(dao.Setting.Columns().K, key).Value()
	return v, err
}
