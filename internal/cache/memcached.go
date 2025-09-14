package cache

import "github.com/bradfitz/gomemcache/memcache"

type MemcachedConnection struct {
	ConnectionStrings []string
	client *memcache.Client
}

var _ Connection = (*MemcachedConnection)(nil)

func NewMemcachedConnection(connStr ...string) (*MemcachedConnection, error) {
	conn := &MemcachedConnection{
		ConnectionStrings: connStr,
	}

	conn.ConnectToMultiple(connStr...)
	return conn, nil
}

func (conn *MemcachedConnection) Connect(connStr string) error {
	conn.client = memcache.New(connStr)
	return nil
}

func (conn *MemcachedConnection) ConnectToMultiple(connStr ...string) error {
	conn.client = memcache.New(connStr...)
	return nil
}

func (conn *MemcachedConnection) Disconnect() error {
	err := conn.client.Close()
	if err != nil {
		return err
	}

	return nil
}

func (conn *MemcachedConnection) AddURL(hash string, url string) error {
	err := conn.client.Set(&memcache.Item{Key: hash, Value: []byte(url)})
	if err != nil {
		return err
	}

	return nil
}

func (conn *MemcachedConnection) GetURL(hash string) (string, error) {
	item, err := conn.client.Get(hash)
	if err != nil {
		return "", err
	}

	return string(item.Value), nil
}
