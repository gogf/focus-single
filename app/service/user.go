package service

import (
	"context"
	"fmt"
	"focus/app/cnt"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"github.com/o1egl/govatar"
)

// 用户管理服务
var User = userService{
	avatarUploadPath:      g.Cfg().GetString(`upload.path`) + `/avatar`,
	avatarUploadUrlPrefix: `/upload/avatar`,
}

type userService struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

func init() {
	// 启动时创建头像存储目录
	if !gfile.Exists(User.avatarUploadPath) {
		if err := gfile.Mkdir(User.avatarUploadPath); err != nil {
			g.Log().Fatal(err)
		}
	}
}

// 获得头像上传路径
func (s *userService) GetAvatarUploadPath() string {
	return s.avatarUploadPath
}

// 获得头像上传对应的URL前缀
func (s *userService) GetAvatarUploadUrlPrefix() string {
	return s.avatarUploadUrlPrefix
}

// 执行登录
func (s *userService) Login(ctx context.Context, input model.UserLoginInput) error {
	userEntity, err := s.GetUserByPassportAndPassword(
		ctx,
		input.Passport,
		s.EncryptPassword(input.Passport, input.Password),
	)
	if err != nil {
		return err
	}
	if userEntity == nil {
		return gerror.New(`账号或密码错误`)
	}
	if err := Session.SetUser(ctx, userEntity); err != nil {
		return err
	}
	// 自动更新上线
	Context.SetUser(ctx, &model.ContextUser{
		Id:       userEntity.Id,
		Passport: userEntity.Passport,
		Nickname: userEntity.Nickname,
		Avatar:   userEntity.Avatar,
	})
	return nil
}

// 注销
func (s *userService) Logout(ctx context.Context) error {
	return Session.RemoveUser(ctx)
}

// 将密码按照内部算法进行加密
func (s *userService) EncryptPassword(passport, password string) string {
	return gmd5.MustEncrypt(passport + password)
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (s *userService) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *model.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns.Passport: passport,
		dao.User.Columns.Password: password,
	}).Scan(&user)
	return
}

// 检测给定的账号是否唯一
func (s *userService) CheckPassportUnique(ctx context.Context, passport string) error {
	n, err := dao.User.Ctx(ctx).Where(dao.User.Columns.Passport, passport).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`账号"%s"已被占用`, passport)
	}
	return nil
}

// 检测给定的昵称是否唯一
func (s *userService) CheckNicknameUnique(ctx context.Context, nickname string) error {
	n, err := dao.User.Ctx(ctx).Where(dao.User.Columns.Nickname, nickname).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`昵称"%s"已被占用`, nickname)
	}
	return nil
}

// 用户注册。
func (s *userService) Register(ctx context.Context, input model.UserRegisterInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var user *model.User
		if err := gconv.Struct(input, &user); err != nil {
			return err
		}
		if err := s.CheckPassportUnique(ctx, user.Passport); err != nil {
			return err
		}
		if err := s.CheckNicknameUnique(ctx, user.Nickname); err != nil {
			return err
		}
		user.Password = s.EncryptPassword(user.Passport, user.Password)
		// 自动生成头像
		avatarFilePath := fmt.Sprintf(`%s/%s.jpg`, s.avatarUploadPath, user.Passport)
		if err := govatar.GenerateFileForUsername(govatar.MALE, user.Passport, avatarFilePath); err != nil {
			return gerror.Wrapf(err, `自动创建头像失败`)
		}
		user.Avatar = fmt.Sprintf(`%s/%s.jpg`, s.avatarUploadUrlPrefix, user.Passport)
		_, err := dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err
	})
}

// 修改个人密码
func (s *userService) UpdatePassword(ctx context.Context, input model.UserPasswordInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		oldPassword := s.EncryptPassword(Context.Get(ctx).User.Passport, input.OldPassword)
		n, err := dao.User.Ctx(ctx).
			Where(dao.User.Columns.Password, oldPassword).
			Where(dao.User.Columns.Id, Context.Get(ctx).User.Id).
			Count()
		if err != nil {
			return err
		}
		if n == 0 {
			return gerror.New(`原始密码错误`)
		}
		newPassword := s.EncryptPassword(Context.Get(ctx).User.Passport, input.NewPassword)
		_, err = dao.User.Ctx(ctx).Data(g.Map{
			dao.User.Columns.Password: newPassword,
		}).Where(dao.User.Columns.Id, Context.Get(ctx).User.Id).Update()
		return err
	})
}

// 获取个人信息
func (s *userService) GetProfileById(ctx context.Context, userId uint) (output *model.UserGetProfileOutput, err error) {
	output = &model.UserGetProfileOutput{}
	if err = dao.User.Ctx(ctx).WherePri(userId).Scan(output); err != nil {
		return nil, err
	}
	output.Stats, err = s.GetUserStats(ctx, userId)
	if err != nil {
		return nil, err
	}
	return
}

