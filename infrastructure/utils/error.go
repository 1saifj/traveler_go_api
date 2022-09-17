package utils

type Error struct {
	Code     int
	Message  string
	Refrence string
}

func NewError(code int, message string, ref string) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		Refrence: ref,
	}
}
