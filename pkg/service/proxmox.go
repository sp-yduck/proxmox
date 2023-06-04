package service

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/sp-yduck/proxmox/pkg/api"
	"github.com/sp-yduck/proxmox/pkg/client"
	"github.com/sp-yduck/proxmox/pkg/service/node"
	storageapi "github.com/sp-yduck/proxmox/pkg/service/node/storage"
	versionapi "github.com/sp-yduck/proxmox/pkg/service/version"
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

// func (c *Service) ClusterStatus() (*cluster.ClusterStatus, error) {
// 	cluster := cluster.Cluster{
// 		Client: c.Client,
// 	}
// 	if err := c.Get("/cluster/status", &cluster); err != nil {
// 		return nil, err
// 	}

// 	return &cluster, nil
// }

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

// func (c *Service) Pools() ([]*poolapi.Pool, error) {
// }

// func (c *Service) Pool() (*poolapi.Pool, error) {
// }

// func (c *Service) CreatePool() (*poolapi.Pool, error) {
// }

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

func (c *Service) Storage(name string) (*storageapi.Storage, error) {
	var storages []*storageapi.Storage
	if err := c.Get("/storage", &storages); err != nil {
		return nil, err
	}
	for _, s := range storages {
		if s.Storage == name {
			s.Client = c
			return s, nil
		}
	}
	return nil, api.ErrNotFound
}

func (c *Service) CreateStorage(name, storageType string, options storageapi.StorageCreateOptions) (*storageapi.Storage, error) {
	var storage *storageapi.Storage
	options.Storage = name
	options.StorageType = storageType
	if err := c.Post("/storage", options, &storage); err != nil {
		return nil, err
	}
	storage.Client = c.Client
	return storage, nil
}

func (c *Service) Version() (*versionapi.Version, error) {
	var version *versionapi.Version
	if err := c.Get("/version", &version); err != nil {
		return nil, err
	}
	return version, nil
}
