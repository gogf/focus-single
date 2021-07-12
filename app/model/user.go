package model

// 创建用户
type UserRegisterInput struct {
	Passport string // 账号
	Password string // 密码(明文)
	Nickname string // 昵称
}

type UserPasswordInput struct {
	OldPassword string // 原密码
	NewPassword string // 新密码
}

// 用户登录
type UserLoginInput struct {
	Passport string // 账号
	Password string // 密码(明文)
}

// 用户信息
type UserGetProfileOutput struct {
	Id       uint           // 用户ID
	Nickname string         // 昵称
	Avatar   string         // 头像地址
	Gender   int            // 性别 0: 未设置 1: 男 2: 女
	Stats    map[string]int // 发布内容数量
}

// 修改用户头像
type UserUpdateAvatarInput struct {
	Avatar string // 头像地址
}

// 修改用户信息
type UserUpdateProfileInput struct {
	Id       uint   // 用户ID
	Nickname string // 昵称
	Avatar   string // 头像地址
	Gender   int    // 性别 0: 未设置 1: 男 2: 女
}

// 查询用户列表输入
type UserGetListInput struct {
	ContentGetListInput
	Id uint
}

// 查询用户详情结果
type UserGetListOutput struct {
	Content *ContentGetListOutput `json:"content"` // 查询用户
	User    *UserGetProfileOutput `json:"user"`    // 查询用户
	Stats   map[string]int        // 发布内容数量
}

type UserGetMessageListInput struct {
	Page       int    `json:"page"`        // 分页码
	Size       int    `json:"size"`        // 分页数量
	TargetType string `json:"target_type"` // 数据类型
	TargetId   int    `json:"target_id"`   // 数据ID
	UserId     uint   `json:"user_id"`     // 用户ID
}

// 查询用户列表查询结果
type UserGetMessageListOutput struct {
	List  []ReplyGetListOutputItem `json:"list"`  // 列表
	Page  int                      `json:"page"`  // 分页码
	Size  int                      `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
	Stats map[string]int           // 发布内容数量
}
