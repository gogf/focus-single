package file

import (
	"context"
	"time"

	"focus-single/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"

	"focus-single/internal/consts"
	"focus-single/internal/dao"
	"focus-single/internal/model"
	"focus-single/internal/model/entity"
)

type sFile struct{}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

// 同一上传文件
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (*model.FileUploadOutput, error) {
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("上传文件路径配置不存在")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	// 同一用户1分钟之内只能上传10张图片
	count, err := dao.File.Ctx(ctx).
		Where(dao.File.Columns().UserId, service.BizCtx().Get(ctx).User.Id).
		WhereGTE(dao.File.Columns().CreatedAt, gtime.Now().Add(time.Minute)).
		Count()
	if err != nil {
		return nil, err
	}
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("您上传得太频繁，请稍后再操作")
	}
	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}
	// 记录到数据表
	data := entity.File{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dateDirName, fileName),
		Url:    "/upload/" + dateDirName + "/" + fileName,
		UserId: service.BizCtx().Get(ctx).User.Id,
	}
	result, err := dao.File.Ctx(ctx).Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &model.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Path: data.Src,
		Url:  data.Url,
	}, nil
}
