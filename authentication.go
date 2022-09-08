package bamboo

import (
	"fmt"
	"net/http"
)

type CredentialsConfigurer interface {
  ApplyCredentials(*http.Request)
}

// SimpleCredentials are the username and password used to communicate with the API
type SimpleCredentialsConfigurer struct {
	Username string
	Password string
}

func (c *SimpleCredentialsConfigurer) ApplyCredentials(req *http.Request) {
	req.SetBasicAuth(c.Username, c.Password)
}

type BearerTokenConfigurer struct {
  Token string
}

func (c *BearerTokenConfigurer) ApplyCredentials(req *http.Request) {
  req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
}
