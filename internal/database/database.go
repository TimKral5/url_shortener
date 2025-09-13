// Package database handles the persistent storage of data in this
// project.
package database

// Connection is an interface all implementations of a database
// wrapper have to comply with.
type Connection interface {
	Connect(connStr string) error
	Disconnect() error
	AddURL(short string, full string) error
	GetURL(short string) (string, error)
}
