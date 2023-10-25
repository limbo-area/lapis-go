package enums

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
