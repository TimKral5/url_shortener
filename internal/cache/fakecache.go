package cache

type FakeCacheConnection struct {
	FailConnect    bool
	FailDisconnect bool
}

func NewFakeCacheConnection() *FakeCacheConnection {
	return &FakeCacheConnection{
		FailConnect:    false,
		FailDisconnect: false,
	}
}

func (self *FakeCacheConnection) Connect(connStr string) error {
	if self.FailConnect {
		return NewCacheError("Could not connect.", 1)
	}
	return nil
}

func (self *FakeCacheConnection) Disconnect() error {
	if self.FailDisconnect {
		return NewCacheError("Could not disconnect.", 2)
	}
	return nil
}
