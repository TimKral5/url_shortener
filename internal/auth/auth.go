package auth

type AuthError struct {
	Message string
	Code    int
}

type AuthConnection interface {
	Connect(connStr string) error
	Disconnect() error
}

func NewAuthError(message string, code int) AuthError {
	return AuthError{
		Message: message,
		Code:    code,
	}
}

func (self AuthError) Error() string {
	return self.Message
}
