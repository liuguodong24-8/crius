package pkgs

const (
	// Success 成功
	Success int32 = iota
	// ErrUnprocessableEntity 请求参数错误
	ErrUnprocessableEntity
	// ErrNotFound 数据未找到
	ErrNotFound
	// ErrInternal 内部错误
	ErrInternal
	// ErrPassword 密码错误
	ErrPassword
	// ErrNumberLimit 数量限制
	ErrNumberLimit
	// ErrNotice 需要返回的提示校验信息
	ErrCheck
)
