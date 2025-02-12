package model

import (
	"context"
	"github.com/20326/vega/pkg/pagination"
	mapset "github.com/deckarep/golang-set"
)

type (
	// permission.
	Permission struct {
		Model

		Name         string   `gorm:"size:64;unique_index" json:"name"`
		Label        string   `gorm:"size:64" json:"label"`
		Describe     string   `gorm:"size:256" json:"describe"`
		Icon         string   `gorm:"size:64" json:"icon"`
		Path         string   `gorm:"size:256" json:"path"`
		Actions      []*Action `gorm:"many2many:permission_actions;" json:"actions"`
		DefaultCheck bool     `gorm:"default:false" json:"defaultCheck"`
		Status       int      `gorm:"default:1" json:"status"`
		Deleted      int      `gorm:"default:0" json:"deleted"`

		Selected     []interface{}  `gorm:"-" json:"selected"` //selected actions
	}

	// PermissionService defines operations for working with system permissions.
	PermissionService interface {
		// Find returns a permission from the datastore.
		Find(context.Context, uint64) (*Permission, error)

		// FindWhere returns a list of users from the datastore.
		FindWhere(PageQuery) ([]*Permission, pagination.Pagination)

		// List returns a list of permissions from the datastore.
		List(context.Context) ([]*Permission, error)

		// Update persists a permission to the datastore.
		Update(context.Context, *Permission) error

		// Delete deletes a permission from the datastore.
		Delete(context.Context, uint64) error

		// Create persists a new permission to the datastore.
		Create(context.Context, *Permission) error
	}
)

func (p *Permission) GetActionIds() mapset.Set{
	actionIds := mapset.NewSet()
	for _, action := range p.Actions {
		actionIds.Add(action.ID)
	}
	return actionIds
}
