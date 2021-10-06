package service

import (
	"context"
	"focus/app/internal/cnt"
	"focus/app/internal/dao"
	"focus/app/internal/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"time"
)

// 文件管理服务
var File = fileService{}

type fileService struct{}

// 同一上传文件
func (s *fileService) Upload(ctx context.Context, input model.FileUploadInput) (*model.FileUploadOutput, error) {
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("上传文件路径配置不存在")
	}
	if input.Name != "" {
		input.File.Filename = input.Name
	}
	// 同一用户1分钟之内只能上传10张图片
	count, err := dao.File.Ctx(ctx).
		Where(dao.File.Columns.UserId, Context.Get(ctx).User.Id).
		WhereGTE(dao.File.Columns.CreatedAt, gtime.Now().Add(time.Minute)).
		Count()
	if err != nil {
		return nil, err
	}
	if count >= cnt.FileMaxUploadCountMinute {
		return nil, gerror.New("您上传得太频繁，请稍后再操作")
	}
	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := input.File.Save(gfile.Join(uploadPath, dateDirName), input.RandomName)
	if err != nil {
		return nil, err
	}
	// 记录到数据表
	data := model.File{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dateDirName, fileName),
		Url:    "/upload/" + dateDirName + "/" + fileName,
		UserId: Context.Get(ctx).User.Id,
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
