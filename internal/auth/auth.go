package auth

type AuthDatabaseConnection interface {
	Disconnect() error
}
