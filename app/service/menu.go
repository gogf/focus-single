package service

import (
	"context"
	"encoding/json"
	"focus/app/model"
)

// 菜单管理服务
var Menu = menuService{}

type menuService struct{}

const (
	settingTopMenusKey = "TopMenus"
)

// 获取顶部菜单
func (s *menuService) SetTopMenus(ctx context.Context, menus []*model.MenuItem) error {
	b, err := json.Marshal(menus)
	if err != nil {
		return err
	}
	return Setting.Set(ctx, settingTopMenusKey, string(b))
}

// 获取顶部菜单
func (s *menuService) GetTopMenus(ctx context.Context) ([]*model.MenuItem, error) {
	var topMenus []*model.MenuItem
	v, err := Setting.GetVar(ctx, settingTopMenusKey)
	if err != nil {
		return nil, err
	}
	err = v.Structs(&topMenus)
	return topMenus, err
}

// 根据给定的Url检索顶部菜单，给定的Url可能只是一个Url Path。
func (s *menuService) GetTopMenuByUrl(ctx context.Context, url string) (*model.MenuItem, error) {
	items, _ := s.GetTopMenus(ctx)
	for _, v := range items {
		if v.Url == url {
			return v, nil
		}
	}
	return nil, nil
}
