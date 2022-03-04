package service

type sTag struct{}

// 标签管理服务
func Tag() *sTag {
	return &sTag{}
}
