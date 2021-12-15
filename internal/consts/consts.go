package consts

const (
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	ContextKeyRequestId      = "RequestId"          // 用于自定义Context RequestId打印
	DefaultAdminId           = 1                    // 默认管理员ID
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
)
