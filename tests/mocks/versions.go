package mocks

import (
	"github.com/sp-yduck/proxmox/tests/mocks/pve6x"
	"github.com/sp-yduck/proxmox/tests/mocks/pve7x"
	"github.com/sp-yduck/proxmox/tests/mocks/types"
)

func ProxmoxVE7x(config types.Config) {
	LoadRoutes(config, pve7x.Routes())
}

func ProxmoxVE6x(config types.Config) {
	LoadRoutes(config, pve6x.Routes())
}
