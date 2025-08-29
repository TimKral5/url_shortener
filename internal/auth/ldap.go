package auth

type LDAPConnection struct {
	AuthDatabaseConnection
	ConnectionString string
}

func NewLDAPConnection(connStr string) LDAPConnection {
	return LDAPConnection{}
}

func (self *LDAPConnection) Disconnect() error {
	return nil
}

