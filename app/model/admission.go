package model

import "context"

type (

	// RoleService defines operations for working with system roles.
	AdmissionService interface {
		// LoadAllPolicy returns all policy from the datastore.
		LoadAllPolicy(context.Context, []*Role) error

		// DeleteAllPolicy deletes policy from the datastore.
		DeleteAllPolicy(context.Context) error

		// Admit can be used to restrict access to authorized users, such as
		// members of an organization in your source control management system.
		Admit(ctx context.Context, user *User, subject string, action string) (bool, error)
	}
)
