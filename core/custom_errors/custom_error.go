package customerrors

type CustomError struct {
	Msg string
}

func NewCustomError(msg string) error {
	return &CustomError{
		Msg: msg,
	}
}

func (c *CustomError) Error() string {
	return c.Msg
}
