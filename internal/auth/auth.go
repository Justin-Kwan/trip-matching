package auth

import (
	// "github.com/pkg/errors"
)

type Auth struct {
	config AuthConfig
}

type AuthConfig struct {
	WsServerApiKey   string
	RestServerApiKey string
}

type Courier struct {
	id       string
	apiKey   string
}

type Consumer struct {
	id       string
	apiKey   string
}

type Authenticator interface {
  Authenticate() (string, error)
}

// func setConfig(authCfg *config.AuthConfig) {

// }


// define interface for both user types somewhere else
func (c *Consumer) Authenticate() (Ticket, error) {
	isAuthorized := wc.CanAccess()
	if !isAuthorized {
		return errors.Errorf("")
	}
	return newTicket(wc)
}

// func (c *Courier) CanAccess() bool {
//   return c.apiKey ==
// }

// func (c *Consumer) CanAccess() bool {
//   return c.apiKey ==
// }
