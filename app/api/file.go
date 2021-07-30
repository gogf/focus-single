package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// 文件管理
var File = fileApi{}

type fileApi struct{}

// @summary 上传文件
// @tags    前台-文件
// @produce json
// @param   file formData file true "文件域"
// @router  /file/upload [POST]
// @success 200 {object} internal.FileUploadRes "请求结果"
func (a *fileApi) Upload(ctx context.Context) (res *internal.FileUploadRes, err error) {
	var (
		request = g.RequestFromCtx(ctx)
		file    = request.GetUploadFile("file")
	)
	if file == nil {
		return nil, gerror.NewCode(gerror.CodeMissingParameter, "请选择需要上传的文件")
	}
	result, err := service.File.Upload(ctx, model.FileUploadInput{
		File:       file,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	res = &internal.FileUploadRes{
		Name: result.Name,
		Url:  result.Url,
	}
	return
}
