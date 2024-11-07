package customerrors

type NotFoundError struct {
	Msg string
}

func NewNotFoundError(msg string) error {
	return &NotFoundError{
		Msg: msg,
	}
}
func (c *NotFoundError) Error() string {
	return c.Msg
}
