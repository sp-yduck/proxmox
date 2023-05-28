package client

import ()

func (c *Client) Ticket(credentials *Credentials) (*Session, error) {
	return c.session, c.Post("/access/ticket", credentials, &c.session)
}
