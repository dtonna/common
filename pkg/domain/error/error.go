package error

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func DomainError(code, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func DomainThirdPartyError(msg string) *Error {
	return &Error{
		Code:    "8889",
		Message: msg,
	}
}
