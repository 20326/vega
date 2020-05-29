package model

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type (
	// SessionData represents the session.
	SessionData struct {
		UID       uint64 // user ID
		Username  string // username
		Roles     string // user roles
		Token     string // user token
		Anonymous bool   // Anonymous user
	}
)

func (sd *SessionData) Get(c *gin.Context, users UserService) (*User, error) {
	switch {
	case isAuthorizationToken(c):
		return sd.fromToken(c, users)
	case isAuthorizationParameter(c):
		return sd.fromToken(c, users)
	default:
		return sd.fromSession(c, users)
	}
}

// Save saves the current session of the specified context.
func (sd *SessionData) Save(c *gin.Context) error {
	session := sessions.Default(c)
	sessionDataBytes, err := json.Marshal(sd)
	if nil != err {
		return err
	}
	session.Set("data", string(sessionDataBytes))

	return session.Save()
}

func (sd *SessionData) Delete(c *gin.Context, user *User, users UserService) error {
	defaultSession := sessions.Default(c)
	defaultSession.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	defaultSession.Clear()
	// delete token form user service
	defaultSession.Save()
	values := map[string]interface{}{
		"token":    "",
		"login_at": time.Now(),
	}
	return users.Updates(c, user, values)
}

func (sd *SessionData) fromSession(c *gin.Context, users UserService) (*User, error) {
	session := sessions.Default(c)
	sessionDataStr := session.Get("data")
	if nil == sessionDataStr {
		return nil, errors.New("not found session data")
	}
	err := json.Unmarshal([]byte(sessionDataStr.(string)), sd)
	if nil != err {
		return nil, err
	}

	defer c.Set("session", sd)
	if sd.UID > 0 {
		return users.Find(c, sd.UID)
	} else if "" != sd.Token {
		return users.FindToken(c, sd.Token)
	} else {
		sd.Anonymous = true
		return nil, errors.New("anonymous user")
	}
}

func (sd *SessionData) fromToken(c *gin.Context, users UserService) (*User, error) {
	sd.Token = extractToken(c)
	if "" == sd.Token {
		return nil, errors.New("not find token")
	}
	// find user by token
	return users.FindToken(c, sd.Token)
}

func isAuthorizationToken(c *gin.Context) bool {
	return c.GetHeader("Authorization") != ""
}

func isAuthorizationParameter(c *gin.Context) bool {
	return c.Query("access_token") != ""
}

func extractToken(c *gin.Context) string {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		bearer = c.Query("access_token")
	}
	return strings.TrimPrefix(bearer, "Bearer ")
}
