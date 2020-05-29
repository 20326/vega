package model

import (
	"context"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
)

var (
	passwordLen     = 5
	errUsernameLen  = errors.New("invalid username length")
	errUsernameChar = errors.New("invalid character in username")
	errPasswordLen  = errors.New("invalid password length")
)

type (
	// User represents a user of the system.
	User struct {
		Model

		Username string     `gorm:"size:32" json:"username"`
		Nickname string     `gorm:"size:32" json:"nickname"`
		Avatar   string     `gorm:"size:255" json:"avatar"`
		BIO      string     `gorm:"size:512" json:"bio"`
		Locale   string     `gorm:"size:32" json:"locale"`
		Password string     `gorm:"size:64" json:"password,omitempty"`
		Email    string     `gorm:"size:64" json:"email"`
		Phone    string     `gorm:"size:16" json:"phone"`
		Status   int        `gorm:"default:1" json:"status"`
		LoginIP  string     `gorm:"size:32" json:"loginIP"`
		LoginAt  *time.Time `json:"loginAt"`
		Token    string     `gorm:"size:255" json:"-"`
		Refresh  string     `gorm:"size:255" json:"-"`
		Expiry   int64      `gorm:"size:255" json:"-"`
		Hash     string     `gorm:"size:64" json:"-"`

		// Role      []Role     `json:"role" gorm:"many2many:user_role;"`
	}

	// UserService defines operations for working with users.
	UserService interface {
		// Find returns a user from the datastore.
		Find(context.Context, uint64) (*User, error)

		// FindName returns a user from the datastore by username.
		FindName(context.Context, string) (*User, error)

		// FindToken returns a user from the datastore by token.
		FindToken(context.Context, string) (*User, error)

		// List returns a list of users from the datastore.
		List(context.Context) ([]*User, error)

		// Create persists a new user to the datastore.
		Create(context.Context, *User) error

		// Update persists an updated user to the datastore.
		Update(context.Context, *User) error

		// Update columns persists an updated user to the datastore.
		Updates(context.Context, *User, interface{}) error

		// Delete deletes a user from the datastore.
		Delete(context.Context, *User) error

		// Count returns a count of users.
		Count(context.Context) (int, error)
	}
)

// Validate the user and returns an error if the
// validation fails.
func (u *User) Validate() error {
	if passwordLen > len(u.Password) {
		return errPasswordLen
	}

	switch {
	case !govalidator.IsByteLength(u.Username, 5, 30):
		return errUsernameLen
	case !govalidator.Matches(u.Username, "^[a-zA-Z0-9_-]+$"):
		return errUsernameChar
	default:
		return nil
	}
}
