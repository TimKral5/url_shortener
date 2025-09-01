package cache

const (
	// FakeConnectError is the code for a failed connection.
	FakeConnectError = iota
	// FakeDisconnectError is the code for a failed termination of a
	// connection.
	FakeDisconnectError
)

// FakeCacheConnection is a mock cache connection for testing.
type FakeCacheConnection struct {
	FailConnect    bool
	FailDisconnect bool
	URLs           map[string]string
}

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
		return Error{
			Message: "Could not connect.",
			Code:    FakeConnectError,
		}
	}

	return nil
}

// Disconnect emulates the termination of a connection. Its behaviour
// can be controlled using the FailDisconnect property.
func (conn *FakeCacheConnection) Disconnect() error {
	if conn.FailDisconnect {
		return Error{
			Message: "Could not disconnect.",
			Code:    FakeDisconnectError,
		}
	}

	return nil
}

// AddURL emulates the creation of a new entry in the cache.
func (conn *FakeCacheConnection) AddURL(short string, full string) error {
	conn.URLs[short] = full

	return nil
}

// GetURL emulates fetching a URL from its hash.
func (conn *FakeCacheConnection) GetURL(short string) (string, error) {
	entry := conn.URLs[short]

	return entry, nil
}
