package errno

/**
	1010001  10 代表服务 1 代表系统级错误  00 代表业务模块
    1020101  10 代表服务 2 业务模块错误    01 代表模块 01 具体某一处
*/
var (
	// Common errors
	OK                  = &Errno{Code: 200, Message: "OK"}
	InternalServerError = &Errno{Code: 1110001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 1110002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation       = &Errno{Code: 1120001, Message: "参数验证没通过"}
	ErrDatabase         = &Errno{Code: 1120002, Message: ""}
	ErrToken            = &Errno{Code: 1120003, Message: ""}
	ErrJaegerInit       = &Errno{Code: 1120004, Message: "jaeger init error"}
	ErrDBNotFoundRecord = &Errno{Code: 1120005, Message: "没有找到该数据"}
)
