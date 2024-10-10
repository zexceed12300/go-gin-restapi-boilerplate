package helpers

type ThrowError struct {
	Code    int
	Message string
}

func (e *ThrowError) Error() string {
	return e.Message
}
