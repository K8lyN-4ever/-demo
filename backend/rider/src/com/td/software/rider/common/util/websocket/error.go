package websocket

type IServiceError interface {
	GetErrorCode() int
	GetParams() []interface{}
}

func NewServiceError(errorCode int, params ...interface{}) *serviceError {
	return &serviceError{
		errorCode: errorCode,
		params:    params,
	}
}

type serviceError struct {
	errorCode int
	params    []interface{}
}

func (se serviceError) GetErrorCode() int {
	return se.errorCode
}

func (se serviceError) GetParams() []interface{} {
	return se.params
}
