package model

import (
	"github.com/gin-gonic/gin"
)

// Session provides session management for
// authenticated users.
type Session interface {
	// Create creates a new user session and writes the
	// session to the http.Response.
	Create(c *gin.Context, user *User) error

	// Delete deletes the user session from the http.Response.
	Delete(c *gin.Context) error

	// Get returns the session from the http.Request. If no
	// session exists a nil user is returned. Returning an
	// error is optional, for debugging purposes only.
	Get(c *gin.Context) (*User, error)
}
