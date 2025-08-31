package database

// MySQLConnection is a wrapper for a connection to a MySQL database.
type MySQLConnection struct {
	ConnectionString string
}

// NewMySQLConnection constructs a new instance of a MySQL connection.
func NewMySQLConnection(connStr string) (MySQLConnection, error) {
	return MySQLConnection{
		ConnectionString: connStr,
	}, nil
}

// Connect establishes a new connection using the provided connection
// string.
func (conn *MySQLConnection) Connect(connStr string) error {
	conn.ConnectionString = connStr

	return nil
}

// Disconnect terminates an active connection.
func (conn *MySQLConnection) Disconnect() error {
	return nil
}
