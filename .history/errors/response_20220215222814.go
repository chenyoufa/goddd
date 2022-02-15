package errors

import "fmt"

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

func WrapResponse(err error, code, status int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:    code,
		Message: fmt.Sprintf(msg, args...),
		ERR:     err,
		Status:  status,
	}
	return res
}
func Wrap400Response(err error, msg string, args ...interface{}) error {
	return Wrap400Response(err, 0, 400, msg, args...)
}
