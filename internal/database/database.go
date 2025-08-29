package database

type DatastoreConnection interface {
	Disconnect() error
}