// 修改个人资料
func (s *userService) GetProfile(ctx context.Context) (*model.UserGetProfileOutput, error) {
	return s.GetProfileById(ctx, Context.Get(ctx).User.Id)
}

// 修改个人头像
func (s *userService) UpdateAvatar(ctx context.Context, input model.UserUpdateProfileInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var (
			err    error
			user   = Context.Get(ctx).User
			userId = user.Id
		)
		_, err = dao.User.Ctx(ctx).Data(model.UserUpdateAvatarInput{
			Avatar: input.Avatar,
		}).Where(dao.User.Columns.Id, userId).Update()
		// 更新登录session Avatar
		if err == nil && user.Avatar != input.Avatar {
			sessionUser := Session.GetUser(ctx)
			sessionUser.Avatar = input.Avatar
			err = Session.SetUser(ctx, sessionUser)
		}
		return err
	})
}

// 修改个人资料
func (s *userService) UpdateProfile(ctx context.Context, input model.UserUpdateProfileInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var (
			err    error
			user   = Context.Get(ctx).User
			userId = user.Id
		)
		n, err := dao.User.Ctx(ctx).
			Where(dao.User.Columns.Nickname, input.Nickname).
			WhereNot(dao.User.Columns.Id, userId).
			Count()
		if err != nil {
			return err
		}
		if n > 0 {
			return gerror.Newf(`昵称"%s"已被占用`, input.Nickname)
		}
		_, err = dao.User.Ctx(ctx).OmitEmpty().Data(input).Where(dao.User.Columns.Id, userId).Update()
		// 更新登录session Nickname
		if err == nil && user.Nickname != input.Nickname {
			sessionUser := Session.GetUser(ctx)
			sessionUser.Nickname = input.Nickname
			err = Session.SetUser(ctx, sessionUser)
		}
		return err
	})

}

// 禁用指定用户
func (s *userService) Disable(ctx context.Context, id uint) error {
	_, err := dao.User.Ctx(ctx).
		Data(dao.User.Columns.Status, cnt.UserStatusDisabled).
		Where(dao.User.Columns.Id, id).
		Update()
	return err
}

// 查询用户内容列表及用户信息
func (s *userService) GetList(ctx context.Context, input model.UserGetListInput) (output *model.UserGetListOutput, err error) {
	output = &model.UserGetListOutput{}
	// 内容列表
	output.Content, err = Content.GetList(ctx, input.ContentGetListInput)
	if err != nil {
		return output, err
	}
	// 用户信息
	output.User, err = User.GetProfileById(ctx, input.UserId)
	if err != nil {
		return output, err
	}
	// 统计信息
	output.Stats, err = s.GetUserStats(ctx, input.UserId)
	if err != nil {
		return output, err
	}
	return
}

// 消息列表
func (s *userService) GetMessageList(ctx context.Context, input model.UserGetMessageListInput) (output *model.UserGetMessageListOutput, err error) {
	output = &model.UserGetMessageListOutput{
		Page: input.Page,
		Size: input.Size,
	}
	var (
		userId = Context.Get(ctx).User.Id
	)
	// 管理员看所有的
	if !s.IsAdminShow(ctx, userId) {
		input.UserId = userId
	}

	replyList, err := Reply.GetList(ctx, model.ReplyGetListInput{
		Page:       input.Page,
		Size:       input.Size,
		TargetType: input.TargetType,
		TargetId:   input.TargetId,
		UserId:     input.UserId,
	})
	if err != nil {
		return nil, err
	}
	output.List = replyList.List
	output.Stats, err = s.GetUserStats(ctx, userId)
	if err != nil {
		return nil, err
	}
	return
}

// 获取文章数量
func (s *userService) GetUserStats(ctx context.Context, userId uint) (map[string]int, error) {
	// 文章统计
	m := dao.Content.Ctx(ctx).Fields(dao.Content.Columns.Type, "count(*) total")
	if userId > 0 && !s.IsAdminShow(ctx, userId) {
		m = m.Where(dao.Content.Columns.UserId, userId)
	}
	statsModel := m.Group(dao.Content.Columns.Type)
	statsAll, err := statsModel.All()
	if err != nil {
		return nil, err
	}
	statsMap := make(map[string]int)
	for _, v := range statsAll {
		statsMap[v["type"].String()] = v["total"].Int()
	}
	// 回复统计
	replyModel := dao.Reply.Ctx(ctx).Fields("count(*) total")
	if userId > 0 && !s.IsAdminShow(ctx, userId) {
		replyModel = replyModel.Where(dao.Reply.Columns.UserId, userId)
	}
	record, err := replyModel.One()
	if err != nil {
		return nil, err
	}
	statsMap["message"] = record["total"].Int()

	return statsMap, nil
}

// 是否是访问管理员的数据
func (s *userService) IsAdminShow(ctx context.Context, userId uint) bool {
	c := Context.Get(ctx)
	if c == nil {
		return false
	}
	if c.User == nil {
		return false
	}
	if userId != c.User.Id {
		return false
	}
	return c.User.IsAdmin
}
