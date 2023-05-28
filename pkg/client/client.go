package client

import (
	"fmt"
	"net/http"
)

const (
	DefaultUserAgent = "go-proxmox/dev"
	TagFormat        = "go-proxmox+%s"
)

// func MakeTag(v string) string {
// 	return fmt.Sprintf(TagFormat, v)
// }

type Option func(*Client)

func NewClient(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL:   baseURL,
		userAgent: DefaultUserAgent,
		log:       &LeveledLogger{Level: LevelError},
	}

	for _, o := range opts {
		o(c)
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	return c
}

func (c *Client) Login(username, password string) error {
	_, err := c.Ticket(&Credentials{
		Username: username,
		Password: password,
	})

	return err
}

func (c *Client) APIToken(tokenID, secret string) {
	c.token = fmt.Sprintf("%s=%s", tokenID, secret)
}

func WithClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

func WithLogins(username, password string) Option {
	return func(c *Client) {
		c.credentials = &Credentials{
			Username: username,
			Password: password,
		}
	}
}

func WithAPIToken(tokenID, secret string) Option {
	return func(c *Client) {
		c.token = fmt.Sprintf("%s=%s", tokenID, secret)
	}
}

// WithSession experimental
func WithSession(ticket, csrfPreventionToken string) Option {
	return func(c *Client) {
		c.session = &Session{
			Ticket:              ticket,
			CsrfPreventionToken: csrfPreventionToken,
		}
	}
}

func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

func WithLogger(logger LeveledLoggerInterface) Option {
	return func(c *Client) {
		c.log = logger
	}
}
