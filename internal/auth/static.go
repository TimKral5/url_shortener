package auth

// StaticAuth is a handler for static authentication using a
// preconfigured token.
type StaticAuth struct {
	Token string
}

var _ Connection = (*StaticAuth)(nil)

// NewStaticAuth constructs a new instance of the StaticAuth handler.
func NewStaticAuth(token string) *StaticAuth {
	return &StaticAuth{
		Token: token,
	}
}

// ValidateCredentials validates a set of credentials and returns
// the result.
func (auth *StaticAuth) ValidateCredentials(_ string, pass string) (bool, error) {
	return pass == auth.Token, nil
}

// GetPermissions returns all permissions of a user.
func (auth *StaticAuth) GetPermissions(_ string) ([]string, error) {
	return []string{"administrator"}, nil
}

// HasAnyPermission returns true if any of the listed permissions is
// present. If the user is an administrator, it returns true by
// default.
func (auth *StaticAuth) HasAnyPermission(_ string, _ []string) (bool, error) {
	return true, nil // default for administrators
}

// HasAllPermissions returns true if all of the listed permissions
// are present. If the user is an administrator, it returns true by
// default.
func (auth *StaticAuth) HasAllPermissions(_ string, _ []string) (bool, error) {
	return true, nil // default for administrators
}
