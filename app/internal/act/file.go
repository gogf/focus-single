package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// 文件管理
	File = fileAct{}
)

type fileAct struct{}

func (a *fileAct) Upload(ctx context.Context, req *api.FileUploadReq) (res *api.FileUploadRes, err error) {
	var (
		request = g.RequestFromCtx(ctx)
		file    = request.GetUploadFile("file")
	)
	if file == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	result, err := service.File.Upload(ctx, model.FileUploadInput{
		File:       file,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	res = &api.FileUploadRes{
		Name: result.Name,
		Url:  result.Url,
	}
	return
}
