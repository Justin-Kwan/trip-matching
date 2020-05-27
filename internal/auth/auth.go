package auth

type AuthConfig struct {
  WsServerApiKey string
  RestServerApiKey string
}

type consumer struct {
  id       string
  apiKey   string
  resource string
}

// type Authenticator interface {
//   Authenticate() Ticket
// }

// func setConfig(authCfg *config.AuthConfig) {
//
// }
//
// func AuthenticateWsClient(): Ticket, error {
//
// }
//
// func (c *consumer) CanAccess() bool {
//   return c.apiKey ==
// }
