package model

// MenuItem 菜单数据结构
type MenuItem struct {
	Name   string      // 显示名称
	Url    string      // 链接地址
	Icon   string      // 图标，可能是class，也可能是iconfont
	Target string      // 打开方式: 空, _blank
	Active bool        // 是否被选中
	Items  []*MenuItem // 子级菜单
}
