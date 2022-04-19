package menu

import (
	"context"
	"encoding/json"

	"focus-single/internal/model"
	"focus-single/internal/service/setting"
)

const settingTopMenusKey = "TopMenus"

// 获取顶部菜单
func SetTopMenus(ctx context.Context, menus []*model.MenuItem) error {
	b, err := json.Marshal(menus)
	if err != nil {
		return err
	}
	return setting.Set(ctx, settingTopMenusKey, string(b))
}

// 获取顶部菜单
func GetTopMenus(ctx context.Context) ([]*model.MenuItem, error) {
	var topMenus []*model.MenuItem
	v, err := setting.Get(ctx, settingTopMenusKey)
	if err != nil {
		return nil, err
	}
	err = v.Structs(&topMenus)
	return topMenus, err
}

// 根据给定的Url检索顶部菜单，给定的Url可能只是一个Url Path。
func GetTopMenuByUrl(ctx context.Context, url string) (*model.MenuItem, error) {
	items, _ := GetTopMenus(ctx)
	for _, v := range items {
		if v.Url == url {
			return v, nil
		}
	}
	return nil, nil
}
