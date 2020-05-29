package model

import (
	"context"
)

type (
	// Perm represents an auth request of the system.
	// Resource
	Resource struct {
		Model

		Name     string   `gorm:"size:256" json:"name"`
		Describe string   `gorm:"size:256" json:"describe"`
		Object   string   `gorm:"size:256" json:"object"`
		Actions  []Action `json:"actions"`
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
