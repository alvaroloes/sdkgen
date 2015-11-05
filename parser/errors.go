package parser

// Error codes
const (
	ErrorCodeNoRootResource = iota + 1
)

type Error struct {
	Code int
	Context, Message string
}

func (err Error) Error() string{
	return err.Context + " -> " + err.Message
}

func NewError(code int, message string) Error {
	return Error {
		Code: code,
		Message: message,
	}
}
