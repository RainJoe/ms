package gateway

type ErrCode int

var ErrMessages = map[ErrCode]string {
	ErrValidateParams: "参数错误",
	ErrFailedToCallUserService: "调用用户服务失败",

}

const (
	ErrValidateParams ErrCode = 10001 + iota
	ErrFailedToCallUserService
)
