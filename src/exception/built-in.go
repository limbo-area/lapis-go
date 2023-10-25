package exception

func BadRequestException() *HttpException {
	return ThrowException(HtptpExceptionOption{
		Code:    BadRequest,
		Message: "Bad Request",
	}, 400)
}

func UnauthorizedException() *HttpException {
	return ThrowException(HtptpExceptionOption{
		Code:    Unauthorized,
		Message: "Unauthorized, please try to login again",
	}, 401)
}

func ForbiddenException() *HttpException {
	return ThrowException(HtptpExceptionOption{
		Code:    Forbidden,
		Message: "You don't have permission to access this",
	}, 403)
}

func NotFoundException(model string) *HttpException {
	message := "Not found " + model
	return ThrowException(HtptpExceptionOption{
		Code:    DataNofFound,
		Message: message,
	}, 200)
}

func ExistedException(model string) *HttpException {
	message := "Already existed" + model
	return ThrowException(HtptpExceptionOption{
		Code:    ExistedData,
		Message: message,
	}, 200)
}

func NotAllowedException() *HttpException {
	return ThrowException(HtptpExceptionOption{
		Code:    NotAllowedAction,
		Message: "Action not allowed",
	}, 200)
}
