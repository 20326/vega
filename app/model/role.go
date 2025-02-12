package model

import (
	"context"
	"github.com/deckarep/golang-set"
)

type (
	// Role represents a user role of the system.
	// Role
	Role struct {
		Model

		Name        string        `gorm:"size:64;unique_index" json:"name"`
		Label       string        `gorm:"size:64" json:"label"`
		Describe    string        `gorm:"size:256" json:"describe"`
		CreateBy    string        `gorm:"size:64" json:"createBy"`
		Status      int           `gorm:"default:1" json:"status"`
		Deleted     int           `gorm:"default:0" json:"deleted"`
		Permissions []*Permission `gorm:"many2many:role_permissions;" json:"permissions"`
		Actions     []*Action     `gorm:"many2many:role_actions;" json:"actions"`
		// Users       []*User       `gorm:"many2many:user_roles;association_jointable_foreignkey:user_id" json:"users"`
	}

	RoleAction struct {
		RoleID       uint64 `sql:"index"`
		ActionID     uint64 `sql:"index"`
		// PermissionID uint64 `sql:"index"`
	}

	// RoleService defines operations for working with system roles.
	RoleService interface {
		// Find returns a role from the datastore.
		Find(context.Context, uint64) (*Role, error)

		// FindName returns a role from the datastore.
		FindName(context.Context, string) (*Role, error)

		// List returns a list of roles from the datastore.
		List(context.Context) ([]*Role, error)

		// Clear Related persists a role to the datastore.
		RelatedClear(context.Context, *Role)

		// Update persists a role to the datastore.
		Update(ctx context.Context, role *Role, actionIDs []interface{}) error

		// Delete deletes a role from the datastore.
		Delete(context.Context, uint64) error

		// Create persists a new role to the datastore.
		Create(context.Context, *Role) error
	}
)

func (u *Role) GetActionIds() mapset.Set{
	actionIds := mapset.NewSet()
	for _, action := range u.Actions {
		actionIds.Add(action.ID)
	}
	return actionIds
}
