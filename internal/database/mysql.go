package database

type MySQLConnection struct {
	ConnectionString string
}

func NewMySQLConnection(connStr string) (MySQLConnection, error) {
	return MySQLConnection{
		ConnectionString: connStr,
	}, nil
}

func (self *MySQLConnection) Connect(connStr string) error {
	self.ConnectionString = connStr
	return nil
}

func (self *MySQLConnection) Disconnect() error {
	return nil
}
