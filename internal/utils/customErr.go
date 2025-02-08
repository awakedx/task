package utils

type CustomErr struct {
	Msg   string
	Cause error
}

func (c *CustomErr) Error() string {
	return c.Msg
}

func (c *CustomErr) Unwrap() error {
	return c.Cause
}

var (
	InternalError = &CustomErr{Msg: "Internal Server error"}
	NotFoundError = &CustomErr{Msg: "Not Found"}
	BadRequestErr = &CustomErr{Msg: "Bad Request"}
)
