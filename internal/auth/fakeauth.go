package auth

type FakeAuthConnection struct {
	FailConnect    bool
	FailDisconnect bool
}

func NewFakeAuthConnection() *FakeAuthConnection {
	return &FakeAuthConnection{
		FailConnect:    false,
		FailDisconnect: false,
	}
}

func (self *FakeAuthConnection) Connect(connStr string) error {
	if self.FailConnect {
		return NewAuthError("Could not connect.", 1)
	}
	return nil
}

func (self *FakeAuthConnection) Disconnect() error {
	if self.FailDisconnect {
		return NewAuthError("Could not disconnect.", 2)
	}
	return nil
}
