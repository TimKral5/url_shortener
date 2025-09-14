// Package cache handles the temporary storage of data for quick
// access.
package cache

// Connection is an interface all implementations of a cache wrapper
// have to comply with.
type Connection interface {
	Connect(connStr string) error
	Disconnect() error
	AddURL(hash string, url string) error
	GetURL(hash string) (string, error)
}
