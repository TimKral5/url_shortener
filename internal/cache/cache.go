// Package cache handles the temporary storage of data for quick
// access.
package cache

// Error is a container for all database related errors.
type Error struct {
	Message string
	Code    int
}

// Connection is an interface all implementations of a cache wrapper
// have to comply with.
type Connection interface {
	Connect(connStr string) error
	Disconnect() error
	AddURL(short string, full string) error
	GetURL(short string) (string, error)
}

// Error returns the error message.
func (err Error) Error() string {
	return err.Message
}
