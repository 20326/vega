package model

import (
	"context"
)

type (
	// Permission represents an auth request of the system.
	// Resource
	Resource struct {
		Model

		ActionID uint64 `json:"actionID"`
		// Name     string `gorm:"size:64" json:"name"`
		// Describe string `gorm:"size:256" json:"describe"`
		Method   string `gorm:"size:64" json:"method"`
		Path     string `gorm:"size:256" json:"path"`
	}

	// ResourceService defines operations for working with system resource.
	ResourceService interface {
		// Find returns a resource from the datastore.
		Find(context.Context, uint64) (*Resource, error)

		// List returns a list of resources from the datastore.
		List(context.Context) ([]*Resource, error)

		// Update persists a resource to the datastore.
		Update(context.Context, *Resource) error

		// Delete deletes a resource from the datastore.
		Delete(context.Context, uint64) error

		// Create persists a new resource to the datastore.
		Create(context.Context, *Resource) error
	}
)
