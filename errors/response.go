package errors

type ResponseError struct {
	Code    int
	Message string
	Status  int
	ERR     error
}

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

func UnWrapResponse(err error) *ResponseError {
	if v, ok := err.(*ResponseError); ok {
		return v
	}
	return nil
}
func WrapResponse(err error, code int, status int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code: code,
	}
}
