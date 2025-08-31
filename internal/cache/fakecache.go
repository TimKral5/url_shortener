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
}

// NewFakeCacheConnection constructs a new fake cache connection.
func NewFakeCacheConnection() *FakeCacheConnection {
	return &FakeCacheConnection{
		FailConnect:    false,
		FailDisconnect: false,
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
