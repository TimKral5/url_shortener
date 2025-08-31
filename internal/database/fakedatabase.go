package database

type FakeDatabaseConnection struct {
	FailConnect    bool
	FailDisconnect bool
}

func NewFakeDatabaseConnection() *FakeDatabaseConnection {
	return &FakeDatabaseConnection{
		FailConnect:    false,
		FailDisconnect: false,
	}
}

func (self *FakeDatabaseConnection) Connect(connStr string) error {
	if self.FailConnect {
		return NewDatabaseError("Could not connect.", 1)
	}
	return nil
}

func (self *FakeDatabaseConnection) Disconnect() error {
	if self.FailDisconnect {
		return NewDatabaseError("Could not disconnect.", 2)
	}
	return nil
}
