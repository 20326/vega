package session

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/20326/vega/app/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

type legacy struct {
	*session
	mapping map[string]string
}

// Legacy returns a session manager that is capable of mapping
// legacy tokens to 1.0 users using a mapping file.
func Legacy(users model.UserService, config *Config) model.Session {
	base := &session{
		secret:     []byte(config.Secret),
		secure:     config.Secure,
		expiration: config.Expiration,
		users:      users,
	}
	out, err := ioutil.ReadFile(config.MappingFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Read session file has some errors!")
	}
	mapping := map[string]string{}
	err = json.Unmarshal(out, &mapping)
	if err != nil {
		//log.Fatal().Err(err).Msg("Unmarshal session file has some errors!")
	}
	return &legacy{base, mapping}
}

func (s *legacy) Get(c *gin.Context) (*model.User, error) {
	switch {
	case isAuthorizationToken(c):
		return s.fromToken(c)
	case isAuthorizationParameter(c):
		return s.fromToken(c)
	default:
		return s.fromSession(c)
	}
}

func (s *legacy) fromToken(c *gin.Context) (*model.User, error) {
	extracted := extractToken(c)

	// determine if the token is a legacy token based on length.
	// legacy tokens are > 64 characters.
	if len(extracted) < 64 {
		return s.users.FindToken(c, extracted)
	}

	token, err := jwt.Parse(extracted, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Legacy token: invalid signature")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Legacy token: invalid claim format")
		}

		// extract the username claim
		claim, ok := claims["text"]
		if !ok {
			return nil, errors.New("Legacy token: invalid format")
		}

		// lookup the username to get the secret
		secret, ok := s.mapping[claim.(string)]
		if !ok {
			return nil, errors.New("Legacy token: cannot lookup user")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return s.users.FindName(c,token.Claims.(jwt.MapClaims)["text"].(string))
}
