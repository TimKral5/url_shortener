// Package database handles the persistent storage of data in this
// project.
package database

// Error is a container for all database related errors.
type Error struct {
	Message string
	Code    int
}

// Connection is an interface all implementations of a database
// wrapper have to comply with.
type Connection interface {
	Connect(connStr string) error
	Disconnect() error
}

// Error returns the error message.
func (err Error) Error() string {
	return err.Message
}
