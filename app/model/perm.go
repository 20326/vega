package model

import "context"

type (
	// permission.
	Perm struct {
		Model

		Name         string   `gorm:"size:256" json:"name"`
		Describe     string   `gorm:"size:256" json:"describe"`
		Resource     Resource `json:"resource"`
		ResourceID   uint64   `json:"resourceID"`
		Action       Action   `json:"action"`
		ActionID     uint64   `json:"actionID"`
		DefaultCheck bool     `gorm:"default:false" json:"defaultCheck"`
		Status       int      `gorm:"default:1" json:"status"`
		Deleted      int      `gorm:"default:0" json:"deleted"`
	}

	// PermService defines operations for working with system permissions.
	PermService interface {
		// Find returns a permission from the datastore.
		Find(context.Context, uint64) (*Perm, error)

		// List returns a list of permissions from the datastore.
		List(context.Context) ([]*Perm, error)

		// Update persists a permission to the datastore.
		Update(context.Context, *Perm) error

		// Delete deletes a permission from the datastore.
		Delete(context.Context, *Perm) error

		// Create persists a new permission to the datastore.
		Create(context.Context, *Perm) error
	}
)
