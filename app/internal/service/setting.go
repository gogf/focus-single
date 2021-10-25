package service

import (
	"context"
	"focus/app/internal/dao"
	"focus/app/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// 持久化Key-Value管理服务
var Setting = settingService{}

type settingService struct{}

// 设置KV。
func (s *settingService) Set(ctx context.Context, key, value string) error {
	_, err := dao.Setting.Ctx(ctx).Data(model.Setting{
		K: key,
		V: value,
	}).Save()
	return err
}

// 查询KV。
func (s *settingService) Get(ctx context.Context, key string) (string, error) {
	v, err := s.GetVar(ctx, key)
	return v.String(), err
}

// 查询KV，返回泛型，便于转换。
func (s *settingService) GetVar(ctx context.Context, key string) (*g.Var, error) {
	v, err := dao.Setting.Ctx(ctx).Fields(dao.Setting.Columns.V).Where(dao.Setting.Columns.K, key).Value()
	return v, err
}
