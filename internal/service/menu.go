package service

import (
	"context"
	"encoding/json"

	"focus-single/internal/model"
)

// Menu 菜单管理服务
var Menu = serviceMenu{}

type serviceMenu struct{}

const (
	settingTopMenusKey = "TopMenus"
)

// SetTopMenus 获取顶部菜单
func (s *serviceMenu) SetTopMenus(ctx context.Context, menus []*model.MenuItem) error {
	b, err := json.Marshal(menus)
	if err != nil {
		return err
	}
	return Setting.Set(ctx, settingTopMenusKey, string(b))
}

// GetTopMenus 获取顶部菜单
func (s *serviceMenu) GetTopMenus(ctx context.Context) ([]*model.MenuItem, error) {
	var topMenus []*model.MenuItem
	v, err := Setting.GetVar(ctx, settingTopMenusKey)
	if err != nil {
		return nil, err
	}
	err = v.Structs(&topMenus)
	return topMenus, err
}

// GetTopMenuByUrl 根据给定的Url检索顶部菜单，给定的Url可能只是一个Url Path。
func (s *serviceMenu) GetTopMenuByUrl(ctx context.Context, url string) (*model.MenuItem, error) {
	items, _ := s.GetTopMenus(ctx)
	for _, v := range items {
		if v.Url == url {
			return v, nil
		}
	}
	return nil, nil
}
