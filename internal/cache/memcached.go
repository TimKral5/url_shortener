package cache

import "github.com/bradfitz/gomemcache/memcache"

// MemcachedConnection is a wrapper for connections to a Memcached
// cache.
type MemcachedConnection struct {
	ConnectionStrings []string
	client            *memcache.Client
}

var _ Connection = (*MemcachedConnection)(nil)

// NewMemcachedConnection constructs a new connection to a Memcached
// instance.
func NewMemcachedConnection(connStr ...string) (*MemcachedConnection, error) {
	conn := &MemcachedConnection{
		ConnectionStrings: connStr,
		client:            nil,
	}

	return conn, conn.ConnectToMultiple(connStr...)
}

// Connect establishes a new connection to a single Memcached
// instance.
func (conn *MemcachedConnection) Connect(connStr string) error {
	conn.client = memcache.New(connStr)

	return nil
}

// ConnectToMultiple establishes a new connection to multiple
// Memcached instances.
func (conn *MemcachedConnection) ConnectToMultiple(connStr ...string) error {
	conn.client = memcache.New(connStr...)

	return nil
}

// Disconnect terminates an active connection to Memcached.
func (conn *MemcachedConnection) Disconnect() error {
	err := conn.client.Close()
	if err != nil {
		return NewMemcachedDisconnectError(err)
	}

	return nil
}

// AddURL inserts a new URL entry into the cache.
func (conn *MemcachedConnection) AddURL(hash string, url string) error {
	err := conn.client.Set(&memcache.Item{
		Key:        hash,
		Value:      []byte(url),
		Flags:      0,
		Expiration: 0,
		CasID:      0,
	})
	if err != nil {
		return NewMemcachedAddError(err)
	}

	return nil
}

// GetURL fetches a URL entry from the cache.
func (conn *MemcachedConnection) GetURL(hash string) (string, error) {
	item, err := conn.client.Get(hash)
	if err != nil {
		return "", NewMemcachedGetError(err)
	}

	return string(item.Value), nil
}
