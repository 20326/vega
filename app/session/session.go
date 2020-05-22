package session

import (
	"encoding/json"
	"strings"

	"github.com/20326/vega/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	// "github.com/satori/go.uuid"
)

// New returns a new cookie-based session management.
func New(users model.UserService, config Config) model.Session {
	return &session{
		secret:     []byte(config.Secret),
		secure:     config.Secure,
		expiration: config.Expiration,
		users:      users,
	}
}

type session struct {
	users      model.UserService
	secret     []byte
	secure     bool
	expiration int64

	UID      uint64 // user ID
	Username string // username
	TOKEN    string // user token

	administrator string // administrator account
	prometheus    string // prometheus account
	autoscaler    string // autoscaler account
}

func (s *session) Create(c *gin.Context, user *model.User) error {
	session := sessions.Default(c)
	//
	s.UID = user.ID
	s.Username = user.Username
	s.TOKEN = user.Token
	sessionDataBytes, err := json.Marshal(s)
	if nil != err {
		return err
	}
	session.Set("data", string(sessionDataBytes))

	return session.Save()
}

func (s *session) Delete(c *gin.Context) error {
	defaultSession := sessions.Default(c)
	defaultSession.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	defaultSession.Clear()
	if err := defaultSession.Save(); nil != err {
		return err
	}
	return nil
}

func (s *session) Get(c *gin.Context) (*model.User, error) {
	switch {
	case isAuthorizationToken(c):
		return s.fromToken(c)
	case isAuthorizationParameter(c):
		return s.fromToken(c)
	default:
		return s.fromSession(c)
	}
}

func (s *session) fromSession(c *gin.Context) (*model.User, error) {

	ret := &session{}
	session := sessions.Default(c)
	sessionDataStr := session.Get("data")
	if nil == sessionDataStr {
		return nil, nil
	}

	err := json.Unmarshal([]byte(sessionDataStr.(string)), ret)
	if nil != err {
		return nil, err
	}
	c.Set("session", ret)

	return s.users.FindName(c, ret.Username)
}

func (s *session) fromToken(c *gin.Context) (*model.User, error) {
	return s.users.FindToken(c,
		extractToken(c),
	)
}

func isAuthorizationToken(c *gin.Context) bool {
	return c.Query("Authorization") != ""
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
