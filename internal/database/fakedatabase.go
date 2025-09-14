package database

// FakeDatabaseConnection is a mock database connection for testing.
type FakeDatabaseConnection struct {
	FailConnect    bool
	FailDisconnect bool
	URLs           map[string]string
}

var _ Connection = (*FakeDatabaseConnection)(nil)

// NewFakeDatabaseConnection constructs a fake database connection.
func NewFakeDatabaseConnection() *FakeDatabaseConnection {
	return &FakeDatabaseConnection{
		FailConnect:    false,
		FailDisconnect: false,
		URLs:           map[string]string{},
	}
}

// Connect emulates the establishment of a new connection. Its
// behaviour can be controlled using the FailConnect property.
func (conn *FakeDatabaseConnection) Connect(_ string) error {
	if conn.FailConnect {
		return NewFakeConnectError()
	}

	return nil
}

// Disconnect emulates the termination of a connection. Its behaviour
// can be controlled using the FailDisconnect property.
func (conn *FakeDatabaseConnection) Disconnect() error {
	if conn.FailDisconnect {
		return NewFakeDisconnectError()
	}

	return nil
}

// AddURL emulates the creation of a new entry in the database.
func (conn *FakeDatabaseConnection) AddURL(hash string, url string) error {
	conn.URLs[hash] = url

	return nil
}

// GetURL emulates fetching a URL from its hash.
func (conn *FakeDatabaseConnection) GetURL(hash string) (string, error) {
	entry := conn.URLs[hash]

	if entry == "" {
		return "", NewFakeNotFoundError(hash)
	}

	return entry, nil
}
