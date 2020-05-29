package model

import "context"

type (
	// Role represents a user role of the system.
	// Role
	Role struct {
		Model

		Name     string `gorm:"size:256" json:"name"`
		Describe string `gorm:"size:256" json:"describe"`
		Deleted  int    `gorm:"default:0" json:"deleted"`
		Perms    []Perm `json:"perms"`
	}

	// RoleService defines operations for working with system roles.
	RoleService interface {
		// Find returns a role from the datastore.
		Find(context.Context, uint64) (*Role, error)

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
