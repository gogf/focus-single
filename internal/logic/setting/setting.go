package setting

import (
	"context"

	"focus-single/internal/service"
	"github.com/gogf/gf/v2/frame/g"

	"focus-single/internal/dao"
	"focus-single/internal/model/entity"
)

type sSetting struct{}

func init() {
	service.RegisterSetting(New())
}

func New() *sSetting {
	return &sSetting{}
}

// 设置KV。
func (s *sSetting) Set(ctx context.Context, key, value string) error {
	_, err := dao.Setting.Ctx(ctx).Data(entity.Setting{
		K: key,
		V: value,
	}).Save()
	return err
}

// 查询KV。
func (s *sSetting) Get(ctx context.Context, key string) (string, error) {
	v, err := s.GetVar(ctx, key)
	return v.String(), err
}

// 查询KV，返回泛型，便于转换。
func (s *sSetting) GetVar(ctx context.Context, key string) (*g.Var, error) {
	v, err := dao.Setting.Ctx(ctx).Fields(dao.Setting.Columns().V).Where(dao.Setting.Columns().K, key).Value()
	return v, err
}
