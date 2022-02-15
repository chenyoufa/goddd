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
