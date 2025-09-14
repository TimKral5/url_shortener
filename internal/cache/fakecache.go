package cache

// FakeCacheConnection is a mock cache connection for testing.
type FakeCacheConnection struct {
	FailConnect    bool
	FailDisconnect bool
	URLs           map[string]string
}

var _ Connection = (*FakeCacheConnection)(nil)

// NewFakeCacheConnection constructs a new fake cache connection.
func NewFakeCacheConnection() *FakeCacheConnection {
	return &FakeCacheConnection{
		FailConnect:    false,
		FailDisconnect: false,
		URLs:           map[string]string{},
	}
}

// Connect emulates the establishment of a connection. Its behaviour
// can be controlled using the FailConnect property.
func (conn *FakeCacheConnection) Connect(_ string) error {
	if conn.FailConnect {
		return NewFakeConnectError()
	}

	return nil
}

// Disconnect emulates the termination of a connection. Its behaviour
// can be controlled using the FailDisconnect property.
func (conn *FakeCacheConnection) Disconnect() error {
	if conn.FailDisconnect {
		return NewFakeDisconnectError()
	}

	return nil
}

// AddURL emulates the creation of a new entry in the cache.
func (conn *FakeCacheConnection) AddURL(hash string, url string) error {
	conn.URLs[hash] = url

	return nil
}

// GetURL emulates fetching a URL from its hash.
func (conn *FakeCacheConnection) GetURL(hash string) (string, error) {
	entry := conn.URLs[hash]

	if entry == "" {
		return "", NewFakeNotFoundError(hash)
	}

	return entry, nil
}
