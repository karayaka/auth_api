package customerrors

type UnAuthorizedError struct {
	Msg string
}

func NewUnAuthorizedError(msg string) error {
	return &CustomError{
		Msg: msg,
	}
}

func (c *UnAuthorizedError) Error() string {
	return c.Msg
}
