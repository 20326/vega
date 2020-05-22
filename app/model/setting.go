package model

import "context"

type (
	// Setting represents a user setting of the system.
	// Setting
	Setting struct {
		Model

		Name     string `sql:"index" gorm:"size:64" json:"name"`
		Value    string `gorm:"type:text" json:"value"`
		Describe string `gorm:"size:text" json:"describe"`
	}

	// SettingService defines operations for working with system settings.
	SettingService interface {
		// Find returns a setting from the datastore.
		Find(context.Context, uint64) (*Setting, error)

		// List returns a list of settings from the datastore.
		List(context.Context) ([]*Setting, error)

		// Update persists a setting to the datastore.
		Update(context.Context, *Setting) error

		// Delete deletes a setting from the datastore.
		Delete(context.Context,  uint64) error

		// Create persists a new setting to the datastore.
		Create(context.Context, *Setting) error
	}
)
