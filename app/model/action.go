package model

import "context"

type (

	// Action
	Action struct {
		Model

		// ActionID  string     `gorm:"size:64" json:"actionID"` // TODO
		Name      string     `gorm:"size:256" json:"name"`
		Describe  string     `gorm:"size:256" json:"describe"`
		Resources []Resource `gorm:"many2many:action_resource;" json:"resources"`
	}

	// ActionService defines operations for working with system actions.
	ActionService interface {
		// Find returns a action from the datastore.
		Find(context.Context, uint64) (*Action, error)

		// List returns a list of actions from the datastore.
		List(context.Context) ([]*Action, error)

		// Update persists a action to the datastore.
		Update(context.Context, *Action) error

		// Delete deletes a action from the datastore.
		Delete(context.Context, uint64) error

		// Create persists a new action to the datastore.
		Create(context.Context, *Action) error
	}
)
