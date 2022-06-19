package errors

type HTTPError struct {
	status  int
	message string
}

func (e HTTPError) HTTPStatusCode() int {
	return e.status
}

func (e HTTPError) Error() string {
	return e.message
}

type BadRequestError struct {
	HTTPError
}

func NewBadRequestError(msg string) BadRequestError {
	e := HTTPError{status: 400, message: msg}
	return BadRequestError{e}
}
