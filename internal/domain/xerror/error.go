package xerror

const (
	ErrNotFound = 404
	ErrInternal = 500
	ErrConflict = 409
)

type Error struct {
	Code       int    `json:"code"`
	DevMessage string `json:"-"`
	Message    string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func New(code int, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
