package cache

type CacheError struct {
	Message string
	Code    int
}

type CacheConnection interface {
	Connect(connStr string) error
	Disconnect() error
}

func NewCacheError(message string, code int) CacheError {
	return CacheError{
		Message: message,
		Code:    code,
	}
}

func (self CacheError) Error() string {
	return self.Message
}
