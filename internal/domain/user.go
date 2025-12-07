package domain

// User represents the user entity in domain layer.
type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
}

// NewUser creates a new User entity.
func NewUser(name, email, password, role string) User {
	return User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

// IsAdmin checks if user has admin role.
func (u User) IsAdmin() bool {
	return u.Role == "admin"
}

// IsCustomer checks if user has customer role.
func (u User) IsCustomer() bool {
	return u.Role == "customer"
}
