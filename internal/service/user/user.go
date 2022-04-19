package user

import (
	"context"
	"fmt"

	"focus-single/internal/model/do"
	"focus-single/internal/service/bizctx"
	"focus-single/internal/service/session"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/o1egl/govatar"

	"focus-single/internal/dao"
	"focus-single/internal/model"
	"focus-single/internal/model/entity"
)

const (
	avatarUploadUrlPrefix = `/upload/avatar`
)

var (
	avatarUploadPath = g.Cfg().MustGet(gctx.New(), `upload.path`).String() + `/avatar`
)

func init() {
	// 启动时创建头像存储目录
	if !gfile.Exists(avatarUploadPath) {
		if err := gfile.Mkdir(avatarUploadPath); err != nil {
			g.Log().Fatal(gctx.New(), err)
		}
	}
}

// 执行登录
func Login(ctx context.Context, in model.UserLoginInput) error {
	userEntity, err := GetUserByPassportAndPassword(
		ctx,
		in.Passport,
		EncryptPassword(in.Passport, in.Password),
	)
	if err != nil {
		return err
	}
	if userEntity == nil {
		return gerror.New(`账号或密码错误`)
	}
	if err = session.SetUser(ctx, userEntity); err != nil {
		return err
	}
	// 自动更新Ctx用户信息
	bizctx.SetUser(ctx, &model.ContextUser{
		Id:       userEntity.Id,
		Passport: userEntity.Passport,
		Nickname: userEntity.Nickname,
		Avatar:   userEntity.Avatar,
	})
	return nil
}

// 注销
func Logout(ctx context.Context) error {
	return session.RemoveUser(ctx)
}

// 将密码按照内部算法进行加密
func EncryptPassword(passport, password string) string {
	return gmd5.MustEncrypt(passport + password)
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
		Password: password,
	}).Scan(&user)
	return
}

// 检测给定的账号是否唯一
func CheckPassportUnique(ctx context.Context, passport string) error {
	n, err := dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`账号"%s"已被占用`, passport)
	}
	return nil
}

// 检测给定的昵称是否唯一
func CheckNicknameUnique(ctx context.Context, nickname string) error {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.Newf(`昵称"%s"已被占用`, nickname)
	}
	return nil
}

// 用户注册。
func Register(ctx context.Context, in model.UserRegisterInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var user *entity.User
		if err := gconv.Struct(in, &user); err != nil {
			return err
		}
		if err := CheckPassportUnique(ctx, user.Passport); err != nil {
			return err
		}
		if err := CheckNicknameUnique(ctx, user.Nickname); err != nil {
			return err
		}
		user.Password = EncryptPassword(user.Passport, user.Password)
		// 自动生成头像
		avatarFilePath := fmt.Sprintf(`%s/%s.jpg`, avatarUploadPath, user.Passport)
		if err := govatar.GenerateFileForUsername(govatar.MALE, user.Passport, avatarFilePath); err != nil {
			return gerror.Wrapf(err, `自动创建头像失败`)
		}
		user.Avatar = fmt.Sprintf(`%s/%s.jpg`, avatarUploadUrlPrefix, user.Passport)
		_, err := dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err
	})
}

// 修改个人密码
func UpdatePassword(ctx context.Context, in model.UserPasswordInput) error {
	oldPassword := EncryptPassword(bizctx.Get(ctx).User.Passport, in.OldPassword)
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Id:       bizctx.Get(ctx).User.Id,
		Password: oldPassword,
	}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New(`原始密码错误`)
	}
	newPassword := EncryptPassword(bizctx.Get(ctx).User.Passport, in.NewPassword)
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Password: newPassword,
	}).Where(do.User{
		Id: bizctx.Get(ctx).User.Id,
	}).Update()
	return err
}

// 获取个人信息
func GetProfileById(ctx context.Context, userId uint) (out *model.UserGetProfileOutput, err error) {
	if err = dao.User.Ctx(ctx).WherePri(userId).Scan(&out); err != nil {
		return nil, err
	}
	// 需要判断nil是否存在,不存在需要判断为空,以防后续nil
	if out == nil {
		return nil, nil
	}
	out.Stats, err = GetUserStats(ctx, userId)
	if err != nil {
		return nil, err
	}
	return
}

// 修改个人资料
func GetProfile(ctx context.Context) (*model.UserGetProfileOutput, error) {
	return GetProfileById(ctx, bizctx.Get(ctx).User.Id)
}

func UpdateAvatar(ctx context.Context, in model.UserUpdateAvatarInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var err error
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Avatar: in.Avatar,
		}).Where(do.User{
			Id: in.UserId,
		}).Update()
		return err
	})
}

// 修改个人资料
func UpdateProfile(ctx context.Context, in model.UserUpdateProfileInput) error {
	var (
		err    error
		user   = bizctx.Get(ctx).User
		userId = user.Id
		cls    = dao.User.Columns()
	)
	count, err := dao.User.Ctx(ctx).Where(cls.Nickname, in.Nickname).WhereNot(cls.Id, userId).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.Newf(`昵称"%s"已被占用`, in.Nickname)
	}
	_, err = dao.User.Ctx(ctx).OmitEmpty().Data(do.User{
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Gender:   in.Gender,
	}).Where(cls.Id, userId).Update()
	// 更新登录session Nickname
	if err == nil && user.Nickname != in.Nickname {
		sessionUser := session.GetUser(ctx)
		sessionUser.Nickname = in.Nickname
		err = session.SetUser(ctx, sessionUser)
	}
	return err
}

// 获取文章数量
func GetUserStats(ctx context.Context, userId uint) (stats map[string]int, err error) {
	stats = make(map[string]int)
	// 文章统计
	if err = doContentStats(ctx, userId, stats); err != nil {
		return nil, err
	}
	// 回复统计
	if err = doReplyStats(ctx, userId, stats); err != nil {
		return nil, err
	}
	return
}

func doContentStats(ctx context.Context, userId uint, stats map[string]int) (err error) {
	var (
		cls = dao.Content.Columns()
		orm = dao.Content.Ctx(ctx).Fields(cls.Type, "COUNT(*) total")
	)
	if userId > 0 && !IsAdmin(ctx, userId) {
		orm = orm.Where(cls.UserId, userId)
	}
	result, err := orm.Group(cls.Type).All()
	if err != nil {
		return err
	}
	for _, v := range result {
		var (
			value = v[cls.Type]
			v2    = v["total"]
		)
		stats[value.String()] = v2.Int()
	}
	return nil
}

func doReplyStats(ctx context.Context, userId uint, stats map[string]int) (err error) {
	orm := dao.Reply.Ctx(ctx).Fields("COUNT(*) total")
	if userId > 0 && !IsAdmin(ctx, userId) {
		orm = orm.Where(dao.Reply.Columns().UserId, userId)
	}
	record, err := orm.One()
	if err != nil {
		return err
	}
	stats["message"] = record["total"].Int()
	return nil
}

// 当前用户是否管理员
func IsCtxAdmin(ctx context.Context) bool {
	var ctxUser = bizctx.Get(ctx).User
	return IsAdmin(ctx, ctxUser.Id)
}

// 判断给定用户是否管理员
func IsAdmin(ctx context.Context, userId uint) bool {
	adminIds := g.Cfg().MustGet(ctx, "settings.adminIds").Uints()
	for _, adminId := range adminIds {
		if userId == adminId {
			return true
		}
	}
	return false
}
