package database

type DatabaseError struct {
	Message string
	Code    int
}

type DatastoreConnection interface {
	Connect(connStr string) error
	Disconnect() error
}

func NewDatabaseError(message string, code int) DatabaseError {
	return DatabaseError{
		Message: message,
		Code:    code,
	}
}

func (self DatabaseError) Error() string {
	return self.Message
}
