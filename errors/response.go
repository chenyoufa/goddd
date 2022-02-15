package errors

type ResponseError struct {
	Code    int
	Message string
	Status  int
	ERR     error
}
