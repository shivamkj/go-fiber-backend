package errors

type HttpError struct {
	Code     int32
	ErrorMsg string
	UserMsg  string
}

func (err *HttpError) Error() string {
	return err.ErrorMsg
}

type InternalHttpError struct {
	HttpError
	err   error
	stack *stack
}

func (err *InternalHttpError) Error() string {
	return err.err.Error()
}

func (err *InternalHttpError) Stack() string {
	return err.stack.StackTrace()
}

func InternalError(msg string, err error) error {
	return &InternalHttpError{
		err:   err,
		stack: callers(),
		HttpError: HttpError{
			Code:     500,
			ErrorMsg: msg,
			UserMsg:  "Internal Server Error",
		}}
}

func BadRequest(msg string) error {
	return &HttpError{Code: 400, ErrorMsg: msg}
}

func Unauthorised(msg string) error {
	return &HttpError{Code: 401, ErrorMsg: msg}
}

func NotFound(msg string) error {
	return &HttpError{Code: 404, ErrorMsg: msg}
}

func Conflict(msg string) error {
	return &HttpError{Code: 409, ErrorMsg: msg}
}
