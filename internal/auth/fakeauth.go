package auth

const (
	// FakeConnectError is the code for a failed connection.
	FakeConnectError = iota
	// FakeDisconnectError is the code for a failed termination of a
	// connection.
	FakeDisconnectError
)

// FakeAuthConnection is a mock auth connection for testing.
type FakeAuthConnection struct {
	FailConnect    bool
	FailDisconnect bool
}

// NewFakeAuthConnection constructs a new fake auth connection.
func NewFakeAuthConnection() *FakeAuthConnection {
	return &FakeAuthConnection{
		FailConnect:    false,
		FailDisconnect: false,
	}
}

// Connect emulates the establishment of a connection. Its behaviour
// can be controlled using the FailConnect property.
func (conn *FakeAuthConnection) Connect(_ string) error {
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
func (conn *FakeAuthConnection) Disconnect() error {
	if conn.FailDisconnect {
		return Error{
			Message: "Could not disconnect.",
			Code:    FakeDisconnectError,
		}
	}

	return nil
}
