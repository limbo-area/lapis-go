package exception

type HttpException struct {
	Code       ExceptioCode
	Message    string
	StatusCode int
}

type HtptpExceptionOption struct {
	Code    ExceptioCode
	Message string
}

type ExceptioCode int32

const (
	BadRequest       ExceptioCode = 400
	Unauthorized     ExceptioCode = 401
	Forbidden        ExceptioCode = 403
	NotFound         ExceptioCode = 404
	DataNofFound     ExceptioCode = 4004
	NotAllowedAction ExceptioCode = 4005
	ExistedData      ExceptioCode = 4009
	InternalServer   ExceptioCode = 500
)

func ThrowException(option HtptpExceptionOption, statusCode int) *HttpException {
	return &HttpException{
		Code:       option.Code,
		Message:    option.Message,
		StatusCode: statusCode,
	}
}

func (http *HttpException) Error() string {
	return http.Message
}
