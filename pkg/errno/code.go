package errno

/**
	10001   1 代表系统级错误  00 代表业务模块
    20101  2 业务模块错误    01 代表模块 01 具体某一处
*/
var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation       = &Errno{Code: 20001, Message: "参数验证没通过"}
	ErrDatabase         = &Errno{Code: 20002, Message: ""}
	ErrToken            = &Errno{Code: 20003, Message: ""}
	ErrJaegerInit       = &Errno{Code: 20004, Message: "jaeger init error"}
	ErrDBNotFoundRecord = &Errno{Code: 20005, Message: "没有找到该数据"}
	ErrTokenInvalid     = &Errno{Code: 20006, Message: "TOKEN无效"}
	ErrAuthInvalid      = &Errno{Code: 20007, Message: "权限不足"}

	ErrUserNameNotUnique  = &Errno{Code: 30001, Message: "用户名已存在"}
	ErrUserNameOrPassword = &Errno{Code: 30002, Message: "用户名或密码错误"}
	ErrUserFreeze         = &Errno{Code: 30003, Message: "该用户已冻结"}
)
