package model

import "context"

type (
	// Role represents a user role of the system.
	// Role
	Role struct {
		Model

		Name        string       `gorm:"size:64;unique_index" json:"name"`
		Label       string       `gorm:"size:64" json:"label"`
		Describe    string       `gorm:"size:256" json:"describe"`
		CreateBy    string       `gorm:"size:64" json:"createBy"`
		Status      int          `gorm:"default:1" json:"status"`
		Deleted     int          `gorm:"default:0" json:"deleted"`
		Users       []*User      `gorm:"many2many:user_roles;association_jointable_foreignkey:user_id" json:"users"`
		Permissions []Permission `json:"permissions"`
	}

	// RoleService defines operations for working with system roles.
	RoleService interface {
		// Find returns a role from the datastore.
		Find(context.Context, uint64) (*Role, error)

		// FindName returns a role from the datastore.
		FindName(context.Context, string) (*Role, error)

		// List returns a list of roles from the datastore.
		List(context.Context) ([]*Role, error)

		// Update persists a role to the datastore.
		Update(context.Context, *Role) error

		// Delete deletes a role from the datastore.
		Delete(context.Context, uint64) error

		// Create persists a new role to the datastore.
		Create(context.Context, *Role) error
	}
)
