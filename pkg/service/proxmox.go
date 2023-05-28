package service

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/sp-yduck/proxmox/pkg/api"
	"github.com/sp-yduck/proxmox/pkg/client"
	"github.com/sp-yduck/proxmox/pkg/service/cluster"
	"github.com/sp-yduck/proxmox/pkg/service/node"
	storageapi "github.com/sp-yduck/proxmox/pkg/service/node/storage"
)

type Service struct {
	*client.Client
}

// to do : base client option
func NewServiceWithLogin(url, user, password string) (*Service, error) {
	base := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	client := client.NewClient(fmt.Sprintf("https://%s/api2/json", url), client.WithClient(&base))
	if err := client.Login(user, password); err != nil {
		return nil, errors.Errorf("failed to login proxmox: %v", err)
	}
	return &Service{Client: client}, nil
}

func (c *Service) Cluster() (*cluster.Cluster, error) {
	cluster := cluster.Cluster{
		Client: c.Client,
	}
	if err := c.Get("/cluster/status", &cluster); err != nil {
		return nil, err
	}

	return &cluster, nil
}

func (c *Service) Nodes() ([]*node.Node, error) {
	var nodes []*node.Node
	if err := c.Get("/nodes", &nodes); err != nil {
		return nil, err
	}
	for _, n := range nodes {
		n.Client = c.Client
	}
	return nodes, nil
}

func (c *Service) Node(name string) (*node.Node, error) {
	var nodes []*node.Node
	if err := c.Get("/nodes", &nodes); err != nil {
		return nil, err
	}
	for _, n := range nodes {
		if n.Node == name {
			n.Client = c.Client
			return n, nil
		}
	}
	return nil, api.ErrNotFound
}

func (c *Service) Storages() ([]*storageapi.Storage, error) {
	var storages []*storageapi.Storage
	if err := c.Get("/storage", &storages); err != nil {
		return nil, err
	}
	for _, s := range storages {
		s.Client = c.Client
	}
	return storages, nil
}

// func (c *Client) Storage(name string) (*Storage, error) {
// 	var storages []*Storage
// 	if err := c.Get("/storage", &storages); err != nil {
// 		return nil, err
// 	}
// 	for _, s := range storages {
// 		if s.Storage == name {
// 			s.client = c
// 			return s, nil
// 		}
// 	}
// 	return nil, ErrNotFound
// }

// func (c *Client) CreateStorage(name, storageType string) (*Storage, error) {
// 	var storage *Storage
// 	data := make(map[string]interface{})
// 	data["storage"] = name
// 	data["type"] = storageType
// 	if err := c.Post("/storage", data, &storage); err != nil {
// 		return nil, err
// 	}
// 	return storage, nil
// }

func (c *Service) Version() (*Version, error) {
	var version *Version
	if err := c.Get("/version", &version); err != nil {
		return nil, err
	}
	return version, nil
}
