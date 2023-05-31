package storage

import (
	"fmt"
)

type Volume struct {
	Format string `json:",omitempty"`
	Path   string `json:",omitempty"`
	Size   int    `json:",omitempty"`
	Used   int    `json:",omitempty"`
}

func (c *Content) Volume() (*Volume, error) {
	var volume *Volume
	path := contentPath(c.Node, c.Storage)
	if err := c.Client.Get(fmt.Sprintf("%s/%s", path, c.VolID), &volume); err != nil {
		return nil, err
	}
	return volume, nil
}
