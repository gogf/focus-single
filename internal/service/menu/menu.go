package menu

import (
	"context"
	"encoding/json"

	"focus-single/internal/service/setting"
)

type Item struct {
	Name   string  // 显示名称
	Url    string  // 链接地址
	Icon   string  // 图标，可能是class，也可能是iconfont
	Target string  // 打开方式: 空, _blank
	Active bool    // 是否被选中
	Items  []*Item // 子级菜单
}

const settingTopMenusKey = "TopMenus"

// 获取顶部菜单
func SetTopMenus(ctx context.Context, menus []*Item) error {
	b, err := json.Marshal(menus)
	if err != nil {
		return err
	}
	return setting.Set(ctx, settingTopMenusKey, string(b))
}

// 获取顶部菜单
func GetTopMenus(ctx context.Context) ([]*Item, error) {
	var topMenus []*Item
	v, err := setting.Get(ctx, settingTopMenusKey)
	if err != nil {
		return nil, err
	}
	err = v.Structs(&topMenus)
	return topMenus, err
}

// 根据给定的Url检索顶部菜单，给定的Url可能只是一个Url Path。
func GetTopMenuByUrl(ctx context.Context, url string) (*Item, error) {
	items, _ := GetTopMenus(ctx)
	for _, v := range items {
		if v.Url == url {
			return v, nil
		}
	}
	return nil, nil
}
