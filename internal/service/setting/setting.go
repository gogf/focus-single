package setting

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"focus-single/internal/dao"
	"focus-single/internal/model/entity"
)

// 设置KV。
func Set(ctx context.Context, key, value string) error {
	_, err := dao.Setting.Ctx(ctx).Data(entity.Setting{
		K: key,
		V: value,
	}).Save()
	return err
}

// 查询KV，返回泛型，便于转换。
func Get(ctx context.Context, key string) (*g.Var, error) {
	v, err := dao.Setting.Ctx(ctx).Fields(dao.Setting.Columns().V).Where(dao.Setting.Columns().K, key).Value()
	return v, err
}
