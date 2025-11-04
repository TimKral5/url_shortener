// Package auth handles the storage of user related data, as well as
// authentication and authorization.
package auth

// Error is a container for all authentication related errors.
type Error struct {
	Message string
	Code    int
}

// Connection is an interface all implementations of an auth database
// wrapper have to comply with.
type Connection interface {
	ValidateCredentials(user string, pass string) (bool, error)
	GetPermissions(user string) ([]string, error)
	HasAnyPermission(user string, permissions []string) (bool, error)
	HasAllPermissions(user string, permissions []string) (bool, error)
}

// Error returns the error message.
func (err Error) Error() string {
	return err.Message
}
