package model

import (
	"context"
	"errors"
	"time"

	"github.com/20326/vega/pkg/pagination"
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

		Username     string     `gorm:"size:64" json:"username"`
		Nickname     string     `gorm:"size:64" json:"nickname"`
		Avatar       string     `gorm:"size:255" json:"avatar"`
		BIO          string     `gorm:"size:512" json:"bio"`
		Locale       string     `gorm:"size:64" json:"locale"`
		Password     string     `gorm:"-" json:"password,omitempty"`
		PasswordHash string     `gorm:"size:64" json:"-"`
		Email        string     `gorm:"size:64" json:"email"`
		Phone        string     `gorm:"size:16" json:"phone"`
		Status       int        `gorm:"default:1" json:"status"`
		LoginIP      string     `gorm:"size:32" json:"loginIP"`
		LoginAt      *time.Time `json:"loginAt"`
		Token        string     `gorm:"size:255" json:"-"`
		Refresh      string     `gorm:"size:255" json:"-"`
		Expiry       int64      `gorm:"size:255" json:"-"`
		Hash         string     `gorm:"size:64" json:"-"`
		Roles        []*Role    `gorm:"many2many:user_roles;association_jointable_foreignkey:role_id" json:"roles"`
		RoleList     []uint64   `gorm:"-" json:"roleList" `
	}

	UserRole struct {
		UserID uint64 `sql:"index"`
		RoleID uint64 `sql:"index"`
	}

	// UserService defines operations for working with users.
	UserService interface {
		// Find returns a user from the datastore.
		Find(context.Context, uint64) (*User, error)

		// FindName returns a user from the datastore by username.
		FindName(context.Context, string) (*User, error)

		// FindToken returns a user from the datastore by token.
		FindToken(context.Context, string) (*User, error)

		// FindWhere returns a list of users from the datastore.
		FindWhere(PageQuery, []string) ([]*User, pagination.Pagination)

		// List returns a list of users from the datastore.
		List(context.Context) ([]*User, error)

		// Create persists a new user to the datastore.
		Create(context.Context, *User) error

		// Update persists an updated user to the datastore.
		Update(context.Context, *User) error

		// Update columns persists an updated user to the datastore.
		Updates(context.Context, *User, interface{}) error

		// Delete deletes a user from the datastore.
		Delete(context.Context, uint64) error

		// Count returns a count of users.
		Count(context.Context) (int, error)

		// RelatedClear returns a user from the datastore by token.
		RelatedClear(context.Context, *User)
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

func (u *User) FillRoleList() {
	u.RoleList = []uint64{}
	for _, role := range u.Roles{
		u.RoleList = append(u.RoleList, role.ID)
	}
}

func (u *User) FillRolePermissionList() {
	u.RoleList = []uint64{}
	for _, role := range u.Roles{
		u.RoleList = append(u.RoleList, role.ID)
	}
}