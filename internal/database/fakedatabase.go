package database

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const (
	// FakeConnectError is the code for a failed connection.
	FakeConnectError = iota
	// FakeDisconnectError is the code for a failed termination of a
	// connection.
	FakeDisconnectError
)

// FakeDatabaseConnection is a mock database connection for testing.
type FakeDatabaseConnection struct {
	FailConnect    bool
	FailDisconnect bool
	URLs           map[string]string
}

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
		return Error{
			Message: "Could not connect.",
			Code:    FakeConnectError,
		}
	}

	return nil
}

// Disconnect emulates the termination of a connection. Its behaviour
// can be controlled using the FailDisconnect property.
func (conn *FakeDatabaseConnection) Disconnect() error {
	if conn.FailDisconnect {
		return Error{
			Message: "Could not disconnect.",
			Code:    FakeDisconnectError,
		}
	}

	return nil
}

// CreateURL emulates the creation of a new shortened URL.
func (conn *FakeDatabaseConnection) CreateURL(full string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(full))
	hashStr := strings.ToUpper(hex.EncodeToString(hash.Sum(nil))[:10])

	conn.URLs[hashStr] = full

	return hashStr, nil
}

// GetURL emulates fetching a URL from its hash.
func (conn *FakeDatabaseConnection) GetURL(hash string) (string, error) {
	entry := conn.URLs[hash]

	return entry, nil
}
